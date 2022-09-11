package completeness

import "github.com/gin-gonic/gin"

func (c *completenessController) SetupRouter(group *gin.RouterGroup) {

	router := group.Group("/completeness")

	router.GET("/", c.FindByPerson)
}
