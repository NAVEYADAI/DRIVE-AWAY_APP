package Manager

import "back/MainController/controller"

// DriverData Struct
type DriverData struct {
	controller.UserIdAndLevel
	DriverName string `json:"driverName"`
}
