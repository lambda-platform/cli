package <%= serviceName %>

import (
	"github.com/labstack/echo/v4"
	"<%= projectName %>/<%= serviceName %>/routes"
)

func Set(e *echo.Echo) {
	routes.Web(e)
	routes.Api(e)
}
