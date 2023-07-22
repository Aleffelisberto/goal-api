package postgres

import "database/sql"

type dbManager struct {
	db *sql.DB
}

var DBManager *dbManager

func Init() {
	db, err := sql.Open("postgres", "postgres://goal_api:goal_api@127.0.0.1:5432/goal_api")
	if err != nil {
		panic(err)
	}

	DBManager = &dbManager{db: db}
}

func GetInstance() *sql.DB {
	return DBManager.db
}
