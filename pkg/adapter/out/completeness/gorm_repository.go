package completeness

import (
	"fmt"

	model "example.com/completude/pkg/application/completeness"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormCompletenessRepository struct {
	gorm         *gorm.DB
	completeness []model.Completeness
}

func NewCompletenessRepository(gorm *gorm.DB) model.Repository {
	return &gormCompletenessRepository{gorm: gorm}
}

func (r *gormCompletenessRepository) Create(person model.Completeness) (*model.Completeness, error) {
	fmt.Printf("Salvando completeness %s\n", person)

	person.Id = uuid.NewString()

	r.completeness = append(r.completeness, person)

	return &person, nil
}

func (r *gormCompletenessRepository) Retrieve(id string) (*model.Completeness, error) {

	for _, person := range r.completeness {
		if person.Id == id {
			return &person, nil
		}
	}

	return nil, nil
}

func (r *gormCompletenessRepository) FindByPerson(personId string, tenantId string) (*[]model.Completeness, error) {

	// search for completeness
	var entitys []Completeness
	filter := Completeness{PersonId: personId, TenantId: tenantId}
	r.gorm.Find(&entitys, filter)

	// map to model
	var models = make([]model.Completeness, 0)
	for _, entity := range entitys {
		models = append(models, entity.toModel())
	}

	return &models, nil
}
