package kartolafc

import (
	"github.com/jhonata-menezes/kartolafc-backend/api"
	"time"
)

var CacheKartolaAtletas api.Atletas
var CacheStatus api.Status
var CacheDestaques api.Destaques


func UpdateStatus() {
	status := api.Status{}
	status.GetStatus()
	CacheStatus = status
	SleepCacheSecond(60)
}

func UpdateDestaques() {
	destaques := api.Destaques{}
	destaques.GetDestaques()
	CacheDestaques = destaques

	SleepCacheSecond(60)
}

func UpdateMercado() {
	mercado := api.Atletas{}
	mercado.GetAtletas()
	CacheKartolaAtletas = mercado

	SleepCacheSecond(60)
}

func UpdateCache() {
	go UpdateStatus()
	go UpdateDestaques()
	go UpdateMercado()
}

func SleepCacheSecond(t int) {
	time.Sleep(time.Duration(t) * time.Second)
}
