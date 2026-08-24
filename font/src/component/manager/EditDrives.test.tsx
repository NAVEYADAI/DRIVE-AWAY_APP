import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../../GlobalState';
import EditDrives from './EditDrives';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function renderPage() {
    localStorage.setItem('level', JSON.stringify(3));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <EditDrives />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('EditDrives', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('shows a loading state before the drives list resolves', () => {
        mockedAxios.post.mockReturnValueOnce(new Promise(() => {}));
        renderPage();
        expect(screen.getByText('טוען נסיעות...')).toBeInTheDocument();
    });

    it('renders the fetched drives once loaded', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: { message: 'Success', List: [{ nameOfDrive: 'קו 5', aviable: true }] },
        });
        renderPage();

        expect(await screen.findByText('קו 5')).toBeInTheDocument();
    });

    it('shows an empty-state message when there are no drives', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { message: 'Success', List: [] } });
        renderPage();

        expect(await screen.findByText('אין נסיעות זמינות')).toBeInTheDocument();
    });

    it('does not crash and falls back to the empty state when the request fails', async () => {
        mockedAxios.post.mockRejectedValueOnce(new Error('network error'));
        renderPage();

        expect(await screen.findByText('אין נסיעות זמינות')).toBeInTheDocument();
    });

    it('toggling a drive switch calls ViewOrNotDrive with that drive', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: { message: 'Success', List: [{ nameOfDrive: 'קו 5', aviable: false }] },
        });
        mockedAxios.post.mockResolvedValueOnce({ data: {} });
        const user = userEvent.setup();
        renderPage();

        const toggle = await screen.findByRole('checkbox');
        expect(toggle).not.toBeChecked();
        await user.click(toggle);

        expect(mockedAxios.post).toHaveBeenLastCalledWith(
            'http://localhost:2909/MainController/ViewOrNotDrive',
            expect.objectContaining({ driveName: 'קו 5' })
        );
    });

    it('deleting a drive alerts success on a resolved request', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: { message: 'Success', List: [{ nameOfDrive: 'קו 5', aviable: true }] },
        });
        mockedAxios.post.mockResolvedValueOnce({ data: { message: 'ok' } });
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderPage();

        await screen.findByText('קו 5');
        await user.click(screen.getByTestId('DeleteIcon'));

        expect(mockedAxios.post).toHaveBeenLastCalledWith(
            'http://localhost:2909/MainController/DeleteDrive',
            expect.objectContaining({ driveName: 'קו 5' })
        );
        expect(alertSpy).toHaveBeenCalledWith('הנסיעה נמחקה בהצלחה');
        alertSpy.mockRestore();
    });

    it('deleting a drive alerts failure when the request rejects', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: { message: 'Success', List: [{ nameOfDrive: 'קו 5', aviable: true }] },
        });
        mockedAxios.post.mockRejectedValueOnce(new Error('network error'));
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderPage();

        await screen.findByText('קו 5');
        await user.click(screen.getByTestId('DeleteIcon'));

        expect(alertSpy).toHaveBeenCalledWith('שגיאה במחיקת נסיעה');
        alertSpy.mockRestore();
    });
});
