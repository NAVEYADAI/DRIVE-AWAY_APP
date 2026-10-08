# Auth — login, sign-up, credentials

## 1. Overview

Everything under `back/internal/auth/`: user creation (sign-up and the shared "create a user" steps
reused by management), the three login endpoints, login throttling, change username/password,
forgot-username, and the mail sender. Plus the frontend login screen and the (unmounted) sign-up
screen.

See `00-shared-conventions.md` for the identity model (no tokens) and response conventions.

Sources: `back/internal/auth/{auth.go,handler/*,service/*,model/*}`,
`back/internal/management/handler/LogIn.go`, `font/src/LogIn/*`, `font/src/component/AddUser.tsx`.

Last checked against the code: 2026-09-29.

## 2. Data Model

A "user" is split across tables keyed by user id (`back/internal/auth/model/`):

| table | struct | key columns | notes |
|---|---|---|---|
| `users` | `Users` | `id`, `userName` | the username |
| `password` | `Password` | `id`, `userid`, `password` | **append-only history**. The current password is the row with the highest `id` (`GetLastPasswordByUserId`: `ORDER BY id DESC`). Plaintext. |
| `mail` | `Mail` | `id`, `userid`, `mail` | |
| `phone` | `Phone` | `id`, `userid`, `phone` | |
| `personalinformation` | `PersonalInformation` | `id`, `fname`, `lname`, `identitycard` (text), `userid` | |
| `doubleregistration` | `DoubleRegistration` | `id`, `userid`, `personalinformationid` | recorded when a new user registers with an ID card that already exists |
| `name` | `Name` | `id`, `idUser`, `name` | **username history**. `ChangeUsername` appends the *old* name here. |
| `trylogin` | `TryLogIn` | `id`, `userName`, `password`, `ip`, `macAddress`, `Cause`, `connect`, `date` | every login attempt, including the typed password in plaintext |

`TryLogIn.Cause` codes (comment in `trylogin.go`): `-1` success, `1` bind error, `2` too many
attempts, `3` username not found, `4` wrong password.

## 3. API Contract

### 3.1 `POST /login` — `auth.Login` (`auth/auth.go:24`) — **used by the frontend**

Body: `{ "UserName", "Password", "IP", "MacAddress" }` (all strings).

| # | condition | status | body | trylogin row |
|---|---|---|---|---|
| 1 | JSON bind fails | 415 | `{"messege":"error in BindJSON"}` | connect=false, Cause=1 |
| 2 | `TryLogInIpAndMAcAddress` fails (§4.2) | 409 | `{"messege":"Too many attempts"}` | false, 2 |
| 3 | `GetUserByUserName` → -1 | 404 | `{"messege":"User not found"}` | false, 3 |
| 4 | latest password `!= Password` | 401 | `{"messege":"Password incorrect"}` | false, 4 |
| 5 | password OK, `GetPermissionsByUserId(id).GetId() != 0` | 200 | `{"messege":"LogIn Success","UserId":<int>,"level":<int>}` | true, -1 |
| 6 | password OK, no permissions row | 401 | `{"messege":"User has no permissions"}` | **true, -1** (recorded as a success before the permission check) |

### 3.2 `POST /loginApi/LogIn` — `handler.LogInUser` (`auth/handler/LogIn.go`)

Same steps 1–4 with key `"message"`, and statuses 415/409/404/401. Success:
`200 {"message":"LogIn Success","UserId":<int>}` (no level, no permission check).

### 3.3 `POST /MainController/LogIn` — `management/handler.LogIn`

Same as 3.2, but success adds `"level": GetPermissionLevelByUserId(id)` (0 if none, still 200).

### 3.4 `POST /loginApi/SingUp` — `handler.CreateUser` (`auth/handler/SingUp.go`)

Body: `{ "UserName","FName","LName","IdentityCard","Password","Email","Phone" }`.

| # | condition | status | body |
|---|---|---|---|
| 1 | bind fails | 415 | `{"message":"error in BindJSON"}` |
| 2 | username invalid (§4.1) | 409 | `{"message":"<username message>"}` (continues to check the password) |
| 3 | password weak (§4.3) | 409 | `{"message":"<password message>"}` (may follow #2 in the same body) |
| 4 | both valid | 201 | `{"message":"User Created"}` |

On success, inserts: `users`, **two** `personalinformation` rows (see §9.3), `password`, `mail`,
`phone`, and a `doubleregistration` row if the ID card already existed. No permissions row, so the
new user **cannot log in via `POST /login`** (§3.1 row 6).

The same create-user sequence is copy-pasted into `AddCompany`, `AddPosition`, and `AddDriver`
(see `02-management.md`).

### 3.5 `POST /loginApi/ChangeUsername`, `POST /loginApi/ChangePassword`

Both re-authenticate by making an HTTP `POST` to the hardcoded
`http://localhost:8080/logIn/LogIn`. The server listens on `:2909` and the route is
`/loginApi/LogIn`, so the request fails and the response is `500 "Failed to send POST request"`
(a JSON string, not an object). See §9.4.

### 3.6 `POST /loginApi/ForgetMyUserName`

Body `{ "Phone","Mail","IdentityCard":int }`. Looks up the three user ids. The `if` body is empty,
so it always returns an **empty 200**. Bind error → 415.

### 3.7 `POST /loginApi/SendMail`

Ignores the request body and sends a fixed test email ("Drive Away" / "This is a test message from")
from a hardcoded Gmail account to a hardcoded recipient. Errors are swallowed and the response is an
empty 200. Unauthenticated, so anyone can trigger it.

### 3.8 Raw CRUD — `/loginApi/{Mail,Password,PersonalInformation,Phone,TryLogin,Users}/…`

Generic Get/GetAll/Create/Update/Delete per table, bodies are the model structs, bind error → 400.
No permission check (see `00-shared-conventions.md` §3.3).

## 4. Business Logic

### 4.1 Username rules — `service.SignUpNewUserName` (checked in this order, first failure wins)

1. already exists in `users` → `"User name already taken"`
2. length not in [5, 20] → `"User name must be between 5 and 20 characters"`
3. not `^[a-zA-Z0-9]+$` → `"User name must contain only letters and numbers"`
4. matches `\d{4}-\d{2}-\d{2}` → `"User name must not contain personal information"`. **Unreachable**,
   because rule 3 already rejects `-`.
5. first char not ASCII letter → `"User name must start with a letter"`
6. else → `(true, "Good user name")`

### 4.2 Login throttling — `service.TryLogInIpAndMAcAddress(IP, MacAddress, now)`

Counts **all** `trylogin` rows ever recorded with `ip = IP AND "macAddress" = MacAddress` and
`connect = false`. If the count is `> 5` → `"Too many attempts"`. There is **no time window** (the
10-minute filter is commented out in `GetTryLogInByIpAndMacAddress`), and a successful login doesn't
reset the count.

**The frontend sends `IP: <username>` and `MacAddress: <password>`** (`LogIn.tsx:33-34`). So the
counter is keyed on the (username, typed-password) pair:
- repeating the *same* wrong password 6 times locks that exact pair forever;
- trying a *different* password each time never trips the limit (no brute-force protection);
- the correct password is never blocked, because failed attempts with it can't exist.

### 4.3 Password strength — `service.IsSecure` (order, first failure wins)

1. `len < 8` → `"Password must be at least 8 characters long"`
2. no uppercase → `"Password must contain at least one uppercase letter"`
3. no lowercase → `"Password must contain at least one lowercase letter"`
4. no digit → `"Password must contain at least one digit"`
5. none of `!@#$%^&*()-_+=[]{}|;:'",.<>?/` → `"Password must contain at least one special character"`
6. else → `(true, "Password is secure")`

## 5. Client-Side Validation

- `LogIn.tsx`: sends only if both username and password are non-empty. Otherwise the click does
  nothing (no message).
- `AddUser.tsx` (shared form for user fields): has a "confirm password" field (`Password2Value`), but
  **no component compares it with `PasswordValue`**. Nothing is enforced client-side, and every rule
  in §4 is checked server-side only.

## 6. UI States & Components

- `/LogIn` → `LogIn`: fields `שם משתמש`, `סיסמה`, button `כניסה`. On 2xx it navigates by level (1–2 →
  `/systemAdministrator`, 3–4 → `/manager`, 5 → `/driver`, anything else stays put). On any error it
  shows `alert("שגיאת התחברות")`.
- `SingUp.tsx` / `MainLogIn.tsx` exist, but **no route mounts them**, and `SingUp`'s submit button
  has no `onClick`. Pinned by `e2e/tests/signup-not-reachable.spec.ts`.
- There's no wildcard route. Unknown paths render an empty main area (`NotFound.tsx` exists but
  isn't routed).

## 7. Cross-Cutting Flows

- **Login → role home**: E2E `login.spec.ts`, `login-failures.spec.ts`, `permission-gating.spec.ts`.
- **Logout**: menu `התנתקות` → level 0, userId -1 → `/LogIn`. E2E `navigation-menu.spec.ts`.

## 8. Edge Cases / Test Scenarios

1. `POST /login` happy path for each level → 200 with `messege`, `UserId`, `level`. **No backend test
   exists for `auth.Login`** (only for `LogInUser`).
2. `POST /login` valid credentials, no permissions row → 401 `"User has no permissions"`, and a
   `trylogin` row with `connect=true`.
3. Password history: after two `AddPassword` calls, only the newest one logs in.
4. User with **no** password row (e.g. created via raw `/loginApi/Users/CreateUser`) → latest password
   is `""`, so a login with `Password: ""` **succeeds** on the backend (the frontend blocks empty
   input, the API doesn't).
5. Throttle: 6 identical failed pairs → 7th attempt 409; a different password on the 7th attempt → not
   throttled.
6. Sign-up with both invalid username and weak password → 409, body contains both messages.
7. Sign-up with an ID card that already exists → `doubleregistration` row pointing at the existing
   `personalinformation.id`.
8. ChangeUsername / ChangePassword → currently always 500 (§3.5).

## 9. Known Discrepancies & Risks

| # | Item | STATUS |
|---|---|---|
| 9.1 | Throttle keyed on (username, password) from the frontend, no time window (§4.2) | Bug candidate |
| 9.2 | `POST /login` records a successful `trylogin` before rejecting a user with no permissions | Bug candidate (minor) |
| 9.3 | `CreateUser`/`AddCompany`/`AddPosition`/`AddDriver` call `service.AddPersonalInformation` (which inserts) **and then** `model.AddPersonalInformation` → 2 `personalinformation` rows per user | Bug candidate |
| 9.4 | ChangeUsername/ChangePassword call `localhost:8080/logIn/LogIn` (wrong port and path). If that ever answered, `Response.UserId` is typed `string` while the login returns an int, and the decode error hits `log.Fatal`, **killing the server** | Bug candidate |
| 9.5 | `ForgetMyUserName` is a stub (empty 200) | Unfinished |
| 9.6 | `SendMail` is a test stub with hardcoded credentials (see `00` §9.4) | Bug candidate — security |
| 9.7 | Username rule 4 unreachable | Intentional-ish (dead code, harmless) |
| 9.8 | Confirm-password field never compared | Bug candidate |
| 9.9 | Empty-password login for users with no password row (§8.4) | Bug candidate — security |
