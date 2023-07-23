package usecases

import "goal-api/internal/entity"

type ListTeamsOutputDto struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Country      string `json:"country"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type ListTeamsUseCase struct {
	TeamRepository entity.TeamRepository
}

func MakeListTeamsUseCase(teamRepository entity.TeamRepository) *ListTeamsUseCase {
	return &ListTeamsUseCase{TeamRepository: teamRepository}
}

func (useCase *ListTeamsUseCase) Execute() ([]*ListTeamsOutputDto, error) {
	teams, err := useCase.TeamRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var teamsOutput []*ListTeamsOutputDto
	for _, team := range teams {
		teamsOutput = append(teamsOutput, &ListTeamsOutputDto{
			ID:           team.ID,
			Name:         team.Name,
			Abbreviation: team.Abbreviation,
			Country:      team.Country,
			CreatedAt:    team.CreatedAt,
			UpdatedAt:    team.UpdatedAt,
		})
	}
	return teamsOutput, err
}
