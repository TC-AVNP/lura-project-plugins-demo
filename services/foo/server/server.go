package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) Open(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Insecure foo\n")
	w.WriteHeader(http.StatusOK)

}

func (s Server) Sleepy(w http.ResponseWriter, req *http.Request) {
	fmt.Println("will respond in 10 seconds")

	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "Finished sleeping\n")
	w.WriteHeader(http.StatusOK)

}

func (s Server) Secured(w http.ResponseWriter, req *http.Request) {
	val, ok := req.Header["Authorization"]
	if !ok || len(val) != 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if val[0] != "password123!" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Secure foo!!!\n")
	w.WriteHeader(http.StatusOK)
}

func (s Server) Run() error {
	http.HandleFunc("/open", s.Open)
	http.HandleFunc("/secured", s.Secured)
	http.HandleFunc("/sleepy", s.Sleepy)

	return http.ListenAndServe(":50051", nil)
}
