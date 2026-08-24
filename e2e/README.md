# E2E tests

Playwright specs driving the real app (frontend + backend) end to end, wired to the isolated `db_test` Postgres instance — **never the dev `db` or production RDS**.

## Run order

From the repo root:

```bash
# 1. Bring up the isolated test DB
docker compose up -d db_test

# 2. Create the dedicated E2E database once (separate from the driveaway_test
#    DB that Go unit tests use, so the two suites never race each other's
#    table truncation on the same server)
docker compose exec -T db_test psql -U postgres -c "CREATE DATABASE driveaway_e2e;"   # ok if it already exists

# 3. Point the backend at db_test/driveaway_e2e and build the frontend image
docker compose -f docker-compose.yml -f docker-compose.e2e.yml up -d --build backend
docker compose build frontend

# 4. Start the frontend on the configured port (3000 by default — if that
#    port is already taken by something else on your machine, run a one-off
#    container on another port and pass BASE_URL, see below)
docker compose -f docker-compose.yml -f docker-compose.e2e.yml up -d frontend

# 5. Run the suite (seeds fixed users via global-setup.ts automatically)
cd e2e
npm test
```

If port 3000 is unavailable locally, run the frontend on an alternate port instead of step 4:

```bash
docker run -d --name e2e-frontend -p 3005:3000 \
  -v "$(pwd)/../font:/app" -v /app/node_modules drive-away_app-frontend
BASE_URL=http://localhost:3005 npm test
```

**When you're done**, point the backend back at the dev DB so local dev isn't left wired to the test database:

```bash
docker compose up -d backend
```

## What's seeded

`global-setup.ts` truncates and reseeds a fixed set of rows in `driveaway_e2e` before the suite runs:

| user | password | level |
|---|---|---|
| `e2e-admin` | `e2e-admin-pass123` | 1 (system administrator) |
| `e2e-manager` | `e2e-manager-pass123` | 3 (manager) |
| `e2e-driver` | `e2e-driver-pass123` | 5 (driver) |

plus one company (`E2E Test Co`). Specs run serially (`workers: 1` in `playwright.config.ts`) since they share this one fixture set — a spec that mutates data (adds a driver/drive) could otherwise race a spec reading the same tables.

## Coverage

- `login.spec.ts` — golden path login → manager home.
- `login-failures.spec.ts` — wrong password / unknown user.
- `permission-gating.spec.ts` — each role can/can't reach the routes it should/shouldn't.
- `manager-crud.spec.ts` — add driver / add drive, confirm they show up in the edit lists.
- `sort-creation.spec.ts` — create a sort/assignment.
- `logout.spec.ts` — session clears and protected routes lock back down.

Per `CLAUDE.md`, this layer covers golden paths and the highest-value edge cases only — exhaustive permutation coverage belongs in the backend/frontend unit tests, not here.
