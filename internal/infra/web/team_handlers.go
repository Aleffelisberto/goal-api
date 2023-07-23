package web

import (
	"encoding/json"
	"goal-api/internal/usecases"
	"log"
	"net/http"
)

type TeamHandlers struct {
	CreateTeamUseCase *usecases.CreateTeamUseCase
	ListTeamsUseCase  *usecases.ListTeamsUseCase
}

func NewTeamHandlers(createTeamUseCase *usecases.CreateTeamUseCase, listTeamsUseCase *usecases.ListTeamsUseCase) *TeamHandlers {
	return &TeamHandlers{
		CreateTeamUseCase: createTeamUseCase,
		ListTeamsUseCase:  listTeamsUseCase,
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

func (handler *TeamHandlers) ListTeamsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := handler.ListTeamsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}
