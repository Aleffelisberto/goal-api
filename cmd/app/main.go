package main

import (
	"goal-api/cmd/app/config"
	"goal-api/internal/infra/db"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db.Init()
	defer db.GetDBInstance().Close()
	router := config.Router()
	go log.Fatal(http.ListenAndServe(":8000", router))
}
