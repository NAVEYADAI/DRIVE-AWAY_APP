import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../GlobalState';
import { LogIn } from './LogIn';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function renderLogIn() {
    return render(
        <GlobalStateProvider>
            <MemoryRouter initialEntries={['/LogIn']}>
                <Routes>
                    <Route path="/LogIn" element={<LogIn />} />
                    <Route path="/systemAdministrator" element={<div>system administrator home</div>} />
                    <Route path="/manager" element={<div>manager home</div>} />
                    <Route path="/driver" element={<div>driver home</div>} />
                </Routes>
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

async function login(user: ReturnType<typeof userEvent.setup>) {
    await user.type(screen.getByLabelText('שם משתמש'), 'someuser');
    await user.type(screen.getByLabelText('סיסמה'), 'somepass');
    await user.click(screen.getByRole('button', { name: 'כניסה' }));
}

describe('LogIn', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('posts credentials to the login endpoint', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { messege: 'ok', UserId: 1, level: 3 } });
        const user = userEvent.setup();
        renderLogIn();

        await login(user);

        expect(mockedAxios.post).toHaveBeenCalledWith(
            'http://localhost:2909/login',
            expect.objectContaining({ UserName: 'someuser', Password: 'somepass' })
        );
    });

    it('navigates admin levels (1-2) to /systemAdministrator', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { messege: 'ok', UserId: 1, level: 1 } });
        const user = userEvent.setup();
        renderLogIn();
        await login(user);
        expect(await screen.findByText('system administrator home')).toBeInTheDocument();
    });

    it('navigates manager levels (3-4) to /manager', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { messege: 'ok', UserId: 1, level: 3 } });
        const user = userEvent.setup();
        renderLogIn();
        await login(user);
        expect(await screen.findByText('manager home')).toBeInTheDocument();
    });

    // Regression test: level 5 (driver) previously had no navigation branch
    // at all, so a driver who logged in successfully was silently stuck on
    // the login screen. Fixed by adding the `level === 5` branch in LogIn.tsx.
    it('navigates driver level (5) to /driver', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { messege: 'ok', UserId: 1, level: 5 } });
        const user = userEvent.setup();
        renderLogIn();
        await login(user);
        expect(await screen.findByText('driver home')).toBeInTheDocument();
    });

    it('shows an error alert and stays on the login screen when the request fails', async () => {
        mockedAxios.post.mockRejectedValueOnce(new Error('network error'));
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderLogIn();

        await login(user);

        expect(alertSpy).toHaveBeenCalledWith('שגיאת התחברות');
        expect(screen.getByRole('button', { name: 'כניסה' })).toBeInTheDocument();
        alertSpy.mockRestore();
    });

    it('does not submit when a field is empty', async () => {
        const user = userEvent.setup();
        renderLogIn();

        await user.type(screen.getByLabelText('שם משתמש'), 'someuser');
        await user.click(screen.getByRole('button', { name: 'כניסה' }));

        expect(mockedAxios.post).not.toHaveBeenCalled();
    });
});
