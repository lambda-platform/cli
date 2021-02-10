package main

import (

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	/*
		|----------------------------------------------
		| Generated Models
		|----------------------------------------------
	*/
	"github.com/lambda-platform/lambda"
	"<%= projectName %>/models/form/caller"
	"<%= projectName %>/models/form/validationCaller"
	gridCaller "<%= projectName %>/models/grid/caller"

	/*
    	|----------------------------------------------
    	| APP, APP ADMIN
    	|----------------------------------------------
    */
    "<%= projectName %>/<%= serviceName %>"
    "<%= projectName %>/appAdmin"
)

func main() {

	lambda := lambda.New(&lambda.Settings{
		ModuleName:"<%= projectName %>",
		GetGridMODEL:gridCaller.GetMODEL,
		GetFormMODEL:caller.GetMODEL,
		GetFormMessages:validationCaller.GetMessages,
		GetFormRules:validationCaller.GetRules,
		/*
			|----------------------------------------------
			| PRO MODULES middlewares based on KRUD
			|----------------------------------------------
		*/
		KrudMiddleWares: []echo.MiddlewareFunc{

		},
		//Back-End permission validation
		KrudWithPermission:true,
	})
	/*
		|----------------------------------------------
		| Echo useful MIDDLEWARES
		|----------------------------------------------
	*/
	lambda.Echo.Use(middleware.Secure())
	lambda.Echo.Use(middleware.Recover())
	//App.Echo.Use(middleware.Logger())
	lambda.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*", "http://localhost:*","http://127.0.0.1:*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Requested-With", "x-requested-with"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	/*
		|----------------------------------------------
		| APP, APP ADMIN
		|----------------------------------------------
	*/
	<%= serviceName %>.Set(lambda.Echo)
    appAdmin.Set(lambda.Echo, false)

	lambda.Start()


}


