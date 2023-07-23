package rgoqu

import (
	"goal-api/internal/entity"

	"github.com/doug-martin/goqu"
)

type TeamRepositoryGoqu struct {
	DB *goqu.Database
}

func NewTeamRepositoryGoqu(db *goqu.Database) *TeamRepositoryGoqu {
	return &TeamRepositoryGoqu{DB: db}
}

func (repo *TeamRepositoryGoqu) Create(team *entity.TeamCreate) error {
	query := goqu.Insert("teams").
		Rows(goqu.Record{
			"name":         team.Name,
			"abbreviation": team.Abbreviation,
			"country":      team.Country,
		})

	sql, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}
