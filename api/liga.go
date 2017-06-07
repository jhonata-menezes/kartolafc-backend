package api

import (
	"log"
	"encoding/json"
	"strconv"
)

const URL_LIGA = "/auth/liga/"

type Liga struct {
	Liga struct{
		LigaId int `json:"liga_id"`
		TimeDonoId int `json:"time_dono_id"`
		ClubeId int `json:"clube_id"`
		Nome string `json:"nome"`
		Descricao string `json:"descricao"`
		Slug string `json:"slug"`
		UrlFlamulaSvg string `json:"url_flamula_svg"`
		Imagem string `json:"imagem"`
		QuantidadeTimes int `json:"quantidade_times"`
		MataMata bool `json:"mata_mata"`
		Tipo string `json:"tipo"`
		TotalTimesLiga int `json:"total_times_liga"`

	} `json:"liga"`

	Times []struct{
		TimeId int `json:"time_id"`
		Nome string `json:"nome"`
		NomeCartola string `json:"nome_cartola"`
		Slug string `json:"slug"`
		FacebookId int64 `json:"facebook_id"`
		UrlEscudoSvg string `json:"url_escudo_svg"`
		FotoPerfil string `json:"foto_perfil"`
		Assinante bool `json:"assinante"`
		Patrimonio float32 `json:"patrimonio"`
		Rankinkg struct{
			Rodada int `json:"rodada"`
			Mes int `json:"mes"`
			Turno int `json:"turno"`
			Campeonato int `json:"campeonato"`
			Patrimonio int `json:"patrimonio"`
		} `json:"ranking"`

		Pontos struct{
			Rodada float32 `json:"rodada"`
			Mes float32 `json:"mes"`
			Turno float32 `json:"turno"`
			Campeonato float32 `json:"campeonato"`
		} `json:"pontos"`

		Variacao struct{
			Mes int `json:"mes"`
			Turno int `json:"turno"`
			Campeonato int `json:"campeonato"`
			Patrimonio int `json:"patrimonio"`
		} `json:"variacao"`

	} `json:"times"`
}

func (l *Liga) GetLiga(page int) {
	request := Request{}
	if res, err := request.Get(URL_LIGA+l.Liga.Slug + "?page=" + strconv.Itoa(page), 10);
		err != nil || res.StatusCode() != 200 {
		log.Println("liga id status", res.StatusCode(), "err", err )
	}else{
		err = json.Unmarshal(res.Body(), &l)
		if err != nil {
			log.Println("liga id parse json")
		}
	}

}
