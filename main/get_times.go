package main

import (
	"gopkg.in/mgo.v2"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"sync"
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"log"
	"time"
	"encoding/json"
	"fmt"
	"strings"
)

type MeuTime api.TimeCompleto

func main() {
	session, err := mgo.Dial(cmd.MongoDBIpPort)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var wg = sync.WaitGroup{}
	chIdTime := make(chan int)

	for g:=0; g<20; g++ {
		collection := session.Copy().DB("kartolafc").C("times")
		go getTime(&chIdTime, &wg, collection)
	}

	inicio := time.Now()
	//hora de lancar os id's para processar
	log.Println("enviando id's")
	for g:=1; g<1000; g++ {
		chIdTime <- g
	}

	for {
		cnt := len(chIdTime)
		log.Println("contagem fila", cnt)
		if cnt  == 0 {
			break
		}
		time.Sleep(10 * time.Second)
	}
	log.Println("fim", time.Since(inicio))
}


func getTime(time *chan int, wg *sync.WaitGroup, c *mgo.Collection) {
	defer wg.Done()
	tentativas := 0
	for t := range *time {
		timeApi := MeuTime{}
		timeApi.TimeCompleto.TimeId = t
		statusCode := timeApi.GetTime()

		if statusCode == 500 {
			*time <- t
		}

		if tentativas >= 100 {
			log.Println("tentativas", tentativas, "excedidas parando a goroutine")
		}

		if timeApi.RodadaAtual <= 0 {
			tentativas++
			continue
		}

		err := c.Insert(timeApi)
		if err != nil {
			log.Println("nao foi possivel inserir", err)
		}
		tentativas = 0
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