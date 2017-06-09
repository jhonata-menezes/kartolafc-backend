package kartolafc

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sort"
	"github.com/jhonata-menezes/kartolafc-backend/api"
)

type AtletaRanking struct {
	AtletaId int `json:"atleta_id"`
}

type AtletasRanking struct {
	Pontuacao float32
	Atletas []AtletaRanking `bson:"atletas"`
	TimeCompleto struct{
		TimeId int
		Assinante bool `json:"assinante" bson:"assinante"`
	} `bson:"timecompleto"`
	Mensagem string `bson:"mensagem"`
}

type SelectAtletasRanking struct {
	Atletas int `bson:"atletas.atletaid" json:"atletas.atletaid"`
	TimeId int `json:"timecompleto.timeid" bson:"timecompleto.timeid"`
	Assinante int `json:"timecompleto.assinante" bson:"timecompleto.assinante"`
	Mensagem int `bson:"mensagem"`
}

func getSelect() SelectAtletasRanking {
	s := SelectAtletasRanking{}
	s.Atletas = 1
	s.TimeId = 1
	s.Assinante = 1
	s.Mensagem = 1
	return s
}

type TimesRanking []AtletasRanking

type TimeRankingFormated struct {
	TimeId int `json:"time_id"`
	Pontuacao float32 `json:"pontuacao"`
	Assinante bool `json:"assinante" bson:"assinante"`
	Atletas []AtletaRanking `bson:"atletas" json:"atletas,omitempty"`
}

// para melhor apresentacao no endpoint
type TimesRankingFormated []TimeRankingFormated

// para apresentar a posicao atraves do ID/Slug
type TimeIdRanking struct {
	TimeId int `json:"time_id"`
	Pontuacao float32 `json:"pontuacao"`
	Posicao int `json:"posicao"`
	Assinante bool `json:"assinante" bson:"assinante"`
}

func (a TimesRankingFormated) Len() int {
	return len(a)
}

func (a TimesRankingFormated) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a TimesRankingFormated) Less(i, j int) bool {
	a[i].Pontuacao = SomaPontuacao(a[i])
	a[j].Pontuacao = SomaPontuacao(a[j])
	return a[i].Pontuacao > a[j].Pontuacao
}

func SomaPontuacao(atletasTime TimeRankingFormated) float32 {
	var soma float32
	for _, a := range atletasTime.Atletas {

		// se o jogador nao existe no map pode retornar um erro
		if _, ok := CachePontuados.Atletas[a.AtletaId]; ok {
			soma+= CachePontuados.Atletas[a.AtletaId].Pontuacao
		}
	}
	return soma
}

func LoadInMemory(collection *mgo.Collection) {
	time.Sleep(5 * time.Second)
	// se houver request enquanto faz o load do db gerar: index out of range
	CacheRankingTimeIdPontuados = make([]TimeIdRanking, 15000000)

	inicio := time.Now()

	var atl TimesRanking
	// selecionando apenas as fields desejadas
	err := collection.Find(bson.M{}).Select(getSelect()).All(&atl)
	if err != nil {
		panic(err)
	}
	log.Println("pegar todos times da collection", time.Since(inicio))

	// formatando os dados e criando array de tamamnho especifico
	atletasFormatado := make([]TimeRankingFormated, len(atl))
	for k, a := range atl {
		// time que nao tem escalacao, não precisa ser ordenado
		if a.Mensagem == "Este time ainda não foi escalado na temporada." {
			continue
		}
		timeTemp := TimeRankingFormated{}
		timeTemp.TimeId = a.TimeCompleto.TimeId
		timeTemp.Pontuacao = a.Pontuacao
		timeTemp.Atletas = a.Atletas
		timeTemp.Assinante = a.TimeCompleto.Assinante
		atletasFormatado[k] = timeTemp
	}
	log.Println("atualizado array de times com a posicao no indice")
	go SortPontuados(atletasFormatado)
}

func SortPontuados(times TimesRankingFormated) {
	// cache local para comparar, se houve mudanca na pontuacao entao e efetuado sort, assim economiza processamento
	CacheLocalPontuados := CachePontuados
	for {
		inicio := time.Now()
		sort.Sort(times)
		CacheRankingPontuados = times
		log.Println("sort", time.Since(inicio))

		// outro array de times, porem a chave e o time_id
		for k, t := range times {
			timeTemp := TimeIdRanking{}
			timeTemp.Pontuacao = t.Pontuacao
			timeTemp.Posicao = (k+1)
			timeTemp.TimeId = t.TimeId
			timeTemp.Assinante = t.Assinante
			CacheRankingTimeIdPontuados[t.TimeId] = timeTemp
		}
		log.Println("atualizado array de times com time_id no indice")

		melhores()
		//aguarda 1 minuto para fazer o sort novamente
		for PontuadosSaoIguais(CacheLocalPontuados, CachePontuados) {
			time.Sleep(60 * time.Second)
		}
		// atualiza o cache local
		CacheLocalPontuados = CachePontuados
	}
}

func melhores() {
	melhores := make([]TimeRankingFormated, 100)
	if len(CacheRankingPontuados) >= 100 {
		for k, temp := range CacheRankingPontuados[:100] {
			melhores[k] = temp
			melhores[k].Atletas = nil
		}
		CacheRankingPontuadosMelhores = melhores

		melhores = make([]TimeRankingFormated, 100)
		qtd := 0
		for _, temp := range CacheRankingPontuados {
			if temp.Assinante == true {
				melhores[qtd] = temp
				melhores[qtd].Atletas = nil
				qtd++
			}

			if qtd >= 99 {
				break
			}
		}
		CacheRankingPontuadosMelhoresPro = melhores
	}
}

func PontuadosSaoIguais(old, new api.Pontuados) bool {
	for atletaId, desc := range new.Atletas {
		if (old.Atletas[atletaId].Pontuacao == desc.Pontuacao) {
			continue
		}
		return false
	}
	return true
}