package <%= serviceName %>

import (
	"github.com/labstack/echo/v4"
	"<%= projectName %>/<%= serviceName %>/middlewares"
	"<%= projectName %>/<%= serviceName %>/routes"
)

func Set(e *echo.Echo) {
    middlewares.Set(e)

	routes.Web(e)
	routes.Api(e)
}
