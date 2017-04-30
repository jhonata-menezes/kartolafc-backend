package kartola

import (
	"github.com/jhonata-menezes/kartola/api"
	"net/http"
	"github.com/pressly/chi"
	"strconv"
)

func GetHome(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("{ \"status\":\"Birll\"}"))
}

func GetStatus(response http.ResponseWriter, request *http.Request) {
	apiStatus := api.Status{}
	apiStatus.GetStatus()

	response.Write(JsonBuild(apiStatus))
}

func GetTimes(response http.ResponseWriter, request *http.Request) {
	timePesquisado := chi.URLParam(request, "q")
	times := api.Times{}
	times.Pesquisa = timePesquisado
	times.GetTimes()

	response.Write(JsonBuild(times))
}

func GetTime(response http.ResponseWriter, request *http.Request) {
	idString := chi.URLParam(request, "id")
	time := api.TimeCompleto{}

	id, err := strconv.Atoi(idString)

	if err != nil {
		response.Write([]byte("{\"status\": \"error\", \"message\": \"id tem que ser um numero\"}"))
	} else {
		time.TimeCompleto.TimeId = id
		time.GetTime()
		response.Write(JsonBuild(time))
	}
}

func GetMercado(response http.ResponseWriter, request *http.Request) {
	atletas := api.Atletas{}
	atletas.GetAtletas()

	response.Write(JsonBuild(atletas))
}