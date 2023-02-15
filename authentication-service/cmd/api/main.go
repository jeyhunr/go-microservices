package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jeyhunr/go-microservices/authentication-service/data"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	fmt.Println("Startign authentication service")

	app := Config{}

	// http server definition
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// starting the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
