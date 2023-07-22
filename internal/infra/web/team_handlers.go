package web

import (
	"encoding/json"
	"goal-api/internal/usecases"
	"log"
	"net/http"
)

type TeamHandlers struct {
	CreateTeamUseCase *usecases.CreateTeamUseCase
}

func NewTeamHandlers(createTeamUseCase *usecases.CreateTeamUseCase) *TeamHandlers {
	return &TeamHandlers{
		CreateTeamUseCase: createTeamUseCase,
	}
}

func (handler *TeamHandlers) CreateTeamHandler(w http.ResponseWriter, r *http.Request) {
	var input usecases.CreateTeamInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.CreateTeamUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
