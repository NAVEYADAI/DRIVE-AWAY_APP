# Sorting — automatic drive → driver assignment (שיבוץ)

## 1. Overview

`POST /Sort` takes every *available* drive and every *available* driver of the caller's company,
assigns drives to drivers with a greedy + backtracking search scored by Mapbox travel time (migrated from Google Distance Matrix, 2026-10), and
stores the result as a new `sort` with `driveofdriver` rows. Drivers see the latest sort's result via
`GetMyDrives` (`02-management.md` §3.3).

Sources: `back/internal/sorting/{Sort.go,mainSort.go,Driver.go,Transportation.go,ChangeToDriver.go,
MinHeap.go,TravelListForDrivers.go,MapboxApi.go}`, `font/src/component/manager/Sort.tsx`.

Last checked against the code: 2026-09-29.

## 2. Data Model

Persisted: `sort`, `driveofdriver` (see `02` §2). In memory (`internal/sorting`):
- `Driver{Id (= driver.id), TimeOfDriver, Gps, WageEarner}`. Every driver starts at **06:00 today**
  (`timeUntilSixAM`) at `driver.address`. `WageEarner` = `driver.isemployee`.
- `Transportation{Id (= drive.id), TimeStart (today at htostart:mtostart), GpsStart, GpsEnd,
  TimeForAllTransportation = FindWay(GpsStart, GpsEnd)}`.
- `ChangeToDriver`: the undo record for one assignment. The driver's new time is
  `TimeStart + duration + 5 min`, and their new position is `GpsEnd`.
- `TravelListForDrivers`: `map[driverId][]driveId`.
- Package-level `var driverMap map[int]Driver`, shared global state (§9.2).

## 3. API Contract — `POST /Sort` (`sorting.CreateSort`)

Body: `{ "userId", "level", "nameOfSort" }`.

| # | condition | status | body |
|---|---|---|---|
| 1 | bind error | 400 | `{"message":"error in bind json"}` |
| 2 | gate result -1/0 (`00` §3.1) | 401 | `{"message":"User has no permissions"}` |
| 3 | `Sort` returns `"we didnt can sort"` | 400 | `{"message":"error in sort "}` (trailing space) |
| 4 | `LoadingDriver` panics (§4, backtracking) | 500 | Gin recovery (no JSON body) |
| 5 | ok | 200 | `{"message":"Sort is created"}`. Inserts `sort(companyid, name)`, then one `driveofdriver(sortid, driveid, driverid)` per assignment. |

With **zero available drives**, row 5 applies: an empty sort is created. `nameOfSort` isn't
validated (an empty name is fine).

## 4. Business Logic — `Sort(drivers, transportations)` (`mainSort.go`)

1. Sort transportations by `TimeStart` (random-pivot quicksort `SortFromTime`, in place).
2. For drive `i` (first visit): build a min-heap over all drivers with
   `Grade(driver, drive) != -1`, where:
   - `t = FindWay(driver.Gps, drive.GpsStart)` (Mapbox Directions API `driving-traffic`, `routes[0].duration` seconds; the Matrix API is not usable for a single pair, it requires >= 2 matrix elements). Places are `"lat,lng"` (used as-is) or free-text addresses (Mapbox Geocoding v6, `country=il`, `language=he`). Geocodes and pair durations are cached in memory for the process lifetime; failures are not cached
   - eligible iff `driver.TimeOfDriver` is before `drive.TimeStart - t`
   - grade = `t` in minutes, **× 0.75 if `WageEarner`**. Lower is better, so salaried drivers are
     preferred.
3. Heap non-empty → pop the best, record the assignment and undo record, move the driver
   (`Gps = GpsEnd`, time = end + 5 min), `i++`.
4. Heap empty → backtrack: `i--`. If `i < 0` → return `"we didnt can sort"`.
5. Revisiting drive `i` (heap exists) → undo its assignment and pop the next candidate. **This branch
   never advances `i`**, so it keeps undoing and popping at the same index until the heap is empty,
   and then `LoadingDriver` **panics** ("No available drivers for transportation"). In practice:
   failure on the first drive → 400, failure on any later drive → 500.

`FindWay` returns `0` on any failure (missing `MAPBOX_ACCESS_TOKEN`, HTTP/parse error, ungeocodable address, unroutable pair), so a
failed Maps call reads as "0 minutes away" and isn't surfaced as an error.

## 5. Client-Side Validation

`Sort.tsx`: no validation. It posts `nameOfSort` (may be empty) with `userId`/`level`.

## 6. UI States & Components

`/manager/Sort` → `Sort.tsx`: field `שם השיבוץ`, button `שליחה`. It shows
`alert("השיבוץ נוצר בהצלחה")` on 2xx and `alert("שגיאה ביצירת השיבוץ")` otherwise. The same page
renders `EditDrives` and `EditDrivers` below, so the manager can toggle availability before sorting.

## 7. Cross-Cutting Flows

Manager toggles availability → creates sort → driver opens `/driver` and sees their drives.
E2E: `e2e/tests/sort-creation.spec.ts` is **`test.fixme`**. Per its header comment, the E2E backend
has no Maps key or route to Google, and the run fails in that environment. Enabling it needs
pointing `mapboxBaseURL` at a stub server for the E2E backend, or a test token.

## 8. Edge Cases / Test Scenarios

1. No available drives → 200, empty sort row, no `driveofdriver`.
2. Drives but no available drivers → 400 `"error in sort "`.
3. Drive starting before 06:00 → no driver is ever eligible, so it fails like #2 (if it's first by
   time) or panics (if it's later).
4. Two drives where the second has no eligible driver → backtrack → **500** (§4 step 5).
5. Salaried vs. non-salaried driver at equal distance → salaried wins (0.75 multiplier).
6. Maps unreachable → all travel times 0 → every driver who is free before `TimeStart` is eligible.
   The grade ties at 0 are broken by heap order.
7. Two concurrent `/Sort` calls → both mutate `driverMap` (run with `-race` to demonstrate).

## 9. Known Discrepancies & Risks

| # | Item | STATUS |
|---|---|---|
| 9.1 | Backtracking branch never advances, so any non-first failure panics → 500. Derived from reading the code. **No test covers backtracking yet** (`Sort_test.go` only has the empty and happy paths). | Bug candidate — high |
| 9.2 | Global `driverMap` → data race across concurrent requests | Bug candidate |
| 9.3 | Maps errors silently become 0-minute travel times | Bug candidate |
| 9.4 | One Maps call per drive plus one per (driver × drive) heap build → cost and latency grow as O(drivers × drives) | Risk |
| 9.5 | `SortId` is re-read via `GetNumOfLastSortByCompanyId` after insert and isn't taken from the insert → race between two sorts of the same company | Bug candidate (minor) |
| 9.6 | Day is always "today" in server local time, and drivers always start at 06:00 | Intentional (current product scope). Confirm with the user. |
