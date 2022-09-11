package completeness

import (
	model "example.com/completude/pkg/application/completeness"
)

func (p *PersonCompleteness) ToModel() *[]model.Completeness {

	models := make([]model.Completeness, 0)

	if p.FullName.Validations != nil {

		model := p.FullName.ToModel()
		model.Type = 100
		model.PersonId = p.PersonId
		model.TenantId = p.TenantId

		models = append(models, model)
	}

	return &models
}

func (dto *Field) ToModel() model.Completeness {

	// convert validation list
	validations := make([]model.Validation, 0)
	for _, validation := range dto.Validations {
		validations = append(validations, validation.ToModel())
	}

	// completeness
	model := model.Completeness{
		Id:          dto.Id,
		Validations: validations,
	}

	return model
}

func (dto *Validation) ToModel() model.Validation {

	return model.Validation{
		Level:            dto.Level,
		Mode:             dto.Mode,
		Source:           dto.Source,
		ComplementText:   dto.Text,
		SourceCreateDate: dto.SourceCreateDate,
	}
}
