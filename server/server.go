package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skradiansys/go/db"
	"github.com/skradiansys/go/internals/user"
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
	postgresDb,dbError := db.NewDb()

	if dbError != nil {
		log.Fatal("Something wrong with the db")
	}

	router := http.NewServeMux()

	userStore := user.NewStore(postgresDb)
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(router)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/",http.StripPrefix("/v1",router))

	server:= http.Server{
		Addr: s.addr,
		Handler: router,
	}

	fmt.Println("Server is runing on port: ",s.addr);

	return server.ListenAndServe()
}