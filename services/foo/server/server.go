package server

import (
	"fmt"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) Open(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Insecure foo\n")
	w.WriteHeader(http.StatusOK)

}

func (s Server) Secured(w http.ResponseWriter, req *http.Request) {
	val, ok := req.Header["Authorization"]
	if !ok || len(val) != 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if val[0] != "password123" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Secure foo!!!\n")
	w.WriteHeader(http.StatusOK)
}

func (s Server) Run() error {
	http.HandleFunc("/open", s.Open)
	http.HandleFunc("/secured", s.Secured)

	return http.ListenAndServe(":50051", nil)
}
