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

func (repo *TeamRepositoryGoqu) FindAll() ([]entity.Team, error) {
	query := goqu.Select().From("teams")

	sql, _, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := repo.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []entity.Team
	for rows.Next() {
		var team entity.Team
		err := rows.Scan(&team.ID, &team.Name, &team.Abbreviation, &team.Country, &team.CreatedAt, &team.UpdatedAt)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}
