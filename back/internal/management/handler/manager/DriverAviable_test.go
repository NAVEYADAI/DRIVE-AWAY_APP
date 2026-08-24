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

func newViewOrNotDriverRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/ViewOrNotDriver", manager.ViewOrNotDriver)
	return r
}

// Regression test for the same "level used as company id" bug fixed in
// Delete_Driver (see DeleteDriver_test.go) — ViewOrNotDriver had the
// identical mistake.
func TestViewOrNotDriver_TogglesAvailability(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	for company.ID == 3 {
		model.DeleteCompany(company)
		company = model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	}
	model.CreatePermissions(model.NewPermissions(managerId, company.ID, 3))

	driverUserId := authmodel.AddUser("DriverUser1")
	authmodel.AddPersonalInformation("Sam", "Driver", "111222333", driverUserId)
	model.CreateDriver(model.NewDriver(driverUserId, company.ID, "5 Driver Ln", true))
	model.CreateDriverAvailable(model.NewDriverAvailable(company.ID, driverUserId, true))

	body, _ := json.Marshal(map[string]any{"userId": managerId, "level": 3, "driverName": "Sam"})

	r := newViewOrNotDriverRouter()
	req := httptest.NewRequest(http.MethodPost, "/ViewOrNotDriver", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())
	all := model.GetAllDriverAvailables()
	require.Len(t, all, 1)
	require.False(t, all[0].Available)
}
