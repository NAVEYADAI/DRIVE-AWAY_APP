package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	authmodel "back/internal/auth/model"
	"back/internal/management/handler"
	"back/internal/management/model"
	"back/internal/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newManagementLogInRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/LogIn", handler.LogIn)
	return r
}

func TestManagementLogIn_ValidCredentialsReturnLevel(t *testing.T) {
	testutil.SetupTestDB(t)
	userId := authmodel.AddUser("SamDriver1")
	authmodel.AddPassword("Secret123!", userId)
	model.CreatePermissions(model.NewPermissions(userId, 1, 3))

	body, _ := json.Marshal(map[string]any{
		"UserName":   "SamDriver1",
		"Password":   "Secret123!",
		"IP":         "127.0.0.1",
		"MacAddress": "AA:BB:CC",
	})

	r := newManagementLogInRouter()
	req := httptest.NewRequest(http.MethodPost, "/LogIn", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var resp struct {
		UserId int `json:"UserId"`
		Level  int `json:"level"`
	}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	require.Equal(t, userId, resp.UserId)
	require.Equal(t, 3, resp.Level)
}

func TestManagementLogIn_WrongPasswordRejected(t *testing.T) {
	testutil.SetupTestDB(t)
	userId := authmodel.AddUser("SamDriver1")
	authmodel.AddPassword("Secret123!", userId)

	body, _ := json.Marshal(map[string]any{
		"UserName":   "SamDriver1",
		"Password":   "WrongPassword1!",
		"IP":         "127.0.0.1",
		"MacAddress": "AA:BB:CC",
	})

	r := newManagementLogInRouter()
	req := httptest.NewRequest(http.MethodPost, "/LogIn", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestManagementLogIn_UnknownUserRejected(t *testing.T) {
	testutil.SetupTestDB(t)

	body, _ := json.Marshal(map[string]any{
		"UserName":   "NoSuchUser1",
		"Password":   "Secret123!",
		"IP":         "127.0.0.1",
		"MacAddress": "AA:BB:CC",
	})

	r := newManagementLogInRouter()
	req := httptest.NewRequest(http.MethodPost, "/LogIn", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

// More than 5 failed attempts from the same IP+MAC should be throttled (see
// authservice.TryLogInIpAndMAcAddress: `DidNotSucceed > 5`) even with the
// right password. A single failed attempt must NOT trigger throttling —
// that was this test's original (wrong) premise; see
// TestManagementLogIn_SingleFailedAttemptDoesNotThrottle below for that case.
func TestManagementLogIn_ThrottlesRepeatedFailedAttempts(t *testing.T) {
	testutil.SetupTestDB(t)
	userId := authmodel.AddUser("SamDriver1")
	authmodel.AddPassword("Secret123!", userId)
	for i := 0; i < 6; i++ {
		authmodel.AddTryLogIn("SamDriver1", "wrong", "10.0.0.1", "AA:BB:CC", false, time.Now(), 4)
	}

	body, _ := json.Marshal(map[string]any{
		"UserName":   "SamDriver1",
		"Password":   "Secret123!",
		"IP":         "10.0.0.1",
		"MacAddress": "AA:BB:CC",
	})

	r := newManagementLogInRouter()
	req := httptest.NewRequest(http.MethodPost, "/LogIn", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusConflict, w.Code, w.Body.String())
}

// A single prior failed attempt (below the >5 threshold) must not throttle
// a subsequent correct login.
func TestManagementLogIn_SingleFailedAttemptDoesNotThrottle(t *testing.T) {
	testutil.SetupTestDB(t)
	userId := authmodel.AddUser("SamDriver1")
	authmodel.AddPassword("Secret123!", userId)
	authmodel.AddTryLogIn("SamDriver1", "wrong", "10.0.0.1", "AA:BB:CC", false, time.Now(), 4)

	body, _ := json.Marshal(map[string]any{
		"UserName":   "SamDriver1",
		"Password":   "Secret123!",
		"IP":         "10.0.0.1",
		"MacAddress": "AA:BB:CC",
	})

	r := newManagementLogInRouter()
	req := httptest.NewRequest(http.MethodPost, "/LogIn", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())
}
