package bootstrap

import (
    "<%= projectName %>/app/models/form/caller"
    "<%= projectName %>/app/models/form/validationCaller"
    gridCaller "<%= projectName %>/app/models/grid/caller"
    /*
    	|----------------------------------------------
    	| Generated Models
    	|----------------------------------------------
    */
    "github.com/lambda-platform/lambda"
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
    })

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
