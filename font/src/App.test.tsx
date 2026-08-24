import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen, within } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import App from './App';

// The page title now renders both in the AppBar and as the page's own
// heading (both correctly kept in sync via setNameOfPage), so text queries
// must be scoped to the main content area to avoid ambiguous matches.
function main() {
    return within(screen.getByRole('main'));
}

// Regression coverage for a real bug: the hamburger menu items (דף הבית /
// שיבוץ נסיעה / התנתקות) used to only close the menu and never actually
// navigate anywhere. Fixed by wiring real useNavigate() handlers in App.tsx.
describe('App hamburger menu', () => {
    beforeEach(() => {
        localStorage.clear();
        window.history.pushState({}, '', '/manager');
    });

    async function openMenu(user: ReturnType<typeof userEvent.setup>) {
        await user.click(screen.getByRole('button', { name: 'menu' }));
    }

    it('"דף הבית" navigates a manager-level user to /manager', async () => {
        localStorage.setItem('level', JSON.stringify(3));
        const user = userEvent.setup();
        render(<App />);

        await openMenu(user);
        await user.click(screen.getByRole('menuitem', { name: 'דף הבית' }));

        expect(await main().findByText('ניהול חברה')).toBeInTheDocument();
    });

    it('"התנתקות" logs the user out and returns to the login screen', async () => {
        localStorage.setItem('level', JSON.stringify(3));
        const user = userEvent.setup();
        render(<App />);

        await openMenu(user);
        await user.click(screen.getByRole('menuitem', { name: 'התנתקות' }));

        expect(await screen.findByText('כניסה למערכת')).toBeInTheDocument();
        expect(JSON.parse(localStorage.getItem('level') ?? '-1')).toBe(0);
    });

    it('"שיבוץ נסיעה" navigates straight to the sort-creation screen', async () => {
        localStorage.setItem('level', JSON.stringify(3));
        const user = userEvent.setup();
        render(<App />);

        await openMenu(user);
        await user.click(screen.getByRole('menuitem', { name: 'שיבוץ נסיעה' }));

        expect(await main().findByText('יצירת שיבוץ')).toBeInTheDocument();
    });
});
