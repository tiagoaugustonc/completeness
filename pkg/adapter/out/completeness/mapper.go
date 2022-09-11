package completeness

import (
	model "example.com/completude/pkg/application/completeness"
)

func (e *Completeness) toModel() model.Completeness {

	var validations []model.Validation
	for _, val := range e.Validations {
		validations = append(validations, val.toModel())
	}

	return model.Completeness{
		Id:          e.Id,
		PersonId:    e.PersonId,
		TenantId:    e.TenantId,
		DataId:      e.DataId,
		Type:        e.Type,
		Validations: validations,
	}

}

func (v *Validation) toModel() model.Validation {

	return model.Validation{
		Level:            v.Level,
		Mode:             v.Mode,
		Source:           v.Source,
		ComplementText:   v.ComplementText,
		UpdateDate:       v.UpdateDate,
		CreateDate:       v.CreateDate,
		SourceCreateDate: v.SourceCreateDate,
	}

}
