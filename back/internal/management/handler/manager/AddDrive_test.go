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

func newAddDriveRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/AddDrive", manager.Add_Drive)
	return r
}

func TestAddDrive_ManagerCreatesDriveAndAvailability(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	model.CreatePermissions(model.NewPermissions(managerId, company.ID, 3))

	body, _ := json.Marshal(map[string]any{
		"userId":       managerId,
		"level":        3,
		"nameOfDriver": "Morning Run",
		"timeToStartH": 8,
		"timeToStartM": 30,
		"gpsStart":     "10,10",
		"gpsEnd":       "20,20",
	})

	r := newAddDriveRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddDrive", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	drives := model.GetAllDrives()
	require.Len(t, drives, 1)
	require.Equal(t, "Morning Run", drives[0].Name)
	require.Equal(t, company.ID, drives[0].CompanyId)

	avail := model.GetAllDriveAvailables()
	require.Len(t, avail, 1)
	require.Equal(t, drives[0].ID, avail[0].IdDrive)
	require.True(t, avail[0].Available)
}

func TestAddDrive_RejectsRequesterWithoutPermission(t *testing.T) {
	testutil.SetupTestDB(t)
	body, _ := json.Marshal(map[string]any{
		"userId":       999,
		"level":        3,
		"nameOfDriver": "Morning Run",
	})

	r := newAddDriveRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddDrive", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.Empty(t, model.GetAllDrives())
}
