package server

import (
	"context"
	"net/http"
	"strconv"
)

type Server struct {
	instances map[int]context.CancelFunc
}

func NewServer() Server {
	return Server{
		instances: make(map[int]context.CancelFunc),
	}
}

func (s Server) StartOne() error {
	lura := NewLuraInstance()
	ctxCancel, err := lura.Spawn(8080)
	if err != nil {
		return err
	}

	s.instances[8080] = ctxCancel
	return nil
}

func (s Server) StartLura(w http.ResponseWriter, req *http.Request) {
	val, ok := req.URL.Query()["port"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}
	port, _ := strconv.Atoi(val[0])
	lura := NewLuraInstance()
	ctxCancel, err := lura.Spawn(port)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	s.instances[port] = ctxCancel
	w.WriteHeader(http.StatusOK)
}

func (s Server) StopLura(w http.ResponseWriter, req *http.Request) {
	val, ok := req.URL.Query()["port"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}
	port, _ := strconv.Atoi(val[0])

	s.instances[port]()
	w.WriteHeader(http.StatusOK)
}

func (s Server) Run() error {
	http.HandleFunc("/start", s.StartLura)
	http.HandleFunc("/stop", s.StopLura)

	return http.ListenAndServe(":8083", nil)
}
