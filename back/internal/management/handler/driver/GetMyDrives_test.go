package driver_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authmodel "back/internal/auth/model"
	"back/internal/management/handler/driver"
	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newGetMyDrivesRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/GetMyDrives", driver.GetMyDrives)
	return r
}

func TestGetMyDrives_ReturnsAssignedDrivesForLatestSort(t *testing.T) {
	testutil.SetupTestDB(t)

	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	driverUserId := authmodel.AddUser("SamDriver1")
	model.CreatePermissions(model.NewPermissions(driverUserId, company.ID, 5))
	model.CreateDriver(model.NewDriver(driverUserId, company.ID, "5 Driver Ln", true))

	model.CreateSort(company.ID, "Sort A")
	sort := model.GetLastSortByCompanyId(company.ID)

	driveId := model.CreateDriveReturnId(model.NewDrive("Morning Run", 8, 0, company.ID, "10,10", "20,20"))
	driverRowId := model.GetDriverIdByUserIdAndCompanyId(driverUserId, company.ID)
	model.CreateDriveOfDriver(sort.ID, driveId, driverRowId)

	body, _ := json.Marshal(map[string]any{"userId": driverUserId, "level": 5})

	r := newGetMyDrivesRouter()
	req := httptest.NewRequest(http.MethodPost, "/GetMyDrives", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var resp struct {
		List []driver.DriveToReturn `json:"list of drives"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	require.Len(t, resp.List, 1)
	require.Equal(t, "Morning Run", resp.List[0].Name)
}

func TestGetMyDrives_RejectsRequesterWithoutPermission(t *testing.T) {
	testutil.SetupTestDB(t)

	body, _ := json.Marshal(map[string]any{"userId": 999, "level": 5})

	r := newGetMyDrivesRouter()
	req := httptest.NewRequest(http.MethodPost, "/GetMyDrives", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}
