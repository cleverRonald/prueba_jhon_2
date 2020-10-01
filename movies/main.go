package main

import (
	"net/http"
	"os"
	"github.com/cleverRonald/prueba_jhon_2/common"
	"github.com/cleverRonald/prueba_jhon_2/routers"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Jhon CTMR vamos a jalar ptmr!!!!!!!!!!!!!!!!!!!!!!!!!!!!!</h1>"))
}

func main() {
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}