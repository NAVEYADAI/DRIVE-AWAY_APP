import { test, expect } from '@playwright/test';

// Golden-path smoke test: a seeded user can log in and land on the correct
// home screen for their permission level. Seeding happens against db_test
// only (see ../README or CLAUDE.md) — never against the dev/prod database.
test('a manager-level user can log in and reach the manager home screen', async ({ page }) => {
  await page.goto('/LogIn');

  await page.getByLabel('שם משתמש').fill('e2e-manager');
  await page.getByLabel('סיסמה').fill('e2e-manager-pass123');
  await page.getByRole('button', { name: 'כניסה' }).click();

  await expect(page).toHaveURL(/\/manager$/);
  // The page title renders both in the AppBar and as the page's own
  // heading, so scope to the heading to avoid a strict-mode ambiguous match.
  await expect(page.getByRole('heading', { name: 'ניהול חברה' })).toBeVisible();
});
