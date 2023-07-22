package handlers_factory

import (
	"database/sql"
	"goal-api/internal/infra/repository"
	"goal-api/internal/infra/web"
	"goal-api/internal/usecases"
)

func MakeTeamHandlers(db *sql.DB) *web.TeamHandlers {
	teamRepository := repository.NewTeamRepositoryMySql(db)
	createTeamUseCase := usecases.MakeCreateTeamUseCase(teamRepository)
	teamHandlers := web.NewTeamHandlers(createTeamUseCase)

	return teamHandlers
}
