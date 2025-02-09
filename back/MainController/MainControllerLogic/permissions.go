package MainControllerLogic

import "back/MainController/MainControllerGormModels"

// פונקציה לבדיקה אם יש למשתמש הרשאה להוסיף סוגי משתמשים
func ReturnPermissionForSingUp(UserId, level int) int {
	id := MainControllerGormModels.GetPermissionByUserIdAndLevel(UserId, level)
	if id.Level != 0 {
		if level == 1 || level == 2 {
			return 1
		}
		if level == 3 || level == 4 {
			return 2
		}
		return 0
	}
	return -1
}
