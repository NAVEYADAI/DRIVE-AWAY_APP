package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"back/internal/auth/handler"
	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func doLogInRequest(t *testing.T, body handler.LogIn) *httptest.ResponseRecorder {
	t.Helper()
	router := gin.New()
	router.POST("/LogIn", handler.LogInUser)

	payload, err := json.Marshal(body)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/LogIn", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestLogInUser_Success(t *testing.T) {
	testutil.SetupTestDB(t)
	userId := model.AddUser("alice")
	model.AddPassword("correct-pass", userId)

	w := doLogInRequest(t, handler.LogIn{UserName: "alice", Password: "correct-pass", IP: "1.2.3.4", MacAddress: "AA:BB"})

	assert.Equal(t, http.StatusOK, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	assert.Equal(t, "LogIn Success", body["message"])
	assert.Equal(t, float64(userId), body["UserId"])
}

func TestLogInUser_WrongPassword(t *testing.T) {
	testutil.SetupTestDB(t)
	userId := model.AddUser("alice")
	model.AddPassword("correct-pass", userId)

	w := doLogInRequest(t, handler.LogIn{UserName: "alice", Password: "wrong-pass", IP: "1.2.3.4", MacAddress: "AA:BB"})

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestLogInUser_UnknownUsername is a regression test that only became
// meaningful after fixing model.GetUserByUserName to return -1 (rather than
// 0) for a nonexistent user — see model/users_test.go. Before that fix this
// request would have fallen through to comparing against user id 0's
// (nonexistent) password and returned 401 "Password incorrect" instead of
// the correct 404 "User not found".
func TestLogInUser_UnknownUsername(t *testing.T) {
	testutil.SetupTestDB(t)

	w := doLogInRequest(t, handler.LogIn{UserName: "ghost", Password: "anything", IP: "1.2.3.4", MacAddress: "AA:BB"})

	assert.Equal(t, http.StatusNotFound, w.Code)
	var body map[string]interface{}
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	assert.Equal(t, "User not found", body["message"])
}

// TestLogInUser_RateLimitedAfterSixFailures is a regression test that only
// became meaningful after fixing the single-quoted 'macAddress' column bug
// in model.GetTryLogInByIpAndMacAddress — see model/trylogin_test.go.
func TestLogInUser_RateLimitedAfterSixFailures(t *testing.T) {
	testutil.SetupTestDB(t)
	userId := model.AddUser("alice")
	model.AddPassword("correct-pass", userId)

	for i := 0; i < 6; i++ {
		doLogInRequest(t, handler.LogIn{UserName: "alice", Password: "wrong-pass", IP: "5.5.5.5", MacAddress: "FF:FF"})
	}

	w := doLogInRequest(t, handler.LogIn{UserName: "alice", Password: "correct-pass", IP: "5.5.5.5", MacAddress: "FF:FF"})

	assert.Equal(t, http.StatusConflict, w.Code, "even the correct password should be rejected once the IP/MAC is rate-limited")
}

func TestLogInUser_MalformedJSON(t *testing.T) {
	testutil.SetupTestDB(t)

	router := gin.New()
	router.POST("/LogIn", handler.LogInUser)
	req := httptest.NewRequest(http.MethodPost, "/LogIn", bytes.NewReader([]byte("not json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
}
