package logs

import (
	"bitbucket.org/nsjostrom/machinable/dsi/interfaces"
	"bitbucket.org/nsjostrom/machinable/middleware"
	"github.com/gin-gonic/gin"
)

// SetRoutes sets all of the appropriate routes to handlers for project collections
func SetRoutes(engine *gin.Engine, datastore interfaces.Datastore) error {
	// create new Logs handler with datastore
	handler := New(datastore)

	logs := engine.Group("/logs")
	logs.Use(middleware.AppUserJwtAuthzMiddleware())
	logs.Use(middleware.AppUserProjectAuthzMiddleware(datastore))
	logs.GET("/", handler.GetProjectLogs)

	return nil
}
