package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mattisongjj/Go-FAS-API/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Listening on localhost:8080...")
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
