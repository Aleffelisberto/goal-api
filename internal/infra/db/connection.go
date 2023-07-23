package db

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

type dbManager struct {
	db   *sql.DB
	goqu *goqu.Database
}

var DBManager *dbManager

func Init() {
	dialect := goqu.Dialect("postgres")
	pgDb, err := sql.Open("postgres", "postgres://goal_api:goal_api@127.0.0.1:5432/goal_api")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("✔️ Postgres connection established")

	db := dialect.DB(pgDb)

	DBManager = &dbManager{db: pgDb, goqu: db}
}

func GetDBInstance() *sql.DB {
	return DBManager.db
}

func GetGoquInstance() *goqu.Database {
	return DBManager.goqu
}
