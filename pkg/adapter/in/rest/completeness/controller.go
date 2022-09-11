package completeness

import (
	"net/http"

	"example.com/completude/pkg/application/completeness"
	"github.com/gin-gonic/gin"
)

type completenessController struct {
	service completeness.Service
}

func NewController(service completeness.Service) *completenessController {

	return &completenessController{
		service: service,
	}
}

func (c *completenessController) FindByPerson(ctx *gin.Context) {

	personId := ctx.Query("person-id")
	tenantId := ctx.Query("tenant-id")

	// call to the service
	if completeness, err := c.service.FindByPerson(personId, tenantId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, completeness)
	}

}
