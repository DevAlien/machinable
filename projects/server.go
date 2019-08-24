package projects

import (
	"github.com/anothrnick/machinable/dsi/interfaces"

	"github.com/anothrnick/machinable/middleware"
	"github.com/anothrnick/machinable/projects/apikeys"
	"github.com/anothrnick/machinable/projects/documents"
	"github.com/anothrnick/machinable/projects/logs"
	"github.com/anothrnick/machinable/projects/resources"
	"github.com/anothrnick/machinable/projects/sessions"
	"github.com/anothrnick/machinable/projects/users"
	"github.com/gin-gonic/gin"
)

// CreateRoutes creates a gin.Engine for the project routes
func CreateRoutes(datastore interfaces.Datastore) *gin.Engine {

	router := gin.Default()
	router.Use(middleware.OpenCORSMiddleware())
	router.Use(middleware.SubDomainMiddleware())

	// set routes -> handlers for each package
	resources.SetRoutes(router, datastore)
	documents.SetRoutes(router, datastore)
	logs.SetRoutes(router, datastore)
	users.SetRoutes(router, datastore)
	sessions.SetRoutes(router, datastore)
	apikeys.SetRoutes(router, datastore)

	return router
}
