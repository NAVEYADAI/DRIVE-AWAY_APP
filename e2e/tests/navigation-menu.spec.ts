import { test, expect } from '@playwright/test';

// Regression test for a previously-fixed bug: the hamburger menu's items
// used to only close the menu without navigating anywhere. Proves the menu
// items genuinely navigate/log out, not just that the menu opens.

test.beforeEach(async ({ page }) => {
  page.on('dialog', (d) => d.accept());
  await page.goto('/LogIn');
  await page.getByLabel('שם משתמש').fill('e2e-manager');
  await page.getByLabel('סיסמה').fill('e2e-manager-pass123');
  await page.getByRole('button', { name: 'כניסה' }).click();
  await expect(page).toHaveURL(/\/manager$/);
});

test('"שיבוץ נסיעה" menu item navigates to the sort-creation screen', async ({ page }) => {
  await page.getByRole('button', { name: 'menu' }).click();
  await page.getByText('שיבוץ נסיעה').click();
  await expect(page).toHaveURL(/\/manager\/Sort$/);
});

test('"דף הבית" menu item navigates back to the role home screen', async ({ page }) => {
  await page.goto('/manager/AddDriver');
  await page.getByRole('button', { name: 'menu' }).click();
  await page.getByText('דף הבית').click();
  await expect(page).toHaveURL(/\/manager$/);
});

test('"התנתקות" menu item logs out and returns to /LogIn', async ({ page }) => {
  await page.getByRole('button', { name: 'menu' }).click();
  await page.getByText('התנתקות').click();
  await expect(page).toHaveURL(/\/LogIn$/);

  // Confirm the session really was cleared, not just a client-side redirect:
  // navigating straight to a gated route afterwards must bounce back to /LogIn.
  await page.goto('/manager');
  await expect(page).toHaveURL(/\/LogIn$/);
});
