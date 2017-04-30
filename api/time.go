package api

import (
	"fmt"
	"log"
	"encoding/json"
)

const URL_TIME_ID = "/time/id/%d"

type Time struct {
	TimeId int `json:"time_id"`
	ClubeId int `json:"clube_id"`
	EsquemaId int `json:"esquema_id"`
	CadunId int `json:"cadun_id"`

	Nome string `json:"nome"`
	NomeCartola string `json:"nome_cartola"`
	Slug string `json:"slug"`
	FacebookId int64 `json:"facebook_id"`
	UrlEscudoPng string `json:"url_escudo_png"`
	UrlEscudoSvg string `json:"url_escudo_svg"`
	FotoPerfil string `json:"foto_perfil"`
	Assinante bool `json:"assinante"`
}

type TimeCompleto struct {
	TimeCompleto Time `json:"time"`
	Mensagem string `json:"mensagem"`
	RodadaAtual int `json:"rodada_atual"`
}

func (t *TimeCompleto) GetTime(){
	request := Request{}
	res, err := request.Get(t.MountUrl(), 10)

	if err != nil {
		log.Println(err)
	}
	if res.StatusCode() != 200 {
		log.Println("time id diferente de 200", res.StatusCode(), "response: ", string(res.Body()))
	}else{
		json.Unmarshal(res.Body(), &t)
	}
}

func (t *TimeCompleto) MountUrl() string {
	return fmt.Sprintf(URL_TIME_ID, t.TimeCompleto.TimeId)
}
