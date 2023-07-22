package config

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"goal-api/internal/infra/db/postgres"
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
	teamHandlers := handlers_factory.MakeTeamHandlers(postgres.GetInstance())
	teamHandlers.CreateTeamHandler(w, r)
}

//	@title			Goal API
//	@version		1.0
//	@description	Goal API
//	@termsOfService	http://swagger.io/terms/
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host
//	@BasePath	/
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", version)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/teams", CreateTeam).Methods("POST")

	return router
}
