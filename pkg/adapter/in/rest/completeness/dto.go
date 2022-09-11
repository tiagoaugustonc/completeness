package completeness

import (
	"time"
)

type PersonCompleteness struct {
	PersonId string `json:"personId"`
	TenantId string `json:"tenantId"`
	FullName Field  `json:"fullname"`
	Sex      Field  `json:"sex"`
	Type     Field  `json:"type"`
	Document Field  `json:"document"`
}

type Field struct {
	Id          string       `json:"id"`
	Validations []Validation `json:"validations"`
}

type Validation struct {
	Level            int       `json:"level"`
	Mode             int       `json:"mode"`
	Source           int       `json:"source"`
	Text             string    `json:"text"`
	SourceCreateDate time.Time `json:"sourceCreateDate"`
}
