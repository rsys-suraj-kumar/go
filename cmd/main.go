package main

import (
	"fmt"
	"log"

	"github.com/skradiansys/go/server"
)

func main() {
	fmt.Println("This is the entry point of this file.")

	server := server.NewApiServer(":8000")

	err := server.Run()

	if err != nil {
		log.Fatal("There was a error runing",err)
	}
}