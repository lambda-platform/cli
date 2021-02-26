package validationCaller

import (
	"<%= projectName %>/app/models/form/validations"
	"github.com/thedevsaddam/govalidator"
)

func GetMessages(schema_id string) map[string][]string {

	switch schema_id {

	case "crud_form":
		return validations.GetCrudFromMessages()

	case "analytic_form":
		return validations.GetAnalyticFormMessages()

	case "notification_target_form":
		return validations.GetNotificationTargetMessages()

	case "menu_form":
		return validations.GetMenuFormMessages()

	case "user_form":
		return validations.GetUserFormMessages()

	case "user_profile":
		return validations.GetUserProfileMessages()

	case "user_password":
		return validations.GetUserPasswordMessages()

	}
	return govalidator.MapData{}

}
