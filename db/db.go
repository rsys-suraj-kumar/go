package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDb() (*sql.DB,error){
	conStr:= os.Getenv("GOOSE_DBSTRING")

	db,err := sql.Open("postgres",conStr)

	if err != nil {
		log.Fatal("error: ",err.Error())
	}

	err2:= db.Ping()

	if err2 != nil {
		log.Fatal("dab is not connected")
	}

	fmt.Println("Sererv db")

	defer db.Close()
	
	return db,nil
}