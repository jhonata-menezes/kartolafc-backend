package main

import (
	"github.com/pressly/chi"
	"github.com/jhonata-menezes/kartola"
	"net/http"
	"fmt"
)

func main(){
	go kartola.UpdateCache()

	router := chi.NewRouter()
	kartola.BuildRoutes(router)
	fmt.Println("Bora Cumpade.")
	http.ListenAndServe("0.0.0.0:5015", router)
}
