package api

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"log"
)

const URL_PONTUADOS = "/atletas/pontuados"

type Pontuado struct {
	Apelido string `json:"apelido"`
	Pontuacao float32 `json:"pontuacao"`
	Scout map[string]int `json:"scout"`
	Foto string `json:"foto"`
	PosicaoId int `json:"posicao_id"`
	ClubeId int `json:"clube_id"`
}

type Clube struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Abreviacao string `json:"abreviacao"`
	Escudos map[string]string
}

type Posicoes struct {
	Id int `json:"id"`
	Nome string `json:"nome"`
	Abreviacao string `json:"abreviacao"`
}

type Pontuados struct {
	Rodada int `json:"rodada"`
	Atletas map[int]Pontuado `json:"atletas"`
	Clubes map[int]Clube `json:"clubes"`
	Posicoes map[int]Posicoes `json:"posicoes"`
	TotalAtletas int `json:"total_atletas"`
}

func (p *Pontuados) GetPontuados() {
	file, err := ioutil.ReadFile("/home/jhonata/.gopath/src/github.com/jhonata-menezes/kartolafc-backend/mock/pontuados.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file,&p); err == nil {
		p.ChangeFormatDefault()
	}

}

func (p *Pontuados) ChangeFormatDefault() {
	for i, des := range p.Atletas {
		foto := p.Atletas[i]
		foto.Foto = strings.Replace(des.Foto, "FORMATO", "140x140", 3)
		p.Atletas[i] = foto
		log.Println(strings.Replace(des.Foto, "FORMATO", "140x140", 3))
	}
}