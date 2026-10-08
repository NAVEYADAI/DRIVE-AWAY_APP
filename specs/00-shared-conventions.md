# Shared Conventions

## 1. Overview

Conventions that apply across every backend domain (`auth`, `management`, `sorting`) and every
frontend screen: how "authentication" works (there is none on the server), how the permission level
is checked, how errors are returned, and how the frontend calls the API. The other spec files point
here and don't repeat this.

Sources: `back/cmd/server/main.go`, `back/internal/auth/auth.go`,
`back/internal/management/service/permissions.go`, `back/internal/management/model/permissions.go`,
`font/src/GlobalState.tsx`, `font/src/LogIn/LogIn.tsx`, `font/src/component/*/…Permissions.tsx`.

Last checked against the code: 2026-09-29.

## 2. Data Model

`permissions` (`back/internal/management/model/permissions.go`):
- `ID int` — `id`
- `UserId int` — `userid`
- `ConpanyId int` — `companyid` (Go field name is misspelled, column is not)
- `Level int` — `level`

**Level values** (derived from who creates what; nothing enforces them as an enum):

| level | role | created by |
|---|---|---|
| 1, 2 | System administrator | Seeded manually / directly in DB (no endpoint creates them) |
| 3 | Company manager | `AddCompany` (`systemadmin/AddCompany.go:75`) |
| 4 | Position in an existing company | `AddPosition` (`systemadmin/AddPosition.go:65`) |
| 5 | Driver | `AddDriver` (`manager/AddDriver.go:68`) |

One user can in principle have several `permissions` rows. Most lookups are keyed by
`(userid, level)`. `GetPermissionsByUserId` / `GetPermissionLevelByUserId` just take whichever row
GORM `Find` returns first.

## 3. Auth Mechanism — client-asserted identity

**There is no server-side session, token, or cookie.** Login returns `UserId` and `level`. The
frontend persists them in `localStorage` (`GlobalState.tsx`) and sends them back in the JSON body of
every business request as `{ "userId": <int>, "level": <int> }` (`handler.UserIdAndLevel`).

### 3.1 The permission gate — `ReturnPermissionForSingUp(UserId, level)` (`management/service/permissions.go`)

```
row = GetPermissionByUserIdAndLevel(UserId, level)   // WHERE userid = ? AND level = ?
if row.Level != 0:
    level 1 or 2 → return 1
    level 3 or 4 → return 2
    otherwise    → return 0
return -1                                            // no row
```

Callers:
- Admin endpoints (`AddCompany`, `AddPosition`): proceed only when the result `== 1`.
- Manager endpoints (`AddDriver`, `AddDrive`, `DeleteDrive`, `DeleteDriver`, `GetDrives`,
  `GetDrivers`, `ViewOrNotDrive`, `ViewOrNotDriver`, `POST /Sort`): reject `-1` **and** `0`, so admins
  (1) and managers (2) pass and drivers (0) don't.
- `GetMyDrives` (driver) **does not** use it. It checks
  `GetPermissionByUserIdAndLevel(...).Level == 0` directly (see the NOTE in `GetMyDrives.go`: using the
  helper there locked every real driver out, because drivers map to `0`).

Rejection body (all of the above): `401 {"message": "User has no permissions"}`.

### 3.2 Company resolution

Manager endpoints resolve "my company" with `GetCompanyIdByLevelAndUserId(userId, level)` → the
`companyid` of the `(userid, level)` permissions row. For admins (level 1/2) this is whatever company
id their permission row holds (often 0).

### 3.3 Security consequences (STATUS: Bug candidate — security)

- Any client can send any `userId`/`level`. Identity is not verified after login. Knowing a manager's
  numeric id is enough to act as them.
- The raw CRUD route groups (`/loginApi/{Mail,Password,PersonalInformation,Phone,TryLogin,Users}/…`
  and `/MainController/{Company,Drive,DriveAvailable,Driver,DriverAvailable,Permissions}/…`) have
  **no permission check at all**. For example, `GET /loginApi/Password/GetAllPassword` returns every
  stored password, and `POST /MainController/Permissions/CreatePermission` can grant any level.
- Passwords are stored and compared **in plaintext** (`password.password`, compared with `!=` in the
  login handlers). Every login attempt, including the typed password, is also stored in plaintext in
  `trylogin.password`.
- `back/internal/auth/handler/SendMail.go:14` contains a hardcoded Gmail sender address and
  app password, committed to git.
- CORS (`main.go`) only allows `http://localhost:<port>` origins. The deployed frontend's origin
  would be rejected by the browser unless it is also on localhost (worth confirming against the
  ECS setup).

## 4. Error / response conventions

- Success and error bodies: `c.IndentedJSON(status, gin.H{"message": "..."})`.
- **`POST /login` (`auth.Login`) uses the key `"messege"`** (sic) for every response. The frontend
  `LoginResponse` interface mirrors that spelling. `/loginApi/LogIn` and `/MainController/LogIn` use
  `"message"`.
- JSON bind failures: `415 {"message": "error in BindJSON"}` in auth handlers;
  `400 {"message": "error in BindJSON"}` in management handlers and raw CRUD routes; `401` with
  `"error in bind json"` in `GetMyDrives`.
- **Double-write pattern.** Handlers that validate username + password (`CreateUser`, `AddCompany`,
  `AddPosition`, `AddDriver`) call `IndentedJSON` once per failing check before returning, and
  `AddCompany`/`AddDriver` write `201 User Created` and then continue to write more. Gin sends the
  **first** status code. The body is the concatenation of every written JSON object. Tests should
  assert the status and that the body *contains* the expected message.
- Several success paths write nothing at all (`AddDrive`, `AddPosition`, `ForgetMyUserName`,
  raw `CreateUser`), which gives an empty `200`.

## 5. Frontend API conventions

- Base URL is the hardcoded literal `http://localhost:2909` in each component.
- All business calls are `axios.post(url, { ...fields, userId, level })`.
- Success/failure feedback is `alert("<Hebrew>")`. Axios rejects on any non-2xx, so every 4xx/5xx
  shows the component's generic failure alert. Backend messages are never shown to the user.

## 6. Client-side gating

| HOC | allows | file |
|---|---|---|
| `systemAdministratorPermissions` | `level === 1 \|\| level === 2` | `font/src/component/systemAdministrator/systemAdministratorPermissions.tsx` |
| `managerPermissions` | `level > 0 && level < 5` | `font/src/component/manager/managerPermissions.tsx` |
| `driverPermissions` | `level > 0 && level < 6` | `font/src/component/drivers/DriverPermission.tsx` |

Anything else → `<Navigate to="/LogIn" />`. Logout (`App.tsx::logOut`) sets `level=0`, `userId=-1`.
The only protection this gives is hiding UI. See §3.3.

## 8. Edge Cases / Test Scenarios (cross-cutting)

1. Manager endpoint with a `(userId, level)` pair that has no permissions row → `401 "User has no permissions"`.
2. Manager endpoint called by a driver (level 5 row exists) → `401` (helper returns `0`).
3. Manager endpoint called by an admin (level 1/2) → passes the gate. The company resolves to the
   admin row's `companyid`.
4. `POST /login` response uses `messege`, and `/loginApi/LogIn` uses `message`. `/loginApi/LogIn` is
   pinned by `back/internal/auth/handler/LogIn_test.go`. **`auth.Login` (`POST /login`, the endpoint
   the frontend actually uses) has no backend test.** Only the frontend `LogIn.test.tsx` (mocked
   axios) and the E2E login specs cover it.
5. Double-write handlers: a sign-up with both an invalid username and an invalid password → status
   409, body contains both messages.

## 9. Known Discrepancies & Risks

| # | Item | STATUS |
|---|---|---|
| 9.1 | No server-side identity; `userId`/`level` trusted from body (§3.3) | Bug candidate — security |
| 9.2 | Raw CRUD groups unauthenticated, including `GetAllPassword` | Bug candidate — security |
| 9.3 | Plaintext passwords in `password` and `trylogin` | Bug candidate — security |
| 9.4 | Hardcoded Gmail app password in `SendMail.go` (in git history) | Bug candidate — security, **rotate the credential** |
| 9.5 | `"messege"` spelling on `POST /login` | Intentional (frontend depends on it) |
| 9.6 | CORS allows only localhost origins | Needs confirmation against the deployment |
