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

func newAddCompanyRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/AddCompany", systemadmin.Add_Company)
	return r
}

func validAddCompanyBody(overrides map[string]any) []byte {
	body := map[string]any{
		"userId":      1,
		"level":       1,
		"companyName": "Acme",
		"companyID":   "AC-1",
		"address":     "1 Main St",
		"fName":       "John",
		"lName":       "Doe",
		"id":          "123456789",
		"userName":    "JohnDoe1",
		"password":    "Secret123!",
		"email":       "john@example.com",
		"phoneNumber": "0500000000",
	}
	for k, v := range overrides {
		body[k] = v
	}
	b, _ := json.Marshal(body)
	return b
}

func TestAddCompany_AdminCreatesCompanyAndUser(t *testing.T) {
	testutil.SetupTestDB(t)
	// Give the requester a real Users row (not just a Permissions row) so its
	// id can't collide with the id the handler assigns the new company-owner
	// user it creates (both come from the same "users" SERIAL sequence).
	adminId := authmodel.AddUser("AdminUser1")
	model.CreatePermissions(model.NewPermissions(adminId, 0, 1)) // requester is a system admin (level 1)

	r := newAddCompanyRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddCompany", bytes.NewReader(validAddCompanyBody(map[string]any{"userId": adminId})))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code, w.Body.String())

	companies := model.GetAllCompanies()
	require.Len(t, companies, 1)
	require.Equal(t, "Acme", companies[0].Name)

	newUserId := authmodel.GetUserByUserName("JohnDoe1")
	require.NotZero(t, newUserId)
	require.NotEqual(t, adminId, newUserId)

	perm := model.GetPermissionsByUserId(newUserId)
	require.Equal(t, companies[0].ID, perm.ConpanyId)
	require.Equal(t, 3, perm.Level, "the newly created company owner should get level 3")
}

func TestAddCompany_RejectsRequesterWithoutAdminPermission(t *testing.T) {
	testutil.SetupTestDB(t)
	model.CreatePermissions(model.NewPermissions(1, 1, 3)) // only a manager, not level 1/2

	r := newAddCompanyRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddCompany", bytes.NewReader(validAddCompanyBody(nil)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.Empty(t, model.GetAllCompanies(), "no company should be created when the requester lacks permission")
}

func TestAddCompany_RejectsWeakPassword(t *testing.T) {
	testutil.SetupTestDB(t)
	model.CreatePermissions(model.NewPermissions(1, 0, 1))

	r := newAddCompanyRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddCompany", bytes.NewReader(validAddCompanyBody(map[string]any{"password": "weak"})))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	require.Empty(t, model.GetAllCompanies())
	require.Equal(t, -1, authmodel.GetUserByUserName("JohnDoe1"), "no user should be created with a weak password")
}

func TestAddCompany_MissingRequiredFieldRejected(t *testing.T) {
	testutil.SetupTestDB(t)
	model.CreatePermissions(model.NewPermissions(1, 0, 1))

	r := newAddCompanyRouter()
	req := httptest.NewRequest(http.MethodPost, "/AddCompany", bytes.NewReader(validAddCompanyBody(map[string]any{"email": ""})))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}
