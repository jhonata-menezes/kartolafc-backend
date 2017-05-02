package kartola

import (
	"github.com/jhonata-menezes/kartola/api"
	"net/http"
	"github.com/pressly/chi"
	"strconv"
)

func GetHome(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	response.Write([]byte("{ \"status\":\"Birll\"}"))
}

func GetStatus(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)

	response.Write(JsonBuild(CacheStatus))
}

func GetTimes(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	timePesquisado := chi.URLParam(request, "q")
	times := api.Times{}
	times.Pesquisa = timePesquisado
	times.GetTimes()

	response.Write(JsonBuild(times))
}

func GetTime(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
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
	responseDefault(response)
	atletas := api.Atletas{}
	atletas.GetAtletas()

	response.Write(JsonBuild(atletas))
}

func responseDefault(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}