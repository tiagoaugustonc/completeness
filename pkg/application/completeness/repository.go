package completeness

type Repository interface {
	Create(completeness Completeness) (*Completeness, error)
	Retrieve(id string) (*Completeness, error)
	FindByPerson(personId string, tenantId string) (*[]Completeness, error)
}
