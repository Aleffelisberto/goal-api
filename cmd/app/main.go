package main

import (
	"goal-api/cmd/app/config"
	"goal-api/internal/infra/db/postgres"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	postgres.Init()
	defer postgres.GetInstance().Close()
	router := config.Router()
	go log.Fatal(http.ListenAndServe(":8000", router))
}
