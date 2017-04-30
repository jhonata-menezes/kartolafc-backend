package main

import (
	"github.com/pressly/chi"
	"github.com/jhonata-menezes/kartola"
	"net/http"
	"fmt"
)

func main(){
	router := chi.NewRouter()
	kartola.BuildRoutes(router)

	fmt.Println("Bora Cumpade.")
	http.ListenAndServe(":5015", router)
}
