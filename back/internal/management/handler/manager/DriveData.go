package manager

import "back/internal/management/handler"

// DriveData Struct
type DriveData struct {
	handler.UserIdAndLevel
	DriveName string `json:"driveName"`
}
