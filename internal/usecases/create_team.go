package usecases

import entity "goal-api/internal/entity"

type CreateTeamInputDto struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Country      string `json:"country"`
}

type CreateTeamUseCase struct {
	TeamRepository entity.TeamRepository
}

func MakeCreateTeamUseCase(teamRepository entity.TeamRepository) *CreateTeamUseCase {
	return &CreateTeamUseCase{TeamRepository: teamRepository}
}

func (useCase *CreateTeamUseCase) Execute(input CreateTeamInputDto) error {
	teamEntity := entity.NewTeamCreate(input.Name, input.Abbreviation, input.Country)

	err := useCase.TeamRepository.Create(teamEntity)
	if err != nil {
		return err
	}
	return nil
}
