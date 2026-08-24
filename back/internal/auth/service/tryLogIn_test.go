package service_test

import (
	"testing"
	"time"

	"back/internal/auth/model"
	"back/internal/auth/service"
	"back/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestTryLogInIpAndMAcAddress_AllowsUnderThreshold(t *testing.T) {
	testutil.SetupTestDB(t)

	for i := 0; i < 5; i++ {
		model.AddTryLogIn("alice", "wrong", "1.2.3.4", "AA:BB", false, time.Now(), 4)
	}

	good, _ := service.TryLogInIpAndMAcAddress("1.2.3.4", "AA:BB", time.Now())
	assert.True(t, good, "5 failed attempts should still be allowed (threshold is more than 5)")
}

// TestTryLogInIpAndMAcAddress_BlocksOverThreshold exercises the brute-force
// protection this service exists for. It only became meaningful once the
// model.GetTryLogInByIpAndMacAddress single-quoted-column bug (see
// model/trylogin_test.go) was fixed — before that fix this test would have
// failed, since the service always saw zero attempts no matter how many
// failed logins actually happened.
func TestTryLogInIpAndMAcAddress_BlocksOverThreshold(t *testing.T) {
	testutil.SetupTestDB(t)

	for i := 0; i < 6; i++ {
		model.AddTryLogIn("alice", "wrong", "1.2.3.4", "AA:BB", false, time.Now(), 4)
	}

	good, message := service.TryLogInIpAndMAcAddress("1.2.3.4", "AA:BB", time.Now())
	assert.False(t, good)
	assert.Equal(t, "Too many attempts", message)
}

func TestTryLogInIpAndMAcAddress_SuccessfulAttemptsDontCount(t *testing.T) {
	testutil.SetupTestDB(t)

	for i := 0; i < 10; i++ {
		model.AddTryLogIn("alice", "correct", "1.2.3.4", "AA:BB", true, time.Now(), -1)
	}

	good, _ := service.TryLogInIpAndMAcAddress("1.2.3.4", "AA:BB", time.Now())
	assert.True(t, good, "only failed (Connect=false) attempts should count toward the threshold")
}

func TestTryLogInIpAndMAcAddress_DifferentIpNotAffected(t *testing.T) {
	testutil.SetupTestDB(t)

	for i := 0; i < 10; i++ {
		model.AddTryLogIn("alice", "wrong", "1.2.3.4", "AA:BB", false, time.Now(), 4)
	}

	good, _ := service.TryLogInIpAndMAcAddress("9.9.9.9", "ZZ:ZZ", time.Now())
	assert.True(t, good)
}
