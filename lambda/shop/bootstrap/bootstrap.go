package bootstrap

import (
    "lambda/lambda/models/form/caller"
    "lambda/lambda/models/form/validationCaller"
    gridCaller "lambda/lambda/models/grid/caller"
    "github.com/labstack/echo/v4"
    //"github.com/lambda-platform/arcGIS"
    "github.com/lambda-platform/chart"
    /*
    	|----------------------------------------------
    	| Generated Models
    	|----------------------------------------------
    */
    "github.com/lambda-platform/lambda"
    "github.com/lambda-platform/agent"
    "github.com/lambda-platform/krud"
    "github.com/lambda-platform/adminmodule"
    /*
    	|----------------------------------------------
    	| Graphql
    	|----------------------------------------------
    */
    "lambda/lambda/graph"
    /*
    	|----------------------------------------------
    	| PRO MODULES
    	|----------------------------------------------
    */

    "github.com/lambda-platform/moqup"

    /*
    	|----------------------------------------------
    	| App
    	|----------------------------------------------
    */
    "lambda/app/middlewares"
    "lambda/routes"
    lambdaUtils "github.com/lambda-platform/lambda/utils"
    /*
    	|----------------------------------------------
    	| Template Utils
    	|----------------------------------------------
    */
    templateUtils "github.com/lambda-platform/template/utils"

    "html/template"
)

func Set() *lambda.Lambda {

    /*
    	|----------------------------------------------
    	| Lambda Platform
    	|----------------------------------------------
    */
    lambda := lambda.New(&lambda.Settings{
        ModuleName:      "lambda",
    })
    KrudMiddleWares := []echo.MiddlewareFunc{
    // arcGIS.MW(caller.GetMODEL, gridCaller.GetMODEL),
    }
    agent.Set(lambda.Echo)
    krud.Set(lambda.Echo,  gridCaller.GetMODEL, caller.GetMODEL, validationCaller.GetMessages, validationCaller.GetRules, KrudMiddleWares, true)

    /*
    	|----------------------------------------------
    	| MODULES
    	|----------------------------------------------
    */
    graph.Set(lambda.Echo)
    adminmodule.Set(lambda.Echo, lambda.ModuleName, gridCaller.GetMODEL)
    chart.Set(lambda.Echo)
    moqup.Set(lambda.Echo)
    middlewares.Set(lambda.Echo)

    /*
    	|----------------------------------------------
    	| Template Register
    	|----------------------------------------------
    */

    templates := lambdaUtils.GetTemplates(lambda.Echo)
    TemplatePath := templateUtils.AbsolutePath()
    //* REGISTER VIEWS */
    templates["admin.html"] = template.Must(template.ParseFiles(
        TemplatePath + "views/paper.html",
    ))
    template.Must(templates["admin.html"].ParseFiles(
        "views/admin.html",
    ))

    /*
    	|----------------------------------------------
    	| ROUTES
    	|----------------------------------------------
    */
    routes.Api(lambda.Echo)
    routes.Web(lambda.Echo)

    return lambda
}
