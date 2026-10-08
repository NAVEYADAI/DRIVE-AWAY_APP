# E2E — AGENTS.md

> Playwright, driving the real stack (frontend container + backend container + `db_test`).
> **`README.md` in this directory is the authoritative run order**: follow it step by step,
> including the last step that points the backend back at the dev DB.

## Rules

- Target DB is `db_test` / database `driveaway_e2e` (via `../docker-compose.e2e.yml`), separate from
  `driveaway_test`, which the Go unit tests use. Never the dev `db`, never the production RDS.
- `global-setup.ts` TRUNCATEs every table and seeds fixed users by `psql` inside the `db_test`
  container. The seeded credentials and levels are listed in `README.md`. A new spec reuses these
  fixtures, and if it needs more rows it adds them to `SEED_SQL` in `global-setup.ts`, not ad hoc in
  the spec.
- The backend must have started against `driveaway_e2e` at least once before seeding, because
  `AutoMigrate` creates the schema.
- `workers: 1`, `fullyParallel: false`. Specs share one fixture set, so keep it that way.
- One project only: Desktop Chrome. `BASE_URL` overrides `http://localhost:3000`.
- The app uses `window.alert` for feedback. Specs accept it with `page.on('dialog', d => d.accept())`.
- Locators: Hebrew labels and roles (`getByLabel('שם משתמש')`, `getByRole('button', { name: 'כניסה' })`).
- A spec that documents a known limitation and not a desired behavior says so in a header
  comment (see `signup-not-reachable.spec.ts`, and `sort-creation.spec.ts`, which is `test.fixme`
  because the sorting path calls the real Google Maps API).

## Scope

The top of the test pyramid: golden paths plus the highest-value edge cases (login, permission
gating, core CRUD, sort creation). Exhaustive permutations belong in `back/` and `font/` unit tests.

Run: `npm --prefix e2e test` (after the README setup steps).
