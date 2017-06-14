package api

import (
	"fmt"
	"log"
	"encoding/json"
	"strings"
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
	Atletas []Atleta `json:"atletas"`
	TimeCompleto Time `json:"time"`
	Patrimonio int `json:"patrimonio"`
	EsquemaId int `json:"esquema_id"`
	ValorTime float32 `json:"valor_time"`
	Mensagem string `json:"mensagem"`
	RodadaAtual int `json:"rodada_atual"`
	Pontos float32 `json:"pontos"`
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
		t.ChangeFormatDefault()
	}
}

func (t *TimeCompleto) MountUrl() string {
	return fmt.Sprintf(URL_TIME_ID, t.TimeCompleto.TimeId)
}

func (a *TimeCompleto) ChangeFormatDefault() {
	for i, des := range a.Atletas {
		a.Atletas[i].Foto = strings.Replace(des.Foto, "FORMATO", "140x140", 3)
	}
}