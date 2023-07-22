package usecases

import entity "goal-api/internal/entity"

type CreateTeamInputDto struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Country      string `json:"country"`
}

// type CreateTeamOutputDto struct {
// 	ID           int
// 	Name         string
// 	Abbreviation string
// 	Country      string
// 	CreatedAt    string
// 	UpdatedAt    string
// }

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
	// return &CreateTeamOutputDto{
	// 	ID:           team.ID,
	// 	Name:         team.Name,
	// 	Abbreviation: team.Abbreviation,
	// 	Country:      team.Country,
	// 	CreatedAt:    team.CreatedAt,
	// 	UpdatedAt:    team.UpdatedAt,
	// }, nil
	return nil
}
