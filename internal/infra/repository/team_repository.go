package repository

import (
	"database/sql"
	"goal-api/internal/entity"
)

type TeamRepositoryMySql struct {
	DB *sql.DB
}

func NewTeamRepositoryMySql(db *sql.DB) *TeamRepositoryMySql {
	return &TeamRepositoryMySql{DB: db}
}

func (repo *TeamRepositoryMySql) Create(team *entity.TeamCreate) error {
	_, err := repo.DB.Exec("INSERT INTO teams (name, abbreviation, country) values ($1 ,$2 ,$3)", team.Name, team.Abbreviation, team.Country)

	if err != nil {
		return err
	}
	return nil
}
