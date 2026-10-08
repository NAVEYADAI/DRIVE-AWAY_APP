# Management — companies, positions, drivers, drives

## 1. Overview

Everything under `back/internal/management/`: the system-admin actions (register a company with its
manager, add a position to an existing company), the manager actions (add, list, toggle, and delete
drivers and drives), and the driver's "my drives" view. Plus the matching frontend screens under
`font/src/component/{systemAdministrator,manager,drivers}/`.

Permission gate, company resolution, and response quirks: see `00-shared-conventions.md` §3–§4.
User-creation rules (username/password validation): see `01-auth.md` §4.

Sources: `back/internal/management/{management.go,handler/**,model/*,service/permissions.go}`,
`font/src/component/**`.

Last checked against the code: 2026-09-29.

## 2. Data Model (`back/internal/management/model/`)

| table | struct | columns | notes |
|---|---|---|---|
| `company` | `Company` | `id`, `name`, `idcompany` (string business id), `address` | |
| `permissions` | `Permissions` | `id`, `userid`, `companyid`, `level` | see `00` §2 |
| `driver` | `Driver` | `id`, `userid`, `companyid`, `address`, `isemployee` | `address` is the driver's starting location for sorting |
| `driveravailable` | `DriverAvailable` | `id`, `idcompany`, `iddriver`, `available` | ⚠️ `iddriver` holds the driver's **user id**, not `driver.id` (`AddDriver.go:73`) |
| `drive` | `Drive` | `id`, `name`, `htostart`, `mtostart`, `gpsstart`, `gpsend`, `companyid` | GPS fields are `"lat,lng"` or free-text addresses passed to Mapbox (`sorting.FindWay`). The UI fields (`LocationField`) allow typing or picking a point on a map, which fills `"lat,lng"` |
| `driveavailable` | `DriveAvailable` | `id`, `idcompany`, `iddrive`, `available` | `iddrive` = `drive.id` |
| `sort` | `Sort` | `id`, `companyid`, `name` | one row per sorting run |
| `driveofdriver` | `DriveOfDriver` | `id`, `sortid`, `driveid`, `driverid` | `driverid` = `driver.id` (not user id) |

Naming trap: `model.GetDriverById(id)` queries **`WHERE userid = ?`** (it takes a user id).

Drivers are identified in the manager UI and API **by first name** (`personalinformation.fname`,
via `authmodel.GetFNameByUserID`). Drives are identified **by name**.

## 3. API Contract — all `POST /MainController/<name>`

Common: bind error → `400 {"message":"error in BindJSON"}`; gate failure →
`401 {"message":"User has no permissions"}`.

### 3.1 System admin (gate: `ReturnPermissionForSingUp == 1`)

**`AddCompany`** (`systemadmin/AddCompany.go`) — body: `companyName, companyID, address, userId,
level, fName, lName, id, userName, password, email, phoneNumber`.

| # | condition | status | effect |
|---|---|---|---|
| 1 | any user field empty | 400 `"error in get value"` | nothing created |
| 2 | username / password invalid | 409 (one or two messages) | nothing created |
| 3 | user valid | **201 `"User Created"`** | user rows created (see `01` §3.4) |
| 3a | …then company fields empty | body gets `400 "error in get value of company"`, but status stays 201 | **user exists with no company and no permission** |
| 3b | …company fields present | body gets `201 "Company Created"` | `company` row + `permissions(newUser, company.id, 3)` |

**`AddPosition`** (`systemadmin/AddPosition.go`) — body: `companyName, userId, level` + user fields.
Same rows 1–3. Then `permissions(newUser, GetCompanyByName(companyName), 4)`. **No check that the
company exists**, so an unknown name yields `companyid = 0`. No final response beyond `201 User Created`.

### 3.2 Manager (gate: result ∉ {-1, 0}, so levels 1–4)

| endpoint | body (beyond userId, level) | success | other outcomes |
|---|---|---|---|
| `AddDriver` | user fields + `address`, `isEmployee` | 201 `"User Created"`; creates user + `permissions(u, myCompany, 5)` + `driver` + `driveravailable(available=true)` | 400 `"error in get value"`, 409 validation |
| `AddDrive` | `nameOfDriver` (**the drive's name**, misnamed), `timeToStartH`, `timeToStartM`, `gpsStart`, `gpsEnd` | **empty 200**; creates `drive` + `driveavailable(true)` | no field validation at all |
| `GetDrivers` | — | 200 `{"message":"Success","List":[{nameOfDriver,aviable}]}` from `driveravailable` of my company | `List` is **`null`** when empty |
| `GetDrives` | — | 200 `{"message":"Success","List":[{nameOfDrive,aviable}]}` from `driveavailable` of my company | `List` `null` when empty |
| `ViewOrNotDriver` | `driverName` | 200 `{"message":"good work"}`; toggles `driveravailable.available` for the first company driver whose first name matches | no match → `401 {}` |
| `DeleteDriver` | `driverName` | 200 `{"message":"Driver Is deleted"}`; deletes the `driveravailable` row only | no match → **empty 200** |
| `ViewOrNotDrive` | `driveName` | 200 `{"message":"good work"}`; toggles `driveavailable.available` | drive not in my company → `401 {}` |
| `DeleteDrive` | `driveName` | 200 `{"message":"Drive deleted"}`; deletes the `driveavailable` row only | drive not in my company → 401 `"User has no permissions"` |

Drive lookups use `GetDriveByName(name)`, the first drive with that name across **all** companies,
then check `drive.CompanyId == myCompany` (the ownership fix, see comments in `DeleteDrive.go`,
`DriveAviable.go`).

"Delete" never deletes the `drive`/`driver`/`permissions` rows. It only removes the availability
row, so the item disappears from the manager lists and from sorting. A deleted driver can still log
in.

### 3.3 Driver — `GetMyDrives` (`handler/driver/GetMyDrives.go`)

Body: `userId, level`. All failures are 401:

| # | condition | body |
|---|---|---|
| 1 | bind error | `"error in bind json"` |
| 2 | no `(userId, level)` permission row | `"User has no permissions"` |
| 3 | that row's company = 0 | `"you didnt have a company "` (trailing space) |
| 4 | company has no `sort` rows | `"you didnt have a Sort "` (trailing space) |
| 5 | `GetUserIdByCompanyIdAndLevel(level, company)` = 0 | `"User has no permissions"` |
| 6 | ok | 200 `{"message":"work","list of drives":[{Name,HToStart,MToStart,GpsStart,GpsEnd}]}` for the company's **latest** sort; `null` if none |

⚠️ Step 5 resolves "the driver" as **the first user with that level in the company**, not the
requester (`tmp.UserId` is never used after step 2). See §9.1.

### 3.4 Raw CRUD — `/MainController/{Company,Drive,DriveAvailable,Driver,DriverAvailable,Permissions}/…`

Generic CRUD, no permission check (`00` §3.3).

## 5. Client-Side Validation

- `CompanyRegistration.tsx`: sends only if all 10 company and user fields are non-empty. Otherwise
  it silently does nothing.
- `AddDriver.tsx`: no client validation. It always posts.
- `AddDrive.tsx`: requires a valid departure time (`alert("יש לבחור שעת יציאה")`), and sends hour and
  minute separately. No other checks.
- Confirm-password: never compared (see `01` §9.8).

## 6. UI States & Components

| route | component | calls | notes |
|---|---|---|---|
| `/systemAdministrator` | `mainHome` | — | cards → CompanyRegistration / AddPosition / NumberOfInlays / `/manager` |
| `/systemAdministrator/CompanyRegistration` | `CompanyRegistration` | `AddCompany` | ⚠️ **payload has no `userId`/`level`** → always 401 → `alert("שגיאה בהוספת חברה")` (§9.2) |
| `/systemAdministrator/AddPosition` | `addPosition` | none | company dropdown is **hardcoded** `["nave","yadai"]`, submit button has no `onClick` |
| `/systemAdministrator/NumberOfInlays` | `NumberOfInlays` | none | **hardcoded** demo data |
| `/manager` | `managerHome` | — | cards → AddDriver, EditDrivers, AddDrive, EditDrives, Sort, EditSort, `/driver` |
| `/manager/AddDriver` | `AddDriver` | `AddDriver` | `alert("נהג נוסף בהצלחה")` / `alert("שגיאה בהוספת נהג")` |
| `/manager/EditDrivers` | `EditDrivers` | `GetDrivers`, `ViewOrNotDriver`, `DeleteDriver` | switch per driver + delete. Toggle/delete responses aren't checked and the list isn't refreshed. |
| `/manager/AddDrive` | `AddDrive` | `AddDrive` | `alert("הנסיעה בוצעה בהצלחה")` / `alert("שגיאה בהוספת נסיעה")` |
| `/manager/EditDrives` | `EditDrives` | `GetDrives`, `ViewOrNotDrive`, `DeleteDrive` | delete alerts `"הנסיעה נמחקה בהצלחה"` / `"שגיאה במחיקת נסיעה"` |
| `/manager/EditSort` | `EditSort` | none | placeholder ("מסך זה נמצא בבנייה") |
| `/driver` | `DriverHome` | `GetMyDrives` | loading → table or `"אין נסיעות זמינות"`. Navigation buttons open Google Maps. ⚠️ crashes on a `null` list (§9.4) |

## 7. Cross-Cutting Flows

- Manager adds driver/drive → it shows up in EditDrivers/EditDrives: E2E `manager-crud.spec.ts`,
  `driver-crud.spec.ts`.
- Role gating for every route: E2E `permission-gating.spec.ts`.
- Manager creates a sort → driver sees assigned drives: see `03-sorting.md` (E2E is `fixme`).

## 8. Edge Cases / Test Scenarios

1. `AddCompany` from the real UI payload (no userId/level) → 401. Pin this until it's fixed (§9.2).
2. `AddCompany` with valid user but empty company fields → status 201, user created, no company and
   no permission.
3. `AddPosition` with a non-existent company name → permission row with `companyid = 0`.
4. Two drivers in one company with the same first name → toggle/delete acts on the first one only.
5. Two companies each with a drive named "X" → the manager of the company whose drive is *not* first
   by id gets 401 on toggle/delete of their own "X" (§9.3).
6. `DeleteDriver` with an unknown name → empty 200. The frontend treats it as success.
7. `GetDrivers`/`GetDrives`/`GetMyDrives` on an empty company → `List`/`list of drives` is `null`.
8. Two drivers in one company, sort assigns drives to driver B → driver B's `GetMyDrives` returns
   driver A's drives (§9.1).
9. Admin (level 1/2) opening manager screens → passes the gate. Company = the admin's permission-row
   company.

## 9. Known Discrepancies & Risks

| # | Item | STATUS |
|---|---|---|
| 9.1 | `GetMyDrives` returns the drives of the *first* driver in the company, not the requester | Bug candidate — high (every driver but one sees the wrong list) |
| 9.2 | `CompanyRegistration.tsx` omits `userId`/`level` → company registration from the UI always fails | Bug candidate — high |
| 9.3 | `GetDriveByName` isn't company-scoped. Duplicate drive names across companies block the legitimate owner. | Bug candidate |
| 9.4 | Go returns `null` for empty lists. `DriverHome` does `drives.length` on it and crashes (its test mocks `[]`, not `null`). `EditDrivers` survives only because the throw is caught. | Bug candidate |
| 9.5 | "Delete" only removes availability rows. Driver keeps login and permissions. | Needs decision (soft-delete may be intended) |
| 9.6 | Drivers identified by first name | Bug candidate (ambiguity) |
| 9.7 | `AddCompany` returns 201 even when the company isn't created | Bug candidate |
| 9.8 | `AddPosition` and `NumberOfInlays` screens are UI stubs with hardcoded data | Unfinished |
| 9.9 | `AddDrive` request field `nameOfDriver` is actually the drive name | Intentional (naming only). Don't rename without updating both sides. |
