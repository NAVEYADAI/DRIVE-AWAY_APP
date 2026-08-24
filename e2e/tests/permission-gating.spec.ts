import { test, expect } from '@playwright/test';

// Confirms the client-side permission HOCs (managerPermissions,
// systemAdministratorPermissions) actually protect routes end-to-end, not
// just in isolated unit tests.
test.describe('permission gating', () => {
  test('a driver-level user cannot reach manager routes', async ({ page }) => {
    page.on('dialog', (d) => d.accept());
    await page.goto('/LogIn');
    await page.getByLabel('שם משתמש').fill('e2e-driver');
    await page.getByLabel('סיסמה').fill('e2e-driver-pass123');
    await page.getByRole('button', { name: 'כניסה' }).click();

    // Driver-level users should land on /driver, not /manager.
    await expect(page).toHaveURL(/\/driver$/);

    // Directly navigating to a manager-only route redirects back to /LogIn.
    await page.goto('/manager');
    await expect(page).toHaveURL(/\/LogIn$/);
  });

  test('a driver-level user cannot reach system administrator routes', async ({ page }) => {
    page.on('dialog', (d) => d.accept());
    await page.goto('/LogIn');
    await page.getByLabel('שם משתמש').fill('e2e-driver');
    await page.getByLabel('סיסמה').fill('e2e-driver-pass123');
    await page.getByRole('button', { name: 'כניסה' }).click();
    await expect(page).toHaveURL(/\/driver$/);

    await page.goto('/systemAdministrator');
    await expect(page).toHaveURL(/\/LogIn$/);
  });

  test('a manager-level user cannot reach system administrator routes', async ({ page }) => {
    page.on('dialog', (d) => d.accept());
    await page.goto('/LogIn');
    await page.getByLabel('שם משתמש').fill('e2e-manager');
    await page.getByLabel('סיסמה').fill('e2e-manager-pass123');
    await page.getByRole('button', { name: 'כניסה' }).click();
    await expect(page).toHaveURL(/\/manager$/);

    await page.goto('/systemAdministrator');
    await expect(page).toHaveURL(/\/LogIn$/);
  });

  test('an admin-level user can reach system administrator routes', async ({ page }) => {
    page.on('dialog', (d) => d.accept());
    await page.goto('/LogIn');
    await page.getByLabel('שם משתמש').fill('e2e-admin');
    await page.getByLabel('סיסמה').fill('e2e-admin-pass123');
    await page.getByRole('button', { name: 'כניסה' }).click();

    await expect(page).toHaveURL(/\/systemAdministrator$/);
  });
});
