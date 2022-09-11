package di

import (
	"example.com/completude/pkg/adapter/out"
	outCompleteness "example.com/completude/pkg/adapter/out/completeness"
	"example.com/completude/pkg/application/completeness"
	"example.com/completude/pkg/application/person"
	"example.com/completude/pkg/infra/config"
	"example.com/completude/pkg/infra/orm"
	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {

	// config
	container.Provide(config.NewConfig)

	// Database
	container.Provide(orm.NewDB)

	// completeness
	container.Provide(outCompleteness.NewCompletenessRepository)
	container.Provide(completeness.NewService)

	// person
	container.Provide(out.NewPersonRepository)
	container.Provide(person.NewService)

	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
