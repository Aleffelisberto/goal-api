package entity

type Team struct {
	ID           int
	Name         string
	Abbreviation string
	Country      string
	CreatedAt    string
	UpdatedAt    string
}

type TeamCreate struct {
	Name         string
	Abbreviation string
	Country      string
}

type TeamRepository interface {
	Create(team *TeamCreate) error
	FindAll() ([]Team, error)
}

func NewTeamCreate(name string, abbreviation string, country string) *TeamCreate {
	return &TeamCreate{
		Name:         name,
		Abbreviation: abbreviation,
		Country:      country,
	}
}
