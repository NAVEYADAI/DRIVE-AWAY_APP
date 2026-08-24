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

func newViewOrNotDriveRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/ViewOrNotDrive", manager.ViewOrNotDrive)
	return r
}

func TestViewOrNotDrive_TogglesAvailability(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	model.CreatePermissions(model.NewPermissions(managerId, company.ID, 3))
	driveId := model.CreateDriveReturnId(model.NewDrive("Morning Run", 8, 0, company.ID, "10,10", "20,20"))
	model.CreateDriveAvailable(model.NewDriveAvailable(company.ID, driveId, true))

	body, _ := json.Marshal(map[string]any{"userId": managerId, "level": 3, "driveName": "Morning Run"})

	r := newViewOrNotDriveRouter()
	req := httptest.NewRequest(http.MethodPost, "/ViewOrNotDrive", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())
	require.False(t, model.GetAvailableDriveByDriveId(driveId).Available)
}

// Regression test for the same cross-company authorization bug fixed in
// Delete_Drive: GetDriveByName isn't scoped to a company, so without checking
// drive.CompanyId == idCompany, any manager could flip another company's
// drive availability just by knowing its name.
func TestViewOrNotDrive_RejectsRequesterFromDifferentCompany(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	otherManagerId := authmodel.AddUser("OtherManager1")
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	otherCompany := model.CreateCompanyWithReturnId("Globex", "GX-1", "2 Side St")
	model.CreatePermissions(model.NewPermissions(managerId, company.ID, 3))
	model.CreatePermissions(model.NewPermissions(otherManagerId, otherCompany.ID, 3))
	driveId := model.CreateDriveReturnId(model.NewDrive("Morning Run", 8, 0, company.ID, "10,10", "20,20"))
	model.CreateDriveAvailable(model.NewDriveAvailable(company.ID, driveId, true))

	body, _ := json.Marshal(map[string]any{"userId": otherManagerId, "level": 3, "driveName": "Morning Run"})

	r := newViewOrNotDriveRouter()
	req := httptest.NewRequest(http.MethodPost, "/ViewOrNotDrive", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.True(t, model.GetAvailableDriveByDriveId(driveId).Available, "a manager from a different company must not toggle another company's drive")
}
