package main

import (
	"fmt"
	"github.com/benmanns/goworker"
	"os"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"log"
	"encoding/json"
	"strings"
	"gopkg.in/mgo.v2"
)

type MeuTime api.TimeCompleto

var channelConnectionMongodb = make(chan *mgo.Collection, 200)

func init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://" + os.Args[1] + ":6379/",
		Connections:    200,
		Queues:         []string{"times"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    200,
		Namespace:      "kartolafc:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("Id", run)

	session, err := mgo.Dial(os.Args[1] + ":27017")
	if err != nil {
		panic(err)
	}

	// gera pool de conexoes no channel
	for i:=0; i<200; i++ {
		channelConnectionMongodb <- session.Copy().DB("kartolafc").C("times")
	}
}

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}

func run(queue string, args ...interface{}) error {
	id, _ := args[0].(json.Number)
	timeId64, _ := id.Int64()
	timeId := int(timeId64)
	log.Println("Processando timeId", timeId)
	collection := <- channelConnectionMongodb
	getTime(timeId, collection)
	channelConnectionMongodb <- collection
	return nil
}

func getTime(timeId int, c *mgo.Collection) {
	timeApi := MeuTime{}
	timeApi.TimeCompleto.TimeId = timeId
	statusCode := timeApi.GetTime()
	log.Println("time id", timeId, "status code", statusCode)
	if statusCode == 500 {
		log.Println("id", timeId, "status 500 enviando para a fila")
		enqueue(timeId)
		return
	}

	if timeApi.Mensagem == "Este time ainda não foi escalado na temporada." {
		return
	}

	err := c.Insert(timeApi)
	if err != nil {
		log.Println("nao foi possivel inserir", err)
	}else{
		log.Println("id", timeId, "salvo no banco")
	}
}


func (t *MeuTime) GetTime() int {
	request := api.Request{}
	status := 0

	for i:=1; i<=8; i++ {
		res, err := request.Get(t.MountUrl(), 10)
		status = res.StatusCode()
		if err != nil {
			log.Println(err)
		}
		if res.StatusCode() != 200 {
			log.Println("tentativa", i,"time id", t.TimeCompleto.TimeId, "diferente de 200", res.StatusCode())
		}else{
			json.Unmarshal(res.Body(), &t)
			t.ChangeFormatDefault()
		}

		if res.StatusCode() == 500 {
			continue
		}else{
			break
		}
	}
	return status
}

func (t *MeuTime) MountUrl() string {
	return fmt.Sprintf(api.URL_TIME_ID, t.TimeCompleto.TimeId)
}

func (a *MeuTime) ChangeFormatDefault() {
	for i, des := range a.Atletas {
		a.Atletas[i].Foto = strings.Replace(des.Foto, "FORMATO", "140x140", 3)
	}
}

func enqueue(timeId int) {
	goworker.Enqueue(&goworker.Job{
		Queue: "times",
		Payload: goworker.Payload{
			Class: "Id",
			Args: []interface{}{timeId},
		},
	})
}