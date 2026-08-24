// Package testutil connects backend tests to the isolated db_test Postgres
// instance (see docker-compose.yml) instead of the dev/production database
// that config.Connect() would otherwise pick up from the environment.
package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"

	authmodel "back/internal/auth/model"
	"back/internal/config"
	managementmodel "back/internal/management/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// lockKey is an arbitrary shared id for a Postgres session advisory lock.
// Go runs different packages' tests as separate OS processes, potentially in
// parallel (`go test ./...` without `-p 1`) — since every package's tests
// share the one db_test database and reset it via TRUNCATE, two packages'
// tests running at the same moment would stomp on each other's fixtures.
// Holding this lock for the lifetime of each test serializes that access
// across every process/package without requiring `-p 1` to be remembered.
const lockKey = 918273645

// allTables lists every table in the schema, used to reset state between
// test runs. Keep in sync with the TableName() of each model below.
var allTables = []string{
	"trylogin",
	"doubleregistration",
	"personalinformation",
	"phone",
	"mail",
	"name",
	"password",
	"users",
	"sort",
	"driveofdriver",
	"driveavailable",
	"driveravailable",
	"drive",
	"driver",
	"permissions",
	"company",
}

func envOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// SetupTestDB connects to the isolated test database (db_test in
// docker-compose.yml, published on localhost:5434 by default), migrates the
// full schema (every model across every domain, so any test can rely on any
// table existing), truncates all tables so the test starts from a clean
// slate, and points config.DB at it. It never touches back/.env or the
// dev/production database.
//
// Call this at the start of any test that needs the database, e.g.:
//
//	func TestSomething(t *testing.T) {
//	    testutil.SetupTestDB(t)
//	    // ... use config.DB / model.* functions as normal
//	}
func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	host := envOrDefault("TEST_DB_HOST", "localhost")
	port := envOrDefault("TEST_DB_PORT", "5434")
	user := envOrDefault("TEST_DB_USER", "postgres")
	password := envOrDefault("TEST_DB_PASSWORD", "postgres")
	dbname := envOrDefault("TEST_DB_DBNAME", "driveaway_test")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("testutil: failed to connect to test DB at %s:%s (is `docker compose up -d db_test` running?): %v", host, port, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("testutil: failed to get underlying *sql.DB: %v", err)
	}
	// Every test opens its own pool via gorm.Open, and none of it was ever
	// being closed — hundreds of tests across parallel packages exhausted
	// Postgres's max_connections. Keep each test's pool tiny and always
	// close it when the test ends.
	sqlDB.SetMaxOpenConns(3)
	sqlDB.SetMaxIdleConns(1)
	t.Cleanup(func() {
		_ = sqlDB.Close()
	})

	ctx := context.Background()
	lockConn, err := sqlDB.Conn(ctx)
	if err != nil {
		t.Fatalf("testutil: failed to acquire a connection for the test lock: %v", err)
	}
	if _, err := lockConn.ExecContext(ctx, "SELECT pg_advisory_lock($1)", lockKey); err != nil {
		t.Fatalf("testutil: failed to acquire cross-process test lock: %v", err)
	}
	t.Cleanup(func() {
		_, _ = lockConn.ExecContext(context.Background(), "SELECT pg_advisory_unlock($1)", lockKey)
		_ = lockConn.Close()
	})

	if err := db.AutoMigrate(
		&authmodel.DoubleRegistration{},
		&authmodel.Mail{},
		&authmodel.Name{},
		&authmodel.Password{},
		&authmodel.Phone{},
		&authmodel.TryLogIn{},
		&authmodel.Users{},
		&authmodel.PersonalInformation{},
		&managementmodel.Company{},
		&managementmodel.Drive{},
		&managementmodel.DriveAvailable{},
		&managementmodel.DriveOfDriver{},
		&managementmodel.Driver{},
		&managementmodel.DriverAvailable{},
		&managementmodel.Permissions{},
		&managementmodel.Sort{},
	); err != nil {
		t.Fatalf("testutil: failed to migrate test DB schema: %v", err)
	}

	for _, table := range allTables {
		if err := db.Exec(fmt.Sprintf(`TRUNCATE TABLE "%s" RESTART IDENTITY CASCADE`, table)).Error; err != nil {
			t.Fatalf("testutil: failed to truncate table %s: %v", table, err)
		}
	}

	config.DB = db
	return db
}
