package api

import (
	"strconv"
	"github.com/valyala/fasthttp"
	"encoding/json"
	"log"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
)

const URL_PARTIDAS = "/partidas"
const URL_PARTIDAS_FUTURA = "https://globoesporte.globo.com/servico/backstage/esportes_campeonato/esporte/futebol/modalidade/futebol_de_campo/categoria/profissional/campeonato/campeonato-brasileiro/edicao/campeonato-brasileiro-2018/fases/fase-unica-seriea-2018/rodada/%d/jogos.html"

type Partida struct {
	ConfrontoNome string `json:"confronto_nome"`
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
	if rodada >= 1 && rodada <= 38 {
		rodadaParse := URL_PARTIDAS + "/" + strconv.Itoa(rodada)
		res, err = request.Get(rodadaParse, 10)
	} else {
		res, err = request.Get(URL_PARTIDAS, 10)
	}
	if err != nil || res.StatusCode() != 200 {
		return
	}
	if err = json.Unmarshal(res.Body(), &p); err != nil {
		log.Println("nao decodificou o json da partida", rodada)
	}
}

func (p *Partidas)GetFuturas(rodada int) {
	end := p.mountUrlFuturas(rodada)
	r := Request{}
	resp, err := r.GetSimple(end, 10)
	if err != nil {
		log.Println(err)
		return
	}
	node, err := html.Parse(bytes.NewReader(resp.Body()))
	if err != nil {
		log.Println(err)
		return
	}

	partidas := parse(node.FirstChild.LastChild.FirstChild)
	p.Rodada = rodada
	p.Partidas = partidas.Partidas
}

func (p *Partidas)mountUrlFuturas(r int) string {
	return fmt.Sprintf(URL_PARTIDAS_FUTURA, r)
}

func parse(n *html.Node) Partidas {
	f := n
	partidas := Partidas{}
	partidas.Partidas = make([]Partida, 0)
	for {
		partida := Partida{}
		n := f.FirstChild.FirstChild
		partida.ConfrontoNome = n.Attr[1].Val
		hora := ""
		if n.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling != nil {
			hora = "às " + n.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.Data
		}
		partida.PartidaData = n.NextSibling.NextSibling.FirstChild.Data + hora
		local:=""
		if (n.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild != nil) {
			local = n.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
		}
		partida.Local = local
		partidas.Partidas = append(partidas.Partidas, partida)
		if f = f.NextSibling; f.NextSibling == nil {
			break
		}
	}
	return partidas
}