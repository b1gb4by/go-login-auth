package http

import (
	"api/interface/controller"
	"api/util"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	l := r.Router.PathPrefix("/auth").Subrouter()

	l.HandleFunc("/register", r.Controllers.RegisterUser.RegisterUser).Methods(http.MethodPost)
	l.HandleFunc("/login", r.Controllers.LoginAuthentication.LoginAuthentication).Methods(http.MethodPost)
	l.HandleFunc("/logout", r.Controllers.Logout.Logout).Methods(http.MethodGet)
	l.HandleFunc("/user", r.Controllers.UserAuthentication.UserAuthentication).Methods(http.MethodGet)
	l.HandleFunc("/forgot", r.Controllers.Forgot.Forgot).Methods(http.MethodPost)
	l.HandleFunc("/reset/password", r.Controllers.Reset.Reset).Methods(http.MethodPost)
	l.HandleFunc("/health_check", r.Controllers.HealthCheck.HealthCheck).Methods(http.MethodGet)

	c := cors.Default().Handler(l)

	logger.Printf("%s", "Mux Routers Start.")
	logger.Fatalf("%s", http.ListenAndServe(":"+r.Port, c))
}
