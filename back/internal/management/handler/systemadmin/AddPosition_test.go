package systemadmin_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	authmodel "back/internal/auth/model"
	"back/internal/management/model"
	"back/internal/management/handler/systemadmin"
	"back/internal/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newAddPositionRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/AddPosition", systemadmin.Add_Position)
	return r
}

func validAddPositionBody(companyName string, adminId int, overrides map[string]any) []byte {
	body := map[string]any{
		"userId":      adminId,
		"level":       1,
		"companyName": companyName,
		"fName":       "Jane",
		"lName":       "Roe",
		"id":          "987654321",
		"userName":    "JaneRoe1",
		"password":    "Secret123!",
		"email":       "jane@example.com",
		"phoneNumber": "0500000001",
	}
	for k, v := range overrides {
		body[k] = v
	}
	b, _ := json.Marshal(body)
	return b
}

func TestAddPosition_CreatesUserWithManagerPermissionOnExistingCompany(t *testing.T) {
	testutil.SetupTestDB(t)
	adminId := authmodel.AddUser("AdminUser1")
	model.CreatePermissions(model.NewPermissions(adminId, 0, 1))
	company := model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")

	r := newAddPositionRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddPosition", bytes.NewReader(validAddPositionBody("Acme", adminId, nil)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code, w.Body.String())

	newUserId := authmodel.GetUserByUserName("JaneRoe1")
	require.Greater(t, newUserId, 0)

	perm := model.GetPermissionsByUserId(newUserId)
	require.Equal(t, company.ID, perm.ConpanyId)
	require.Equal(t, 4, perm.Level, "AddPosition should grant level 4 (manager) permission")
}

func TestAddPosition_RejectsRequesterWithoutAdminPermission(t *testing.T) {
	testutil.SetupTestDB(t)
	managerId := authmodel.AddUser("ManagerUser1")
	model.CreatePermissions(model.NewPermissions(managerId, 1, 3))
	model.CreateCompanyWithReturnId("Acme", "AC-1", "1 Main St")

	r := newAddPositionRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddPosition", bytes.NewReader(validAddPositionBody("Acme", managerId, map[string]any{"level": 1})))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.Equal(t, -1, authmodel.GetUserByUserName("JaneRoe1"))
}
