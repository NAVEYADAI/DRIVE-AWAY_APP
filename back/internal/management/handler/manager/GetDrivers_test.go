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

func newGetDriversRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/GetDrivers", manager.Get_Drivers)
	return r
}

func TestGetDrivers_ReturnsOnlyOwnCompanysDrivers(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")
	otherCompany := model.CreateCompanyWithReturnId("Globex", "GX-1", "2 Side St")
	model.CreatePermissions(model.NewPermissions(managerId, company.ID, 3))

	samId := authmodel.AddUser("SamDriver1")
	authmodel.AddPersonalInformation("Sam", "Driver", "111", samId)
	model.CreateDriverAvailable(model.NewDriverAvailable(company.ID, samId, true))

	otherId := authmodel.AddUser("OtherDriver1")
	authmodel.AddPersonalInformation("Otherguy", "Driver", "222", otherId)
	model.CreateDriverAvailable(model.NewDriverAvailable(otherCompany.ID, otherId, true))

	body, _ := json.Marshal(map[string]any{"userId": managerId, "level": 3})

	r := newGetDriversRouter()
	req := httptest.NewRequest(http.MethodPost, "/GetDrivers", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var resp struct {
		List []manager.DriverToReturn `json:"List"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	require.Len(t, resp.List, 1)
	require.Equal(t, "Sam", resp.List[0].NameOfDriver)
}
