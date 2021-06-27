package http

import (
	"api/interface/controller"
	"api/util"
	"net/http"

	"github.com/gorilla/mux"
)

type Routing struct {
	Controllers *controller.AppController
	Port        string
	Router      *mux.Router
}

func NewRouting(ctrls *controller.AppController, port string) *Routing {
	return &Routing{
		Controllers: ctrls,
		Port:        port,
		Router:      mux.NewRouter(),
	}
}

func (r *Routing) SetRouting() {
	logger := util.NewStdLogger()

	b := r.Router.PathPrefix("/go_auth").Subrouter()

	b.HandleFunc("/health_check", r.Controllers.HealthCheck.HealthCheck).Methods(http.MethodGet)

	logger.Printf("%s", "Mux Routers Start.")
	logger.Fatalf("%s", http.ListenAndServe(":"+r.Port, b))
}
