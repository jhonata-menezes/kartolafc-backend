package main

import (
	"github.com/pressly/chi"
	"github.com/jhonata-menezes/kartola"
	"net/http"
	"fmt"
	"github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
)

func main(){
	go kartola.UpdateCache()

	router := chi.NewRouter()
	router.Use(middleware.DefaultCompress)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	kartola.BuildRoutes(router)
	fmt.Println("Bora Cumpade.")
	http.ListenAndServe("0.0.0.0:5015", router)
}
