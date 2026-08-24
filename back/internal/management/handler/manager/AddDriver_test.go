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

func newAddDriverRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/AddDriver", manager.Add_Driver)
	return r
}

func validAddDriverBody(managerId int, overrides map[string]any) []byte {
	body := map[string]any{
		"userId":      managerId,
		"level":       3,
		"address":     "5 Driver Ln",
		"isEmployee":  true,
		"fName":       "Sam",
		"lName":       "Driver",
		"id":          "111222333",
		"userName":    "SamDriver1",
		"password":    "Secret123!",
		"email":       "sam@example.com",
		"phoneNumber": "0500000002",
	}
	for k, v := range overrides {
		body[k] = v
	}
	b, _ := json.Marshal(body)
	return b
}

func TestAddDriver_ManagerCreatesDriverForOwnCompany(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	model.CreatePermissions(model.NewPermissions(managerId, company.ID, 3))

	r := newAddDriverRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddDriver", bytes.NewReader(validAddDriverBody(managerId, nil)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code, w.Body.String())

	newUserId := authmodel.GetUserByUserName("SamDriver1")
	require.Greater(t, newUserId, 0)

	driver := model.GetDriverById(newUserId)
	require.Equal(t, company.ID, driver.CompanyId)
	require.Equal(t, "5 Driver Ln", driver.Address)
	require.True(t, driver.IsEmployee)

	perm := model.GetPermissionsByUserId(newUserId)
	require.Equal(t, 5, perm.Level, "a newly added driver should get level 5")

	avail := model.GetAllDriverAvailables()
	require.Len(t, avail, 1)
	require.True(t, avail[0].Available)
}

func TestAddDriver_RejectsRequesterWithoutPermission(t *testing.T) {
	testutil.SetupTestDB(t)
	// No permissions row at all for this requester.
	r := newAddDriverRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddDriver", bytes.NewReader(validAddDriverBody(999, nil)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.Equal(t, -1, authmodel.GetUserByUserName("SamDriver1"))
}

func TestAddDriver_MalformedJSONRejected(t *testing.T) {
	testutil.SetupTestDB(t)
	r := newAddDriverRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddDriver", bytes.NewReader([]byte("{not json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}
