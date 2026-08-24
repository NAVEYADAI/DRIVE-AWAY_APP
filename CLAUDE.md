# DRIVE-AWAY_APP — Project Conventions

Monorepo: Go backend (`back/`, standard layout under `cmd/` + `internal/`) and React/TypeScript/Vite frontend (`font/`). Local dev runs via `docker-compose.yml` (services: `db`, `backend`, `frontend`). A separate `db_test` service provides an isolated Postgres instance for automated tests — **never point tests at `db` or the production RDS.**

## Testing policy (adopted 2026-08-24)

This project runs a real test suite. The rules below are load-bearing, not aspirational — treat them as required steps of "done," not follow-up work.

1. **Every bug fix ships with a regression test.** When a bug is found (by a human, in review, or by Claude), the fix is not complete until a test exists that reproduces the exact failure path that led to the bug and fails without the fix. Write the failing test first when practical, confirm it fails, then fix and confirm it passes. Name/comment the test so the original symptom is traceable (e.g. `TestLogin_EmptyPasswordDoesNotCrash` with a one-line comment on what broke and why).
2. **Backend (`back/`)**: Go's standard `testing` package + `testify` (`require`/`assert`) for assertions. Tests live alongside the code they test (`*_test.go` in the same package) except where an existing package already uses a colocated `_test` external-test-package convention (`internal/sorting`). Tests that touch the database use the `back/internal/testutil` helper to connect to the isolated test DB — never call `config.Connect()` directly in a test.
3. **Frontend (`font/`)**: Vitest + React Testing Library for unit/component tests (colocated as `*.test.tsx` next to the component). Test user-visible behavior (what renders, what a click does) over implementation details.
4. **E2E (`e2e/`)**: Playwright, driving the real running stack (frontend + backend + `db_test`) through actual browser interactions for the critical end-to-end flows (login, core CRUD, sort creation). E2E tests are the top of the pyramid — cover golden paths and the highest-value edge cases here, not exhaustive permutations (that's what unit tests are for).
5. **Test database isolation**: the `db_test` Postgres service (`docker-compose.yml`) is dedicated to automated tests, on its own port/volume, fully separate from the `db` dev service and from the production RDS instance referenced in `back/.env`. Backend tests default to connecting to `db_test` via `back/internal/testutil`; never wire a test to read `back/.env` or the production DSN.
6. Run before considering backend work done: `cd back && go test ./...`. Run before considering frontend work done: `cd font && npm test`. Run before considering a cross-stack change done: the relevant `e2e/` spec.

## Architecture notes

- Backend follows the standard Go project layout: `back/cmd/server` (entrypoint), `back/internal/{config,auth,management,sorting}` (domain packages, each with `handler/model/service/routes` sub-packages where applicable). See `back/internal/config/db.go` for the DB connection (env-driven: `DB_HOST`/`DB_PORT`/`DB_USER`/`DB_PASSWORD`/`DB_DBNAME`/`DB_SSLMODE`).
- Frontend is Vite + TypeScript (migrated from Create React App). MUI is the component library; `font/src/theme.ts` is the single source of styling tokens (palette, shape, typography) — don't hardcode colors/spacing in components when a theme token exists.
- CORS on the backend allows any `http://localhost:<port>` origin in dev (see `back/cmd/server/main.go`) since the frontend dev-server port can shift.
