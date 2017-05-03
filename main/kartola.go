package main

import (
	"github.com/pressly/chi"
	kartolafc "github.com/jhonata-menezes/kartolafc-backend"
	"net/http"
	"fmt"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
	"github.com/jhonata-menezes/kartolafc-backend/cmd"
	"log"
)

func main() {

	go kartolafc.UpdateCache()

	router := chi.NewRouter()
	router.Use(middleware.DefaultCompress)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	kartolafc.BuildRoutes(router)
	fmt.Println("Bora Cumpade.")
	log.Println("listen", cmd.ServerBind)
	err := http.ListenAndServe(cmd.ServerBind, router)
	if err != nil {
		panic(err)
	}
}
