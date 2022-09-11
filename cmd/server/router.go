package main

import (
	"github.com/gin-gonic/gin"

	completenessCtl "example.com/completude/pkg/adapter/in/rest/completeness"
	personCtl "example.com/completude/pkg/adapter/in/rest/person"
	"example.com/completude/pkg/application/completeness"
	"example.com/completude/pkg/application/person"
)

func (s *server) SetupRouter() {

	// Config the server
	apiGroup := s.engine.Group("api/v1")

	// Person route
	s.personRouter(apiGroup)

	// completeness route
	s.completenessRouter(apiGroup)

}

func (s *server) personRouter(group *gin.RouterGroup) {

	if err := s.container.Invoke(func(personSrv person.Service) {

		personCtl := personCtl.NewController(personSrv)
		personCtl.SetupRouter(group)

	}); err != nil {
		panic(err)
	}

}

func (s *server) completenessRouter(group *gin.RouterGroup) {

	if err := s.container.Invoke(func(completenessSrv completeness.Service) {

		completenessCtl := completenessCtl.NewController(completenessSrv)
		completenessCtl.SetupRouter(group)

	}); err != nil {
		panic(err)
	}

}
