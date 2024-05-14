package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skradiansys/go/db"
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
	_,dbError := db.NewDb()

	if dbError != nil {
		log.Fatal("Something wrong with the db")
	}

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