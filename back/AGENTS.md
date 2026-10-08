# Backend — AGENTS.md

> DRIVE-AWAY — Go backend (Gin + GORM + PostgreSQL). Transportation company management:
> companies, drivers, drives, and an automatic drive→driver assignment ("sort" / שיבוץ).
> Read `../specs/00-shared-conventions.md` before touching auth, permissions, or anything
> cross-cutting. It is the exact reference, taken from reading the code. This file is only
> an orientation.

## Tech Stack

- **Language:** Go 1.21, module name `back` (imports look like `back/internal/...`)
- **HTTP:** Gin (`github.com/gin-gonic/gin`), CORS via `gin-contrib/cors`
- **ORM:** GORM + `gorm.io/driver/postgres`. The schema is created by `AutoMigrate` at startup
  (`internal/auth/auth.go::LogIn`, `internal/management/management.go::MainController`). There are
  no migration files.
- **External:** Mapbox Geocoding v6 + Directions (`internal/sorting/MapboxApi.go`, token from env
  `MAPBOX_ACCESS_TOKEN`); Gmail SMTP (`internal/auth/handler/SendMail.go`)
- **Testing:** `testing` + `testify` (`require`/`assert`); DB tests go through `internal/testutil`
- **Port:** `:2909` (hardcoded in `cmd/server/main.go`)

## Layout

```
cmd/server/main.go          entrypoint: config.Connect(), CORS, route groups
internal/config/db.go       global config.DB (*gorm.DB), env-driven DSN
internal/auth/              users, passwords, mail, phone, personal info, login attempts
  handler/                  SingUp, LogIn, ChangeUsername, ChangePassword, SendMail, ForgetMyUserName
  model/                    GORM structs + package-level CRUD funcs on config.DB
  routes/                   generic REST CRUD per table (/loginApi/<Table>/...)
  service/                  validation (username rules, password strength, login throttling)
internal/management/        companies, drivers, drives, permissions, availability, sorts
  handler/{systemadmin,manager,driver}/   role-scoped business endpoints (/MainController/...)
  model/, routes/, service/ same pattern as auth
internal/sorting/           the assignment algorithm + POST /Sort handler
internal/testutil/          SetupTestDB(t) — the ONLY way tests touch a database
```

## Route groups (from `cmd/server/main.go`)

| Prefix | Wired in | What |
|---|---|---|
| `POST /login` | `auth.Login` | The login the frontend actually uses (returns `UserId` + `level`) |
| `/loginApi/...` | `auth.LogIn` | Sign-up, the other login, change username/password, and raw CRUD for every auth table |
| `/MainController/...` | `management.MainController` | Role-scoped business endpoints + raw CRUD for every management table |
| `POST /Sort` | `sorting.CreateSort` | Run the assignment algorithm for the caller's company |

## Architecture patterns (follow these, and know where they fall short)

### Models are package-level functions on the global `config.DB`

Every model file has a struct with explicit `gorm:"column:..."` tags, a `TableName()` returning a
lowercase table name, and free functions (`AddX`, `GetXById`, `UpdateX`, ...) that call
`config.DB` directly. No repository interfaces, no context passing. New models follow the same shape.
**If you add a model, also add it to the `AutoMigrate` list in its domain's wiring file AND to
`testutil.SetupTestDB` (both the `AutoMigrate` call and the `allTables` truncate list)**, or tests
will run against a missing or dirty table.

Lookups mostly signal "not found" with zero values (`ID == 0`, or a `-1` sentinel in some auth
models like `GetUserByUserName`) and not with errors. GORM errors are almost never checked. Read the
specific function before you assume which sentinel it uses.

### Auth = "trust the client" (no tokens)

There is **no session, JWT, or cookie**. Login returns `UserId` and `level`, and the frontend stores
them in `localStorage`. Every `/MainController/*` business endpoint takes `{userId, level}` in the
JSON body (`handler.UserIdAndLevel`) and checks that a matching `permissions` row exists. Anyone who
knows or guesses a userId+level pair can act as that user. Don't present a new endpoint as
"secured" just because it runs this check. See `specs/00-shared-conventions.md` §3.

### Permission levels

`permissions.level`: `1`/`2` = system administrator, `3` = company manager (created by AddCompany),
`4` = position in an existing company (AddPosition), `5` = driver. The shared gate is
`management/service.ReturnPermissionForSingUp(userId, level)` → `1` (admin), `2` (manager), `0`
(other existing level, e.g. drivers), `-1` (no row). Manager endpoints reject `-1` and `0`. Admin
endpoints require `== 1`. **Don't reuse this helper as a generic "is valid user" check.** It maps
drivers to `0`, which is the same "rejected" value. `GetMyDrives` had exactly this bug (see the
NOTE comment in `handler/driver/GetMyDrives.go`).

### Response shape

`c.IndentedJSON(status, gin.H{"message": ...})`. Watch for these inconsistencies, which tests must
match exactly:
- `POST /login` (`auth.Login`) uses the misspelled key **`"messege"`**, and the frontend
  `LoginResponse` type relies on it. `/loginApi/LogIn` and `/MainController/LogIn` use `"message"`.
- Bind errors return 400 in management and **415** in auth handlers.
- Several handlers write **two responses** (e.g. `201 User Created` followed by
  `201 Company Created`, or a 409 for the username followed by a 409 for the password). Gin keeps
  the first status. Assert on status + first body, not on the full concatenated body.

## Testing (binding: see root `CLAUDE.md` testing policy)

- Run: `cd back && go test ./...`. DB tests need `docker compose up -d db_test` (localhost:5434,
  DB `driveaway_test`).
- **Test package convention:** almost every test file is an **external test package** (`model_test`,
  `manager_test`, `service_test`, `sorting_test`, ...). `sorting` also has one in-package file
  (`MapboxApi_test.go`) for unexported helpers. New tests: use `<pkg>_test` unless you need
  unexported access.
- DB tests start with `testutil.SetupTestDB(t)`. It migrates everything, truncates every table, takes
  a cross-process advisory lock (so `go test ./...` in parallel is safe), and points `config.DB` at
  the test DB. Never call `config.Connect()` in a test, never read `back/.env`.
- Handler tests build a bare Gin router (`gin.SetMode(gin.TestMode); r := gin.New(); r.POST(...)`)
  and drive it with `httptest`. See `internal/management/handler/manager/DeleteDriver_test.go` for
  the reference style, including the "what broke and why" comment on regression tests.
- Anything calling `FindWay` hits the real Mapbox API unless it's stubbed. `mapboxBaseURL` and `mapsHTTPClient` in
  `MapboxApi.go` are package vars so tests can point them at an `httptest` server. `FindWay` caches
  geocodes and durations per process (`cacheMu`), so reset those maps in tests.

## Database & environment rules

- `config.Connect()` loads `back/.env` (via godotenv) if present. That file points at the
  **production RDS**. Never read it, print it, or run anything against it. Local dev uses the `db`
  compose service, tests use `db_test`, E2E uses `db_test`/`driveaway_e2e`.
- Schema changes happen via `AutoMigrate` on server start, i.e. **on whatever DB the server
  connects to**, including prod on the next deploy. A column rename or type change needs the user's
  explicit go-ahead before it merges to `main` (push to `main` deploys, see `.github/workflows/deploy.yml`).
- Formatting: a Claude Code hook runs `gofmt -w` on every `.go` file edited. A few files were
  never gofmt'ed, so the first edit to one of them also shows whitespace changes in the diff.
