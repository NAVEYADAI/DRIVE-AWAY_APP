package model_test

import (
	"testing"
	"time"

	"back/internal/auth/model"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddTryLogIn_AndGetById(t *testing.T) {
	testutil.SetupTestDB(t)

	now := time.Now()
	model.AddTryLogIn("alice", "pw", "1.2.3.4", "AA:BB", true, now, -1)

	all := model.GetAllTryLogIn()
	require.Len(t, all, 1)
	assert.Equal(t, "alice", all[0].UserName)
	assert.True(t, all[0].Connect)
	assert.Equal(t, -1, all[0].Cause)

	got := model.GetTryLogInById(all[0].ID)
	assert.Equal(t, "alice", got.UserName)
}

func TestUpdateAndDeleteTryLogIn(t *testing.T) {
	testutil.SetupTestDB(t)

	model.AddTryLogIn("alice", "pw", "1.2.3.4", "AA:BB", false, time.Now(), 4)
	all := model.GetAllTryLogIn()
	require.Len(t, all, 1)

	entry := all[0]
	entry.Cause = -1
	model.UpdateTryLogIn(entry)
	updated := model.GetTryLogInById(entry.ID)
	assert.Equal(t, -1, updated.Cause)

	model.DeleteTryLogIn(entry)
	assert.Empty(t, model.GetAllTryLogIn())
}

// TestGetTryLogInByIpAndMacAddress_MatchesRealAttempts is a regression test.
// The query built its WHERE clause as `ip = ? AND 'macAddress' = ?` — the
// column name was wrapped in single quotes, which SQL treats as a *string
// literal* ("macAddress") rather than an identifier. The macAddress column
// was therefore never actually compared; the clause only matched if the
// literal string "macAddress" happened to equal the parameter, which never
// happens in practice. GetTryLogInByIpAndMacAddress always returned an empty
// slice, which meant service.TryLogInIpAndMAcAddress's "more than 5 failed
// attempts" brute-force protection could never trigger, no matter how many
// failed logins occurred for a given IP/MAC.
func TestGetTryLogInByIpAndMacAddress_MatchesRealAttempts(t *testing.T) {
	testutil.SetupTestDB(t)

	model.AddTryLogIn("alice", "pw", "1.2.3.4", "AA:BB", false, time.Now(), 4)
	model.AddTryLogIn("alice", "pw", "1.2.3.4", "AA:BB", false, time.Now(), 4)
	model.AddTryLogIn("bob", "pw", "9.9.9.9", "CC:DD", false, time.Now(), 4)

	results := model.GetTryLogInByIpAndMacAddress("1.2.3.4", "AA:BB", time.Now())

	assert.Len(t, results, 2, "must return the attempts that actually match this ip+mac pair")
}

func TestGetTryLogInByIpAndMacAddress_NoMatch(t *testing.T) {
	testutil.SetupTestDB(t)

	model.AddTryLogIn("alice", "pw", "1.2.3.4", "AA:BB", false, time.Now(), 4)

	results := model.GetTryLogInByIpAndMacAddress("9.9.9.9", "ZZ:ZZ", time.Now())
	assert.Empty(t, results)
}
