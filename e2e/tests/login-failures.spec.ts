import { test, expect } from '@playwright/test';

test.describe('login failure paths', () => {
  test('wrong password shows an error and does not navigate away', async ({ page }) => {
    page.on('dialog', (d) => d.accept());
    await page.goto('/LogIn');

    await page.getByLabel('שם משתמש').fill('e2e-manager');
    await page.getByLabel('סיסמה').fill('definitely-wrong-password');
    await page.getByRole('button', { name: 'כניסה' }).click();

    await expect(page).toHaveURL(/\/LogIn$/);
  });

  test('nonexistent username shows an error and does not navigate away', async ({ page }) => {
    page.on('dialog', (d) => d.accept());
    await page.goto('/LogIn');

    await page.getByLabel('שם משתמש').fill('no-such-user-e2e');
    await page.getByLabel('סיסמה').fill('whatever123');
    await page.getByRole('button', { name: 'כניסה' }).click();

    await expect(page).toHaveURL(/\/LogIn$/);
  });
});
