package handlers

import (
	"encoding/json"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/config"
	agentModels "github.com/lambda-platform/agent/models"
	agentUtils "github.com/lambda-platform/agent/utils"
	krudModels "github.com/lambda-platform/krud/models"

	puzzleModels "github.com/lambda-platform/lambda/DB/DBSchema/models"
	"github.com/lambda-platform/lambda/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)


type Permissions struct {
	DefaultMenu string `json:"default_menu"`
	Extra       interface{} `json:"extra"`
	MenuID      int `json:"menu_id"`
	Permissions interface{} `json:"permissions"`
}



func Index(UseNotify bool) echo.HandlerFunc {
	return func(c echo.Context) error {

		User := agentUtils.AuthUserObject(c)


		Role := agentModels.Role{}

		DB.DB.Where("id = ?", User["role"]).Find(&Role)

		Permissions_ := Permissions{}

		json.Unmarshal([]byte(Role.Permissions), &Permissions_)


		Menu := puzzleModels.VBSchema{}
		DB.DB.Where("id = ?",Permissions_.MenuID).Find(&Menu)
		MenuSchema := new(interface{})
		json.Unmarshal([]byte(Menu.Schema), &MenuSchema)
		Kruds := []krudModels.Krud{}
		DB.DB.Where("deleted_at IS NULL").Find(&Kruds)


		FirebaseConfig := config.LambdaConfig.Notify.FirebaseConfig
		//csrfToken := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
		csrfToken := ""
		return c.Render(http.StatusOK, "control.html", map[string]interface{}{
			"UseNotify":       UseNotify,
			"title":       config.LambdaConfig.Title,
			"extraStyles":       config.LambdaConfig.ControlPanel.ExtraStyles,
			"extraScripts":       config.LambdaConfig.ControlPanel.ExtraScripts,
			"primaryColor":       config.LambdaConfig.ControlPanel.PrimaryColor,
			"themeColors":       config.LambdaConfig.ControlPanel.ThemeColors,
			"themeMode":       config.LambdaConfig.ControlPanel.ThemeMode,
			"favicon":     config.LambdaConfig.Favicon,
			"logo":     config.LambdaConfig.Logo,
			"logo_light":     config.LambdaConfig.ControlPanel.LogoLight,
			"logo_dark":     config.LambdaConfig.ControlPanel.LogoDark,
			"brandBtnUrl":     config.LambdaConfig.ControlPanel.BrandBtnURL,
			"permissions": Permissions_,
			"menu":        MenuSchema,
			"cruds":       Kruds,
			"withCrudLog":        config.LambdaConfig.WithCrudLog,
			"User":        User,
			"data_form_custom_elements": config.LambdaConfig.DataFormCustomElements,
			"firebase_config":           FirebaseConfig,
			"mix":                       utils.Mix,
			"csrfToken":                       csrfToken,

		})
	}
}


func Form(c echo.Context) error {

	//csrfToken := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	csrfToken := ""
	schema_id := c.Param("schema_id")
	id := c.Param("id")
	hidemap := c.QueryParam("hidemap")

	return c.Render(http.StatusOK, "form.html", map[string]interface{}{
		"title":       config.LambdaConfig.Title,
		"favicon":     config.LambdaConfig.Favicon,
		"mix":                       utils.Mix,
		"schema_id":                       schema_id,
		"data_form_custom_elements": config.LambdaConfig.DataFormCustomElements,
		"id":                       id,
		"csrfToken":                       csrfToken,
		"hidemap":                       hidemap,
	})

}
