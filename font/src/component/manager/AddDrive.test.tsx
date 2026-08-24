import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../../GlobalState';
import AddDrive from './AddDrive';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function renderAddDrive() {
    localStorage.setItem('level', JSON.stringify(3));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <AddDrive />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

async function pickAStartTime(user: ReturnType<typeof userEvent.setup>) {
    await user.click(screen.getByLabelText('שעת יציאה'));
    await user.keyboard('0830AM');
}

describe('AddDrive', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('renders the drive fields', () => {
        renderAddDrive();
        expect(screen.getByLabelText('שם הקו')).toBeInTheDocument();
        expect(screen.getByRole('button', { name: 'שליחה' })).toBeInTheDocument();
    });

    // Regression coverage for validation added concurrently with this test
    // pass: submitting without picking a start time used to send
    // dayjs(null).hour()/.minute() (NaN) straight to the backend. It now
    // blocks submission with a dedicated alert instead.
    it('blocks submission and alerts when no start time has been picked', async () => {
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderAddDrive();

        await user.type(screen.getByLabelText('שם הקו'), 'קו 5');
        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(alertSpy).toHaveBeenCalledWith('יש לבחור שעת יציאה');
        expect(mockedAxios.post).not.toHaveBeenCalled();
        alertSpy.mockRestore();
    });

    it('submits the drive data to the AddDrive endpoint once a start time is picked', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { message: 'ok' } });
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderAddDrive();

        await user.type(screen.getByLabelText('שם הקו'), 'קו 5');
        await pickAStartTime(user);
        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(mockedAxios.post).toHaveBeenCalledWith(
            'http://localhost:2909/MainController/AddDrive',
            expect.objectContaining({ nameOfDriver: 'קו 5', timeToStartH: 8, timeToStartM: 30 })
        );
        expect(alertSpy).toHaveBeenCalledWith('הנסיעה בוצעה בהצלחה');
        alertSpy.mockRestore();
    });

    it('shows an error alert when the request fails', async () => {
        mockedAxios.post.mockRejectedValueOnce(new Error('network error'));
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderAddDrive();

        await pickAStartTime(user);
        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(alertSpy).toHaveBeenCalledWith('שגיאה בהוספת נסיעה');
        alertSpy.mockRestore();
    });
});
