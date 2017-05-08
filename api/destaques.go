package api

import (
	"log"
	"encoding/json"
	"strings"
)

const URL_DESTAQUES = "/mercado/destaques"

type Destaque struct {
	Atleta struct{
		Atleta_id int `json:"atleta_id"`
		Nome string `json:"nome"`
		Apelido string `json:"apelido"`
		Foto string `json:"foto"`
		PrecoEditorial int `json:"preco_editorial"`
	} `json:"Atleta"`
	Escalacoes int `json:"escalacoes"`
	Clube string `json:"clube"`
	EscudoClube string `json:"escudo_clube"`
	Posicao string `json:"posicao"`
}

type Destaques []Destaque

func (d *Destaques) GetDestaques() {
	request := Request{}
	res, err := request.Get(URL_DESTAQUES, 10 )
	if err != nil {
		log.Println("destaques", err)
	}
	if res.StatusCode() != 200 {
		log.Println("destaques diferente de 200 status", res.StatusCode())
	}else{
		json.Unmarshal(res.Body(), &d)
		d.ChangeFormatDefault()
	}
}

func (d *Destaques) ChangeFormatDefault() {
	for i, des := range *d {
		(*d)[i].Atleta.Foto = strings.Replace(des.Atleta.Foto, "FORMATO", "140x140", 3)
	}
}
