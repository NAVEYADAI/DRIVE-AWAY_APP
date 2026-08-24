import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../../GlobalState';
import AddDriver from './AddDriver';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function renderAddDriver() {
    localStorage.setItem('level', JSON.stringify(3));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <AddDriver />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('AddDriver', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('renders the user fields and the driver-specific header fields', () => {
        renderAddDriver();
        expect(screen.getByLabelText('שם פרטי')).toBeInTheDocument();
        expect(screen.getByLabelText('כתובת')).toBeInTheDocument();
        expect(screen.getByRole('checkbox', { name: 'שכיר' })).toBeInTheDocument();
    });

    it('submits the combined form data to the AddDriver endpoint', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { message: 'ok' } });
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderAddDriver();

        await user.type(screen.getByLabelText('שם פרטי'), 'ישראל');
        await user.type(screen.getByLabelText('כתובת'), 'רחוב הרצל 1');
        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(mockedAxios.post).toHaveBeenCalledWith(
            'http://localhost:2909/MainController/AddDriver',
            expect.objectContaining({ fName: 'ישראל', address: 'רחוב הרצל 1', isEmployee: false })
        );
        expect(alertSpy).toHaveBeenCalledWith('נהג נוסף בהצלחה');
        alertSpy.mockRestore();
    });

    it('shows an error alert when the request fails', async () => {
        mockedAxios.post.mockRejectedValueOnce(new Error('network error'));
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderAddDriver();

        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(alertSpy).toHaveBeenCalledWith('שגיאה בהוספת נהג');
        alertSpy.mockRestore();
    });
});
