package config

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"goal-api/internal/infra/db"
	"goal-api/internal/infra/web/handlers_factory"

	_ "goal-api/docs"

	_ "github.com/lib/pq"
)

func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Goal API version: 1.0.0"))
}

type TeamCreate struct {
	Name         string `json:"name" example:"Nome da equipe"`
	Abbreviation string `json:"abbreviation" example:"ABCD"`
	Country      string `json:"country" example:"Brasil"`
}

type Team struct {
	ID           int    `json:"id" example:"1"`
	Name         string `json:"name" example:"Nome da equipe"`
	Abbreviation string `json:"abbreviation" example:"ABCD"`
	Country      string `json:"country" example:"Brasil"`
	CreatedAt    string `json:"created_at" example:"2023/07/22 22:35:22"`
	UpdatedAt    string `json:"updated_at" example:"2023/07/22 22:35:22"`
}

// CreateTeam
//
//	@Summary		Create a team
//	@Description	Create a team
//	@Tags			Team
//	@Accept			json
//	@Produce		json
//	@Param			team	body	TeamCreate	true	"Informações do time a ser criado"
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/teams [POST]
func CreateTeam(w http.ResponseWriter, r *http.Request) {
	teamHandlers := handlers_factory.MakeTeamHandlers(db.GetGoquInstance())
	teamHandlers.CreateTeamHandler(w, r)
}

// ListTeams
//
// @Summary		List teams
// @Description	Get a list of teams
// @Tags			Team
// @Accept		json
// @Produce		json
// @Success		200	{array} Team
// @Failure		500
// @Router		/teams [GET]
func ListTeams(w http.ResponseWriter, r *http.Request) {
	teamHandlers := handlers_factory.MakeTeamHandlers(db.GetGoquInstance())
	teamHandlers.ListTeamsHandler(w, r)
}

// @title			Goal API
// @version		1.0
// @description	Goal API
// @termsOfService	http://swagger.io/terms/
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath	/
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", version)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/teams", ListTeams).Methods("GET")
	router.HandleFunc("/teams", CreateTeam).Methods("POST")

	return router
}
