package completeness

import "fmt"

type Service interface {
	FindByPerson(personId string, tenantId string) (*[]Completeness, error)
}

type completenessService struct {
	repo Repository
}

func NewService(completenessRepository Repository) Service {

	return &completenessService{
		repo: completenessRepository,
	}
}

func (p *completenessService) FindByPerson(personId string, tenantId string) (*[]Completeness, error) {
	fmt.Printf("FindByPerson personId: %s, tenantId: %s\n", personId, tenantId)

	return p.repo.FindByPerson(personId, tenantId)
}
