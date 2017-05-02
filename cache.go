package kartola

import (
	"github.com/jhonata-menezes/kartola/api"
	"time"
)

var CacheKartolaAtletas api.Atletas
var CacheStatus api.Status
var TempoEspera = 60

func UpdateStatus() {
	status := api.Status{}
	status.GetStatus()
	CacheStatus = status
}

func UpdateCache() {
	UpdateStatus()

	time.Sleep(time.Duration(TempoEspera) * time.Second)
}
