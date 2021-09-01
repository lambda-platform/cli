package controllers

import (
    "encoding/json"
    "github.com/labstack/echo/v4"
    agentModels "github.com/lambda-platform/agent/models"
    agentUtils "github.com/lambda-platform/agent/utils"
    krudModels "github.com/lambda-platform/krud/models"
    "github.com/lambda-platform/lambda/DB"
    puzzleModels "github.com/lambda-platform/lambda/models"
    "github.com/lambda-platform/lambda/config"
    "github.com/lambda-platform/lambda/utils"
    "net/http"
)

type Permissions struct {
    DefaultMenu string      `json:"default_menu"`
    Extra       interface{} `json:"extra"`
    MenuID      int         `json:"menu_id"`
    Permissions interface{} `json:"permissions"`
}

func AdminIndex(UseNotify bool) echo.HandlerFunc {
    return func(c echo.Context) error {
        User := agentUtils.AuthUserObject(c)
        Role := agentModels.Role{}
        DB.DB.Where("id = ?", User["role"]).Find(&Role)
        Permissions_ := Permissions{}
        json.Unmarshal([]byte(Role.Permissions), &Permissions_)
        Menu := puzzleModels.VBSchema{}
        DB.DB.Where("id = ?", Permissions_.MenuID).Find(&Menu)
        MenuSchema := new(interface{})
        json.Unmarshal([]byte(Menu.Schema), &MenuSchema)
        Kruds := []krudModels.Krud{}
        DB.DB.Where("deleted_at IS NULL").Find(&Kruds)

        FirebaseConfig := config.LambdaConfig.Notify.FirebaseConfig
        //csrfToken := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
        csrfToken := ""
        return c.Render(http.StatusOK, "admin.html", map[string]interface{}{
            "UseNotify":                 UseNotify,
            "title":                     config.LambdaConfig.Title,
            "extraStyles":               config.LambdaConfig.ControlPanel.ExtraStyles,
            "extraScripts":              config.LambdaConfig.ControlPanel.ExtraScripts,
            "primaryColor":              config.LambdaConfig.ControlPanel.PrimaryColor,
            "themeColors":               config.LambdaConfig.ControlPanel.ThemeColors,
            "themeMode":                 config.LambdaConfig.ControlPanel.ThemeMode,
            "favicon":                   config.LambdaConfig.Favicon,
            "logo":                      config.LambdaConfig.Logo,
            "logo_light":                config.LambdaConfig.ControlPanel.LogoLight,
            "logo_dark":                 config.LambdaConfig.ControlPanel.LogoDark,
            "brandBtnUrl":               config.LambdaConfig.ControlPanel.BrandBtnURL,
            "permissions":               Permissions_,
            "menu":                      MenuSchema,
            "cruds":                     Kruds,
            "withCrudLog":               config.LambdaConfig.WithCrudLog,
            "User":                      User,
            "data_form_custom_elements": config.LambdaConfig.DataFormCustomElements,
            "firebase_config":           FirebaseConfig,
            "mix":                       utils.Mix,
            "csrfToken":                 csrfToken,
        })
    }
}
