package main

import (
	"fmt"
	"os"

	outCompleteness "example.com/completude/pkg/adapter/out/completeness"
	"example.com/completude/pkg/infra/di"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

}

func run() error {

	container := di.BuildContainer()
	engine := gin.Default()

	//teste
	if err := container.Invoke(func(db *gorm.DB) {
		db.AutoMigrate(&outCompleteness.Completeness{}, &outCompleteness.Validation{})
	}); err != nil {
		panic(err)
	}

	//

	server := NewServer(container, engine)
	server.SetupRouter()

	return server.Start()
}
