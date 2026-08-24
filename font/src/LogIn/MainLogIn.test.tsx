import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../GlobalState';
import { MainLogIn } from './MainLogIn';

vi.mock('axios');

function renderMainLogIn() {
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <MainLogIn />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('MainLogIn', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('shows the sign-up form by default (switch defaults to checked)', () => {
        renderMainLogIn();
        expect(screen.getByText('הרשמה למערכת')).toBeInTheDocument();
    });

    it('toggling the switch shows the login form instead', async () => {
        const user = userEvent.setup();
        renderMainLogIn();

        await user.click(screen.getByRole('checkbox'));

        expect(screen.getByText('כניסה למערכת')).toBeInTheDocument();
        expect(screen.queryByText('הרשמה למערכת')).not.toBeInTheDocument();
    });
});
