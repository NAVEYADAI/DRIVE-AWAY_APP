import { test, expect } from '@playwright/test';

// FINDING (not a bug per se, documenting current state): font/src/LogIn/SingUp.tsx
// exists and font/src/LogIn/MainLogIn.tsx toggles between LogIn/SingUp via a
// Switch, but MainLogIn is never mounted by any <Route> in font/src/App.tsx —
// only <Route path="/LogIn" element={<LogIn />} /> is registered. There is
// currently no way to reach the sign-up form through the running app. This
// test pins that behavior down; update/remove it once sign-up is wired up.
test('there is no sign-up entry point reachable from the login screen', async ({ page }) => {
  await page.goto('/LogIn');
  await expect(page.getByRole('button', { name: /הרשמ/ })).toHaveCount(0);
  await expect(page.getByRole('link', { name: /הרשמ/ })).toHaveCount(0);

  await page.goto('/SingUp');
  // No matching route -> falls through to the catch-all Navigate("/LogIn") at "/"...
  // actually there is no wildcard route either, so React Router renders nothing.
  // Assert we did NOT land on a working sign-up form regardless.
  await expect(page.getByLabel('אימות סיסמה')).toHaveCount(0);
});
