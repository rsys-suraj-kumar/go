package server

import (
	"fmt"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working"))
	})

	server:= http.Server{
		Addr: s.addr,
		Handler: router,
	}

	fmt.Println("Server is runing on port: ",s.addr);

	return server.ListenAndServe()
}