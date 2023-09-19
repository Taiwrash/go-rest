package main

import (
	"net/http"

	"github.com/Taiwrash/go-rest/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	println("Starting GO API services")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Error(err)
	}
}
