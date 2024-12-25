package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServe(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	//? create router
	router := mux.NewRouter()

	//? register sservices

	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, router)

}
