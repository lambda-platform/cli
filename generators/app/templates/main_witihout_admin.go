package main

import (
	"github.com/labstack/echo/v4"
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
    	| APP
    	|----------------------------------------------
    */
    "<%= projectName %>/<%= serviceName %>"

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
		| APP, APP ADMIN
		|----------------------------------------------
	*/
	<%= serviceName %>.Set(lambda.Echo)

	lambda.Start()


}


