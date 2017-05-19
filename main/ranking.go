package main

import (
	"gopkg.in/mgo.v2"
	"time"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
	"net/http"
	"log"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"github.com/jhonata-menezes/kartolafc-backend"
	"strconv"
	"encoding/json"
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetSocketTimeout(2 * time.Hour)
	session.SetMode(mgo.Monotonic, true)
	colllection := session.DB("kartolafc").C("times")

	kartolafc.LoadInMemory(colllection)

	router := chi.NewRouter()
	router.Use(middleware.DefaultCompress)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	router.Get("/ranking/melhores", GetMelhoresRanking)
	router.Get("/ranking/time/id/:id", GetRankingTimeId)

	log.Println("listen", cmd.ServerBind)
	err = http.ListenAndServe(cmd.ServerBind, router)
	if err != nil {
		panic(err)
	}
}

func GetMelhoresRanking(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	render.JSON(response, request, kartolafc.CacheRankingPontuadosMelhores)
}

func GetRankingTimeId(response http.ResponseWriter, request *http.Request) {
	responseDefault(response)
	timeId, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil ||
		timeId < 1 ||
		timeId >= 15000000 ||
		kartolafc.CacheRankingTimeIdPontuados[timeId].TimeId == 0{

		var out map[string]string
		json.Unmarshal([]byte("{ \"status\":\"error\", \"message\":\"id informado nao existe\"}"), &out)
		render.JSON(response, request, out)
	}else{
		render.JSON(response, request, kartolafc.CacheRankingTimeIdPontuados[timeId])
	}
}

func responseDefault(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}