package api

import (
	"encoding/json"
	"log"
	"io/ioutil"
)

const URI_STATUS = "/mercado/status"

type Status struct {
	RodadaAtual int `json:"rodada_atual"`
	StatusMercado int `json:"status_mercado"`
	EsquemaDefaultId int `json:"esquema_default_id"`
	TimesEscalados int `json:"times_escalados"`
	GameOver bool `json:"game_over"`
	MercadoPosRadada bool `json:"mercado_pos_radada"`
	Reativar bool `json:"reativar"`
	Aviso string `json:"aviso"`
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
		log.Println(err)
	}else{
		if res.StatusCode != 200 {
			log.Println("endpoint status com code")
		}else{
			by, _ := ioutil.ReadAll(res.Body)
			err = json.Unmarshal(by, &c)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
