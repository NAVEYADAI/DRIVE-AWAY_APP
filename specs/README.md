# specs/

Exact specification documents for the app's features, written by reading the code directly (not
from memory or from the original intent). Backend and frontend go together in one file per domain,
so each file is the single source of truth for deriving tests.

## How to use these files

Each file follows a fixed template. How each section maps to a test type:

| section | what to derive from it |
|---|---|
| §3 API Contract (error-branch table) | backend unit/handler tests: each row = at least one test case |
| §4 Business Logic | unit tests for services and algorithms (`service/`, `sorting/`) |
| §5 Client-Side Validation | frontend unit tests for validation |
| §6 UI States & Components | frontend unit tests for rendering states |
| §7 Cross-Cutting Flows | E2E scenarios (`e2e/tests/`) |
| §8 Edge Cases / Test Scenarios | the cases that must be covered, the basis for writing the actual test files |
| §9 Known Discrepancies & Risks | cases where a test has to decide: pin the current behavior (STATUS: Intentional) or mark it as a known issue until fixed (STATUS: Bug candidate) |

**Important:** error messages, JSON keys, and field names in these files are copied **verbatim**
from the code, including misspellings (`messege`, `aviable`, `SingUp`) and trailing spaces. When the
code changes, update the matching file along with it. These files document what the code does, not
what it was meant to do.

**Relationship to the testing policy** (root `CLAUDE.md`): fixing a "Bug candidate" from §9 is a bug
fix like any other. It requires a regression test that fails before the fix, and it updates the row
here (STATUS → `Fixed <date>`, plus a pointer to the test).

## Existing files

- [`00-shared-conventions.md`](./00-shared-conventions.md): identity (no tokens, client-asserted
  `userId`/`level`), permission levels and the permission gate, response/error conventions, frontend
  API patterns. The other files point here.
- [`01-auth.md`](./01-auth.md): users and the tables they're split across, the three login endpoints,
  throttling, sign-up rules, change username/password, forgot-username, mail.
- [`02-management.md`](./02-management.md): companies, positions, drivers, drives, availability,
  "my drives", and all the admin/manager/driver screens.
- [`03-sorting.md`](./03-sorting.md): the assignment algorithm (`POST /Sort`), Google Maps
  dependency, backtracking behavior.

## What's still missing

Nothing big is undocumented. If a new domain is added, number it `04-`, `05-`, and so on, and add it
to the list above.
