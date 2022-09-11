package completeness

import (
	"time"
)

type Completeness struct {
	Id          string
	PersonId    string
	TenantId    string
	DataId      string
	Type        int
	Validations []Validation
}

type Validation struct {
	Completeness     Completeness
	Level            int
	Mode             int
	Source           int
	UpdateDate       time.Time
	CreateDate       time.Time
	SourceCreateDate time.Time
	ComplementText   string
}
