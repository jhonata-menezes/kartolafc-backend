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
		Dia int
		Mes int
		Ano int
		Hora int
		Minuto int
		Timestamp int
	}
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
