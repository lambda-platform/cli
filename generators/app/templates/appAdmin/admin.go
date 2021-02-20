package appAdmin
import (
	"html/template"
	"<%= projectName %>/appAdmin/handlers"
	"github.com/labstack/echo/v4"
	"github.com/lambda-platform/agent/agentMW"
	"github.com/lambda-platform/lambda/config"
	"github.com/lambda-platform/lambda/utils"

)

func Set(e *echo.Echo, UseNotify bool) {

	if config.Config.App.Migrate == "true"{
		//utils.AutoMigrateSeed()
	}

	/* REGISTER VIEWS */
	templates := utils.GetTemplates(e)
	templates["admin.html"] = template.Must(template.ParseFiles("appAdmin/templates/admin.html"))
	templates["form.html"] = template.Must(template.ParseFiles("appAdmin/templates/form.html"))

	/* ROUTES */
	e.GET("/admin", handlers.Index(UseNotify), agentMW.IsLoggedInCookie)
	e.GET("/form/:schema_id/:id", handlers.Form)




}
