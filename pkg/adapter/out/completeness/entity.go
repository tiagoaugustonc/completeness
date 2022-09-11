package completeness

import "time"

type Completeness struct {
	Id          string `gorm:"primaryKey"`
	PersonId    string
	TenantId    string
	DataId      string
	Type        int
	Validations []Validation
}

type Validation struct {
	CompletenessId   int
	Level            int
	Mode             int
	Source           int
	UpdateDate       time.Time
	CreateDate       time.Time
	SourceCreateDate time.Time
	ComplementText   string
}
