# Frontend — AGENTS.md

> DRIVE-AWAY — React 18 + TypeScript + Vite + MUI 5. The UI is Hebrew-only.
> (The directory really is named `font/`, not `front/`.)

## Tech Stack

- **Build/dev:** Vite 5 (`npm run dev`, port 3000; `npm run build` = `tsc -b && vite build`)
- **UI:** MUI 5 (`@mui/material`, `@mui/icons-material`, `@mui/x-date-pickers` + `dayjs`)
- **Routing:** `react-router-dom` 6, all routes declared in `src/App.tsx`
- **HTTP:** `axios`, called directly from components (no API client module)
- **Maps:** `mapbox-gl`, loaded lazily by `component/common/LocationField` (map picker for location fields). Needs the *public* token `VITE_MAPBOX_TOKEN` at build/dev time; without it the map button is disabled and typing still works. Location values are `"lat,lng"` or free-text addresses (the backend accepts both).
- **Testing:** Vitest (jsdom) + React Testing Library + `user-event`, setup in `src/setupTests.ts`

## Layout

```
src/main.tsx                ThemeProvider(theme) + CssBaseline + <App/>
src/App.tsx                 AppBar + menu + <Routes>, one route per screen
src/GlobalState.tsx         Context: level, userId, nameOfPage, each mirrored to localStorage
src/theme.ts                THE source of palette/shape/typography tokens
src/types.ts                shared request/response + form types
src/LogIn/                  LogIn (mounted), SingUp + MainLogIn (NOT mounted by any route), NotFound
src/component/common/       ActionCard, LoadingScreen, the shared building blocks
src/component/systemAdministrator/   level 1–2 screens
src/component/manager/      level 1–4 screens (company manager)
src/component/drivers/      driver screen (DriverHome)
```

## Conventions

### Auth state and route gating

- Login (`src/LogIn/LogIn.tsx`) posts to `POST http://localhost:2909/login` and stores
  `response.data.level` / `.UserId` in `GlobalState`, which persists them to `localStorage`. There
  is no token. The backend trusts the `{userId, level}` that each screen sends in the request body
  (see `../specs/00-shared-conventions.md` §3).
- Screens are gated by HOCs wrapping the default export. A disallowed level gets
  `<Navigate to="/LogIn" />`:
  - `systemAdministratorPermissions`: `level === 1 || level === 2`
  - `managerPermissions`: `level > 0 && level < 5` (so admins can open manager screens too)
  - `driverPermissions`: `level > 0 && level < 6`
  New screens must be wrapped the same way. This is client-side only and doesn't protect anything.
- Each screen calls `setNameOfPage("<Hebrew title>")` in a `useEffect` to set the AppBar title.

### API calls

- The backend URL is **hardcoded** as `http://localhost:2909/...` in every component. There is no
  env var and no shared base URL. Keep new calls consistent with that until someone deliberately
  introduces a config, and don't half-migrate.
- Business endpoints are all `POST` with `{ ...fields, userId, level }` in the body.
- User feedback is `alert("<Hebrew>")` on success or failure. Tests spy on `window.alert`.

### Styling

- `src/theme.ts` is the single source of colors, radius, and typography. Use `sx` with theme keys
  (`color="primary"`, `bgcolor: 'background.default'`, spacing numbers). Don't hardcode hex values
  in components. Screens use the same card pattern: `Stack` → `Paper elevation={3}`
  (`maxWidth` 420–960, `borderRadius: 3`) → icon + `Typography variant="h6"` header → form `Stack`.
- All user-facing text is Hebrew, inline in the component (there's no strings file).

## Testing (binding: see root `CLAUDE.md` testing policy)

- Run: `cd font && npm test` (uses `cross-env NODE_OPTIONS=--no-experimental-webstorage vitest run`.
  That flag is needed so jsdom's `localStorage` is used instead of Node's built-in one).
- Tests are colocated: `Foo.tsx` ↔ `Foo.test.tsx`. Every component currently has one.
- Standard harness (see `src/component/manager/AddDriver.test.tsx`):
  `vi.mock('axios')` + `vi.mocked(axios, true)`; seed `localStorage.setItem('level', ...)` **before**
  render (GlobalState reads it in its initializer); wrap in
  `<GlobalStateProvider><MemoryRouter>…</MemoryRouter></GlobalStateProvider>`; `localStorage.clear()`
  + `vi.clearAllMocks()` in `beforeEach`.
- Query by what users see: `getByLabelText('<Hebrew label>')`, `getByRole('button', { name: '...' })`.
  Assert on the axios payload and the alert text, not on internal state.
- Type-check before calling frontend work done: `npx tsc -b` (the same check `npm run build` runs first).
