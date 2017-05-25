package main

import (
	"github.com/pressly/chi"
	kartolafc "github.com/jhonata-menezes/kartolafc-backend"
	"net/http"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/jhonata-menezes/kartolafc-backend/api"
)

func main() {

	go kartolafc.UpdateCache()

	// temporario, para criar pontuados mesmo com mercado aberto
	fileByte, err := ioutil.ReadFile("./pontuados.json")
	if err == nil {
		p := api.Pontuados{}
		json.Unmarshal(fileByte, &p)
		kartolafc.CachePontuados = p
	}
	//-------------------------------------------------------------------------------------

	router := chi.NewRouter()
	router.Use(middleware.DefaultCompress)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	kartolafc.BuildRoutes(router)
	log.Println("Bora Cumpade.")
	log.Println("listen", cmd.ServerBind)
	err = http.ListenAndServe(cmd.ServerBind, router)
	if err != nil {
		panic(err)
	}
}
