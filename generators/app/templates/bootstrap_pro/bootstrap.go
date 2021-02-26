package bootstrap

import (
    "<%= projectName %>/app/models/form/caller"
    "<%= projectName %>/app/models/form/validationCaller"
    gridCaller "<%= projectName %>/app/models/grid/caller"
    "github.com/labstack/echo/v4"
    "github.com/lambda-platform/arcGIS"
    "github.com/lambda-platform/chart"
    /*
    	|----------------------------------------------
    	| Generated Models
    	|----------------------------------------------
    */
    "github.com/lambda-platform/lambda"
    /*
    	|----------------------------------------------
    	| Graphql
    	|----------------------------------------------
    */
    //"<%= projectName %>/app/graph"
    /*
    	|----------------------------------------------
    	| PRO MODULES
    	|----------------------------------------------
    */
    "github.com/lambda-platform/crudlogger"
    "github.com/lambda-platform/dataanalytic"
    "github.com/lambda-platform/moqup"
    "github.com/lambda-platform/notify"
    /*
    	|----------------------------------------------
    	| App
    	|----------------------------------------------
    */
    "<%= projectName %>/app/middlewares"
    "<%= projectName %>/routes"
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
        ModuleName:      "<%= projectName %>",
        GetGridMODEL:    gridCaller.GetMODEL,
        GetFormMODEL:    caller.GetMODEL,
        GetFormMessages: validationCaller.GetMessages,
        GetFormRules:    validationCaller.GetRules,
        /*
        	|----------------------------------------------
        	| PRO MODULES middlewares based on KRUD
        	|----------------------------------------------
        */
        KrudMiddleWares: []echo.MiddlewareFunc{
            crudlogger.MW(),
            notify.MW(),
            arcGIS.MW(caller.GetMODEL, gridCaller.GetMODEL),
        },
        //Back-End permission validation
        KrudWithPermission: true,
    })

    /*
    	|----------------------------------------------
    	| Graphql
    	|----------------------------------------------
    */
    //graph.Set(lambda.Echo)

    /*
    	|----------------------------------------------
    	| PRO MODULES
    	|----------------------------------------------
    */
    crudlogger.Set(lambda.Echo)
    notify.Set(lambda.Echo)
    arcGIS.Set(lambda.Echo)
    dataanalytic.Set(lambda.Echo)
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
