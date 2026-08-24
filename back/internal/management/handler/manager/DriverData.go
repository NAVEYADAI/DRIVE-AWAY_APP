package manager

import "back/internal/management/handler"

// DriverData Struct
type DriverData struct {
	handler.UserIdAndLevel
	DriverName string `json:"driverName"`
}
