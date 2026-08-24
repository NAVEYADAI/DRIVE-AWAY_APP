package sorting

import (
	"back/internal/management/model"
	"back/internal/management/service"
	"back/internal/management/handler"

	"fmt"

	"github.com/gin-gonic/gin"
)


type Sortstruct struct {
	handler.UserIdAndLevel

	NameOfSort string `json:"nameOfSort"`
}

func CreateSort(c *gin.Context) {
	var tmp Sortstruct
	err := c.ShouldBindJSON(&tmp)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "error in bind json"})
		return
	}
	// בדיקת הרשאות
	level := service.ReturnPermissionForSingUp(tmp.UserId, tmp.Level)
	if level == -1 || level == 0 {
		c.IndentedJSON(401, gin.H{"message": "User has no permissions"}) // 401
		return
	}
	idCompany := model.GetCompanyIdByLevelAndUserId(tmp.UserId, tmp.Level)

	listOfDrives := model.GetDriveAviablesByCompanyId(idCompany)

	var finalyListOfDrives []Transportation
	var drive model.Drive
	for i := 0; i < len(listOfDrives); i++ {
		if listOfDrives[i].Available == true {
			drive = model.GetDriveById(listOfDrives[i].IdDrive)
			finalyListOfDrives = append(finalyListOfDrives, NewTransportation(drive.ID, drive.MToStart,
				drive.HToStart, drive.GpsStart, drive.GpsEnd))
		}
	}

	listOfDrivers := model.GetAllDriversAviableInCompany(idCompany)

	var finalyListOfDrivers []Driver
	var driver model.Driver
	for i := 0; i < len(listOfDrivers); i++ {
		if listOfDrivers[i].Available == true {
			driver = model.GetDriverById(listOfDrivers[i].IdDriver)
			finalyListOfDrivers = append(finalyListOfDrivers, NewDriver(driver.ID, driver.Address, driver.IsEmployee))
		}

	}
	mapOfSort, message := Sort(finalyListOfDrivers, finalyListOfDrives)
	if message == "we didnt can sort" {
		c.IndentedJSON(400, gin.H{"message": "error in sort "})
		return
	}
	model.CreateSort(idCompany, tmp.NameOfSort)
	SortId := model.GetNumOfLastSortByCompanyId(idCompany)
	for driver, destinations := range mapOfSort {
		fmt.Printf("Driver %d:\n", driver)
		for _, destination := range destinations {
			fmt.Printf("  Destination: %d\n", destination)
			model.CreateDriveOfDriver(SortId, destination, driver)
		}
	}
	c.IndentedJSON(200, gin.H{"message": "Sort is created"})
}
