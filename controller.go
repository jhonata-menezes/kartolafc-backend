package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"net/http"
	"github.com/pressly/chi"
	"strconv"
	"github.com/pressly/chi/render"
	"github.com/jhonata-menezes/kartolafc-backend/notification"
	"gopkg.in/mgo.v2/bson"
	"log"
	"io/ioutil"
	"encoding/json"
)

type DefaultMessage struct {
	Status string `json:"status"`
	Mensagem string `json:"mensagem"`
}

type LoginSenha struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type TokenGLobo struct {
	Token string `json:"token"`
}

func (l *TokenGLobo) Bind(r *http.Request) error {
	return nil
}

func (l *LoginSenha) Bind(r *http.Request) error {
	return nil
}

func GetHome(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, DefaultMessage{"ok", "Birll"})
}

func GetStatus(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, CacheStatus)
}

func GetTimes(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	timePesquisado := chi.URLParam(request, "q")
	times := api.Times{}
	times.Pesquisa = timePesquisado
	times.GetTimes()

	render.JSON(response, request, times)
}

func GetTimeHistorico(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	idString := chi.URLParam(request, "id")
	rodadaString := chi.URLParam(request, "rodada")
	time := api.TimeHistorico{}

	id, err := strconv.Atoi(idString)
	if err != nil {
		render.JSON(response, request, DefaultMessage{"error", "parametor id tem que ser numerico"})
		return
	}

	rodada, err := strconv.Atoi(rodadaString)
	if err != nil {
		render.JSON(response, request, DefaultMessage{"error", "parametor rodada tem que ser numerico"})
		return
	}

	if rodada >= CacheStatus.RodadaAtual && CacheStatus.StatusMercado != 1 {
		render.JSON(response, request, DefaultMessage{"error", "rodada informada invalida"})
		return
	}
	// pega uma conexao com mongodb da fila
	c := <- ChannelCollectionTime

	// existe na collection?
	err = c.Database.C("times_historico").Find(bson.M{"timecompleto.timeid": id, "rodadaatual": rodada}).One(&time)

	if err != nil {
		for i:=0; i<3; i++ {
			time.TimeCompleto.TimeId = id
			time.RodadaAtual = rodada
			time.GetTime()

			if time.TimeCompleto.Nome != "" {
				break
			}
		}
		if time.TimeCompleto.Nome == "" {
			render.JSON(response, request, DefaultMessage{"error", "problemas ao recuperar historico"})
		} else {
			c.Database.C("times_historico").Insert(time)
			render.JSON(response, request, time)
		}
	} else {
		render.JSON(response, request, time)
	}
	ChannelCollectionTime <- c
}

func GetTime(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	idString := chi.URLParam(request, "id")
	time := api.TimeCompleto{}

	id, err := strconv.Atoi(idString)

	if err != nil {
		render.JSON(response, request, DefaultMessage{"error", "parametor id tem que ser um numerico"})
	} else {
		// pega uma conexao com mongodb da fila
		c := <- ChannelCollectionTime
		time.TimeCompleto.TimeId = id
		rodadaAtual := CacheStatus.RodadaAtual
		if (CacheStatus.StatusMercado == 1) {
			rodadaAtual-=1
		}
		err := c.Find(bson.M{"timecompleto.timeid": id, "rodadaatual": rodadaAtual}).One(&time)

		// caso nao encontre na collection, requisita a api do cartolafc
		if err != nil {
			time.GetTime()
			render.JSON(response, request, time)
			log.Println(err)
			if err = c.Insert(time); time.TimeCompleto.Nome != "" && err == nil {
				ChannelAddTimeRanking <- time
				log.Println("time adicionado a collection e ao ranking", time.TimeCompleto.TimeId)
			}

		} else {
			// caso exista na base, retorna o registro
			render.JSON(response, request, time)
		}
		// devolve a conexao para a fila
		ChannelCollectionTime <- c
	}
}

func GetMercado(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, CacheKartolaAtletas)
}

func GetDestaques(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, CacheDestaques)
}

func GetLigas(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	nome := chi.URLParam(request, "q")
	pesquisa := api.PesquisaLigas{}
	pesquisa.GetPesquisaLigas(nome)
	render.JSON(response, request, pesquisa)
}

func GetLiga(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	var page int
	pageString := chi.URLParam(request, "page")
	if pageString != "" {
		pageOne , err := strconv.Atoi(pageString)
		if err != nil || pageOne < 1 || pageOne > 5 {
			render.JSON(response, request, DefaultMessage{"error", "id informado nao é numerico ou menor que 1 ou maior que 5"})
			return;
		}
		page = pageOne
	}
	slug := chi.URLParam(request, "id")
	liga := api.Liga{}
	liga.Liga.Slug = slug
	liga.GetLiga(page)
	render.JSON(response, request, liga)
}

func GetPontuados(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)

	render.JSON(response, request, CachePontuados)
}

func GetPartida(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	partidaString := chi.URLParam(request, "partida")
	partida, err := strconv.Atoi(partidaString)
	if err != nil || partida < 0 || partida > 38 {
		render.JSON(response, request, DefaultMessage{"error", "Rodada invalida."})
		return
	}
	if (partida >= 0 && partida <= 38) {
		render.JSON(response, request, CachePartidas[partida])
	}
}

func AddNotificacao(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	client := notification.Subscription{}
	render.DecodeJSON(request.Body, &client)
	notification.ChannelSubscribe <- client
	render.JSON(response, request,  DefaultMessage{"ok", "Inscrito com sucesso"})
}

func GetMelhoresRanking(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	if len(CacheRankingPontuadosMelhores) == 0 {
		render.JSON(response, request, DefaultMessage{"error", "melhores da rodada nao esta disponivel no momento"})
		return
	}
	render.JSON(response, request, CacheRankingPontuadosMelhores)
}

func GetMelhoresRankingPro(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	if len(CacheRankingPontuadosMelhores) == 0 {
		render.JSON(response, request, DefaultMessage{"error", "melhores da rodada nao esta disponivel no momento"})
		return
	}
	render.JSON(response, request, CacheRankingPontuadosMelhoresPro)
}

func GetRankingTimeId(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	timeId, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil ||
		timeId < 1 ||
		timeId >= 15000000 ||
		CacheRankingTimeIdPontuados[timeId].TimeId == 0{

		render.JSON(response, request, DefaultMessage{"error", "id informado nao existe"})
	}else{
		render.JSON(response, request, CacheRankingTimeIdPontuados[timeId])
	}
}

func GetPontuacaHistorico(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	atletaId, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil || atletaId <= 0 || atletaId >= 100000 {
		render.JSON(response, request, DefaultMessage{"error", "id do atleta nao e valido"})
	} else {
		render.JSON(response, request, CacheHistoricoAtleta[atletaId])
	}
}

func PostLogin(response http.ResponseWriter, request *http.Request) {
	data := &LoginSenha{}
	render.Bind(request, data)
	login := api.Login{}
	login.Login(data.Email, data.Senha)
	render.JSON(response, request, login)
}

func GetMeuTime(response http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("token")
	if token != "" {
		meuTime := api.MeuTime{}
		meuTime.Get(token)
		render.JSON(response, request, meuTime)
		return
	}
	render.JSON(response, request, DefaultMessage{"error", "token vazio"})
}

func PostSalvarTime(response http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("token")
	bodyBytes, _ := ioutil.ReadAll(request.Body)
	escalacao := api.SalvarTime{}
	json.Unmarshal(bodyBytes, &escalacao)
	escalacao.Post(token)
	render.JSON(response, request, escalacao)
}



func responseDefault(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}