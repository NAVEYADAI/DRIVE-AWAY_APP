import { test, expect } from '@playwright/test';

test.beforeEach(async ({ page }) => {
  page.on('dialog', (d) => d.accept());
  await page.goto('/LogIn');
  await page.getByLabel('שם משתמש').fill('e2e-manager');
  await page.getByLabel('סיסמה').fill('e2e-manager-pass123');
  await page.getByRole('button', { name: 'כניסה' }).click();
  await expect(page).toHaveURL(/\/manager$/);
});

test('adding a driver makes it appear in the drivers list', async ({ page }) => {
  const uniqueSuffix = Date.now().toString().slice(-6);
  const firstName = `e2eDriver${uniqueSuffix}`;

  await test.step('fill and submit the add-driver form', async () => {
    await page.goto('/manager/AddDriver');
    await page.getByLabel('שם פרטי').fill(firstName);
    await page.getByLabel('שם משפחה').fill('Test');
    await page.getByLabel('תעודת זהות').fill(uniqueSuffix);
    await page.getByLabel('שם משתמש').fill(`user${uniqueSuffix}`);
    await page.getByLabel('סיסמה', { exact: true }).fill('DriverPass123!');
    await page.getByLabel('אימות סיסמה').fill('DriverPass123!');
    await page.getByLabel('כתובת מייל').fill(`e2e${uniqueSuffix}@example.com`);
    await page.getByLabel('מספר טלפון').fill('0500000000');
    await page.getByLabel('כתובת', { exact: true }).fill('Test Address 1');
    await page.getByRole('button', { name: 'שליחה' }).click();
  });

  await test.step('the new driver shows up in the drivers list', async () => {
    await page.goto('/manager/EditDrivers');
    await expect(page.getByText(firstName)).toBeVisible({ timeout: 10_000 });
  });
});
