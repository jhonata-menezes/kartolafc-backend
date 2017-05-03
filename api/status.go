package api

import (
	"encoding/json"
)

const URI_STATUS = "/mercado/status"

type Status struct {
	RodadaAtual int `json:"rodada_atual"`
	StatusMercado int `json:"status_mercado"`
	EsquemaDefaultId int `json:"esquema_default_id"`
	TimesEscalados int `json:"times_escalados"`
	Fechamento struct{
		Dia int `json:"dia"`
		Mes int `json:"mes"`
		Ano int `json:"ano"`
		Hora int `json:"hora"`
		Minuto int `json:"minuto"`
		Timestamp int `json:"timestamp"`
	} `json:"fechamento"`
}

func(c *Status) GetStatus(){
	request:= Request{}

	res, err := request.Get(URI_STATUS, 15)
	if err != nil {
		panic(err)
	}
	if res.StatusCode() != 200 {
		panic("endpoint status com code")
	}

	err = json.Unmarshal(res.Body(), &c)
	if err != nil {
		panic(err)
	}
}
