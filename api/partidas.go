package api

import (
	"strconv"
	"github.com/valyala/fasthttp"
	"encoding/json"
	"log"
)

const URL_PARTIDAS = "/partidas"

type Partida struct {
	ClubeCasaId int `json:"clube_casa_id"`
	ClubeCasaPosicao int `json:"clube_casa_posicao"`
	ClubeVisitanteId int `json:"clube_visitante_id"`
	ClubeVisitantePosicao int `json:"clube_visitante_posicao"`
	PartidaData string `json:"partida_data"`
	Local string `json:"local"`
	Valida bool `json:"valida"`
	PlacarOficialMandante int `json:"placar_oficial_mandante"`
	PlacarOficialVisitante int `json:"placar_oficial_visitante"`
	UrlConfronto string `json:"url_confronto"`
	UrlTransmissao string `json:"url_transmissao"`
	AproveitamentoMandante []string `json:"aproveitamento_mandante"`
	AproveitamentoVisitante []string `json:"aproveitamento_visitante"`
}

type Partidas struct {
	Partidas []Partida `json:"partidas"`
	Clubes map[int] struct{
		Id int `json:"id"`
		Nome string `json:"nome"`
		Abreviacao string `json:"abreviacao"`
		Posicao int `json:"posicao"`
		Escudos map[string]string
	} `json:"clubes"`
	Rodada int `json:"rodada"`
	Mensagem string `json:"mensagem"`
}

func (p *Partidas) Get(rodada int) {
	request := Request{}
	var res *fasthttp.Response
	var err error
	if rodada >= 1 && rodada <= 20 {
		rodadaParse := URL_PARTIDAS + "/" + strconv.Itoa(rodada)
		res, err = request.Get(rodadaParse, 10)
	} else {
		res, err = request.Get(URL_PARTIDAS, 10)
	}
	if err != nil || res.StatusCode() != 200 {
		log.Println("nao foi possivel obter a rodada", rodada, err)
		return
	}
	if err = json.Unmarshal(res.Body(), &p); err != nil {
		log.Println("nao decodificou o json da partida", rodada)
	}
}
