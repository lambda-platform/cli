package validationCaller

import (
	"github.com/thedevsaddam/govalidator"
	"lambda/lambda/models/form/validations"
)

func GetRules(schema_id string) map[string][]string {

	switch schema_id {

	case "crud_form":
		return validations.GetCrudFromRules()

	case "analytic_form":
		return validations.GetAnalyticFormRules()

	case "notification_target_form":
		return validations.GetNotificationTargetRules()

	case "menu_form":
		return validations.GetMenuFormRules()

	case "user_form":
		return validations.GetUserFormRules()

	case "user_profile":
		return validations.GetUserProfileRules()

	case "user_password":
		return validations.GetUserPasswordRules()

	}
	return govalidator.MapData{}

}
