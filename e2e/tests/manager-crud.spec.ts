import { test, expect } from '@playwright/test';

async function loginAsManager(page: import('@playwright/test').Page) {
  await page.goto('/LogIn');
  await page.getByLabel('שם משתמש').fill('e2e-manager');
  await page.getByLabel('סיסמה').fill('e2e-manager-pass123');
  await page.getByRole('button', { name: 'כניסה' }).click();
  await expect(page).toHaveURL(/\/manager$/);
}

// Adding a driver and confirming it shows up is covered by driver-crud.spec.ts.
test.describe('manager CRUD flow', () => {
  test('adding a drive makes it appear in the drives list', async ({ page }) => {
    page.on('dialog', (d) => d.accept());
    await loginAsManager(page);

    await page.goto('/manager/AddDrive');
    await page.getByLabel('שם הקו').fill('קו-בדיקה-E2E');
    await page.getByLabel('שעת יציאה').fill('08:00 AM');
    await page.getByLabel('מיקום התחלה').fill('31.7,34.8');
    await page.getByLabel('מיקום סיום').fill('31.8,34.9');
    await page.getByRole('button', { name: 'שליחה' }).click();

    await page.goto('/manager/EditDrives');
    await expect(page.getByText('קו-בדיקה-E2E')).toBeVisible();
  });
});
