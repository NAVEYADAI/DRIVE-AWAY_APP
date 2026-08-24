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

func newDeleteDriverRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/DeleteDriver", manager.Delete_Driver)
	return r
}

// Delete_Driver resolves the requester's company with
// model.GetPermissionLevelByUserId(tmpDriver.UserId) — but that function
// returns the requester's permission LEVEL, not their company id. It's then
// passed straight into GetAllDriversInCompany as if it were a company id.
// With a manager at level 3 whose real company id isn't 3, the driver lookup
// silently searches the wrong "company" and never finds anyone to delete.
func TestDeleteDriver_ManagerRemovesDriverAvailability(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	// Pick a company id that does NOT equal the manager's permission level
	// (3), so the bug (using level as if it were company id) is caught
	// rather than accidentally passing by coincidence.
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

	r := newDeleteDriverRouter()
	req := httptest.NewRequest(http.MethodPost, "/DeleteDriver", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())
	require.Empty(t, model.GetAllDriverAvailables(), "the driver's availability row should be removed")
}
