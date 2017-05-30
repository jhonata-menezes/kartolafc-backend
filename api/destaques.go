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
		PrecoEditorial float32 `json:"preco_editorial"`
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
		if err = json.Unmarshal(res.Body(), &d); err != nil {
			log.Println("json destaques", err)
		}else{
			d.ChangeFormatDefault()
		}
	}
}

func (d *Destaques) ChangeFormatDefault() {
	for i, des := range *d {
		(*d)[i].Atleta.Foto = strings.Replace(des.Atleta.Foto, "FORMATO", "140x140", 3)
	}
}
