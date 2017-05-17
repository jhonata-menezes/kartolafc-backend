package main

import (
	"gopkg.in/mgo.v2"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sort"
	"time"
	"strconv"
	"os"
)

type Atleta struct {
	AtletaId int `json:"atleta_id"`
}

type Atletas struct {
	Pontuacao float32
	Atletas []Atleta `bson:"atletas"`
	TimeCompleto struct{
		TimeId int
	} `bson:"timecompleto"`
}

type Times []Atletas

func (a Times) Len() int {
	return len(a)
}

func (a Times) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Times) Less(i, j int) bool {
	time1 := SomaPontuacao(a[i])
	time2 := SomaPontuacao(a[j])
	a[i].Pontuacao = time1
	a[j].Pontuacao = time2
	return  time1 > time2
}

func SomaPontuacao(atletasTime Atletas) float32 {
	var soma float32
	for _, a := range atletasTime.Atletas {
		soma+= Pontuados.Atletas[a.AtletaId].Pontuacao
	}
	return soma
}

var Pontuados api.Pontuados

func main() {
	limit, _:= strconv.Atoi(os.Args[1])
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetSocketTimeout(2 * time.Hour)
	session.SetMode(mgo.Monotonic, true)
	colllection := session.DB("kartolafc").C("times")

	inicio := time.Now()
	Pontuado := api.Pontuados{}
	Pontuado.GetPontuados()

	Pontuados = Pontuado

	var atl Times
	err = colllection.Find(bson.M{}).Limit(limit).All(&atl)

	if err != nil {
		panic(err)
	}
	log.Println("pegar todos times da collection", time.Since(inicio))

	inicio = time.Now()

	sort.Sort(atl)

	log.Println("sort", time.Since(inicio))

	log.Printf("exibindo registro 0 ordenado %#v \n", atl[0])

}