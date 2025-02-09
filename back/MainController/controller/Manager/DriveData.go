package Manager

import "back/MainController/controller"

// DriveData Struct
type DriveData struct {
	controller.UserIdAndLevel
	DriveName string `json:"driveName"`
}
