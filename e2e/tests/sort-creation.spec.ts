import { test, expect } from '@playwright/test';

// The sorting algorithm (back/internal/sorting/mainSort.go) unconditionally
// calls FindWay() for every driver/drive pairing, which makes a real HTTP
// call to the Google Maps Distance Matrix API (back/internal/sorting/GoogleMapApi.go).
// No API key is configured for the E2E backend (MY_GOOGLE_API_KEY is unset
// in docker-compose.e2e.yml) and this environment has no route to the real
// Google API either way, so the call fails and Sort() always returns "we
// didnt can sort" here — confirmed via direct backend logs, not a bug in
// this test or in the seeded fixture data (one available drive + driver is
// seeded in global-setup.ts specifically for this test). Fixing this
// properly means either mocking/injecting FindWay's HTTP client for E2E or
// providing a real (test) API key — flagged for follow-up rather than
// worked around here.
test.fixme('a manager can create a sort/assignment', async ({ page }) => {
  page.on('dialog', (d) => d.accept());
  await page.goto('/LogIn');
  await page.getByLabel('שם משתמש').fill('e2e-manager');
  await page.getByLabel('סיסמה').fill('e2e-manager-pass123');
  await page.getByRole('button', { name: 'כניסה' }).click();
  await expect(page).toHaveURL(/\/manager$/);

  await page.goto('/manager/Sort');
  await page.getByLabel('שם השיבוץ').fill('שיבוץ-בדיקה-E2E');

  const dialogPromise = page.waitForEvent('dialog');
  await page.getByRole('button', { name: 'שליחה' }).click();
  const dialog = await dialogPromise;
  expect(dialog.message()).toBe('השיבוץ נוצר בהצלחה');
  await dialog.accept();
});
