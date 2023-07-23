package handlers_factory

import (
	"goal-api/internal/infra/repository/rgoqu"
	"goal-api/internal/infra/web"
	"goal-api/internal/usecases"

	"github.com/doug-martin/goqu"
)

func MakeTeamHandlers(db *goqu.Database) *web.TeamHandlers {
	teamRepository := rgoqu.NewTeamRepositoryGoqu(db)
	createTeamUseCase := usecases.MakeCreateTeamUseCase(teamRepository)
	teamHandlers := web.NewTeamHandlers(createTeamUseCase)

	return teamHandlers
}
