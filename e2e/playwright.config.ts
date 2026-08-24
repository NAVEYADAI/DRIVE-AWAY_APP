import { defineConfig, devices } from '@playwright/test';

// E2E runs against the real stack (frontend + backend) wired to the
// isolated db_test database — see ../docker-compose.e2e.yml and
// ../CLAUDE.md. Never point this at the dev `db` or production RDS.
export default defineConfig({
  testDir: './tests',
  globalSetup: './global-setup.ts',
  fullyParallel: false,
  workers: 1,
  retries: process.env.CI ? 2 : 0,
  reporter: 'list',
  use: {
    baseURL: process.env.BASE_URL ?? 'http://localhost:3000',
    trace: 'retain-on-failure',
    screenshot: 'only-on-failure',
  },
  projects: [
    { name: 'chromium', use: { ...devices['Desktop Chrome'] } },
  ],
});
