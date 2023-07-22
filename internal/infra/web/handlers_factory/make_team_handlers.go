package handlers_factory

import (
	"database/sql"
	"fmt"
	"goal-api/internal/infra/repository"
	"goal-api/internal/infra/web"
	"goal-api/internal/usecases"
)

func MakeTeamHandlers(db *sql.DB) *web.TeamHandlers {
	teamRepository := repository.NewTeamRepositoryMySql(db)
	createTeamUseCase := usecases.MakeCreateTeamUseCase(teamRepository)
	teamHandlers := web.NewTeamHandlers(createTeamUseCase)

	fmt.Println("Has passed hereeee")

	return teamHandlers
}
