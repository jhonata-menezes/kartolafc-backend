package main

import (
	"github.com/jhonata-menezes/kartolafc-backend"
	"log"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"strconv"
)

type PontuadoSimples struct {
	Apelido string `json:"apelido"`
	Pontuacao float32 `json:"pontuacao"`
	Scout map[string]int `json:"scout"`
	PosicaoId int `json:"posicao_id"`
	ClubeId int `json:"clube_id"`
}

type PontuadosSimples struct {
	Rodada int `json:"rodada"`
	Time int64 `json:"time"`
	Atletas map[string]PontuadoSimples `json:"atletas"`
}

func main() {
	var updated = false
	go kartolafc.UpdatePontuados() //apenas os pontuados

	session, err:= mgo.Dial(cmd.MongoDBIpPort)

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetSocketTimeout(10 * time.Minute)
	session.SetMode(mgo.Monotonic, true)

	collectionRodada := session.DB("kartolafc").C("pontuacao_rodada")

	pontuados := api.Pontuados{}
	pontuados.GetPontuados()
	save(pontuados, collectionRodada)

	log.Println("Bora Cumpade.")


	for {
		pontuados,updated = Update(pontuados)
		if updated {
			save(pontuados, collectionRodada)
		}
		time.Sleep(10 * time.Second)
	}
}


func Update(pontuadosOld api.Pontuados) (api.Pontuados, bool) {
	r:= api.Pontuados{}
	r.GetPontuados()

	if r.Rodada != 0 {
		if pontuadosSaoIguais(pontuadosOld, r) {
			log.Println("são iguais")
		}else{
			log.Println("nao sao iguais")
			return r, true //se for diferente entao considerar esse valores
		}
	}else{
		log.Println("nao retornou nada")
	}
	return pontuadosOld, false
}

func pontuadosSaoIguais(old, new api.Pontuados) bool {
	for atletaId, desc := range new.Atletas {
		//log.Println("id", atletaId, "old", desc.Pontuacao, "new", new.AtletasRanking[atletaId].Pontuacao)
		if (old.Atletas[atletaId].Pontuacao == desc.Pontuacao) {
			continue
		}
		return false
	}
	return true
}

func save(p api.Pontuados, c *mgo.Collection) {
	timeunix := time.Now().Unix()
	newPontuados := PontuadosSimples{}

	newPontuados.Atletas = map[string]PontuadoSimples{}
	newPontuados.Rodada = p.Rodada
	newPontuados.Time = timeunix
	for atletaId, desc := range p.Atletas {
		auxPontuados := PontuadoSimples{}
		auxPontuados.Pontuacao = desc.Pontuacao
		auxPontuados.Apelido = desc.Apelido
		auxPontuados.Scout = desc.Scout
		auxPontuados.ClubeId = desc.ClubeId
		auxPontuados.PosicaoId = desc.PosicaoId
		newPontuados.Atletas[strconv.Itoa(atletaId)] = auxPontuados
	}

	//fmt.Printf("%#v", newPontuados)

	if err := c.Insert(newPontuados); err != nil {
		log.Println("nao possivel inserir no mongodb os pontuados")
	}

}
