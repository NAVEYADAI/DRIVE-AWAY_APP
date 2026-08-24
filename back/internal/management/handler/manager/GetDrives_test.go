package manager_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authmodel "back/internal/auth/model"
	"back/internal/management/handler/manager"
	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newGetDrivesRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/GetDrives", manager.Get_Drives)
	return r
}

func TestGetDrives_ReturnsOnlyOwnCompanysDrives(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	otherCompany := model.CreateCompanyWithReturnId("Globex", "GX-1", "2 Side St")
	model.CreatePermissions(model.NewPermissions(managerId, company.ID, 3))

	driveId := model.CreateDriveReturnId(model.NewDrive("Morning Run", 8, 0, company.ID, "10,10", "20,20"))
	model.CreateDriveAvailable(model.NewDriveAvailable(company.ID, driveId, true))

	otherDriveId := model.CreateDriveReturnId(model.NewDrive("Other Run", 9, 0, otherCompany.ID, "1,1", "2,2"))
	model.CreateDriveAvailable(model.NewDriveAvailable(otherCompany.ID, otherDriveId, true))

	body, _ := json.Marshal(map[string]any{"userId": managerId, "level": 3})

	r := newGetDrivesRouter()
	req := httptest.NewRequest(http.MethodPost, "/GetDrives", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var resp struct {
		List []manager.DriveToReturn `json:"List"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	require.Len(t, resp.List, 1)
	require.Equal(t, "Morning Run", resp.List[0].NameOfDrive)
}
