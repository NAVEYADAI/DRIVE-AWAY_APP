import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../../GlobalState';
import EditDrivers from './EditDrivers';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function renderPage() {
    localStorage.setItem('level', JSON.stringify(3));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <EditDrivers />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('EditDrivers', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('shows a loading state before the drivers list resolves', () => {
        mockedAxios.post.mockReturnValueOnce(new Promise(() => {})); // never resolves
        renderPage();
        expect(screen.getByText('טוען נהגים...')).toBeInTheDocument();
    });

    it('renders the fetched drivers once loaded', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: { message: 'Success', List: [{ nameOfDriver: 'דני', aviable: true }, { nameOfDriver: 'משה', aviable: false }] },
        });
        renderPage();

        expect(await screen.findByText('דני')).toBeInTheDocument();
        expect(screen.getByText('משה')).toBeInTheDocument();
        expect(mockedAxios.post).toHaveBeenCalledWith(
            'http://localhost:2909/MainController/GetDrivers',
            expect.objectContaining({ userId: -1, level: 3 })
        );
    });

    it('shows an empty-state message when there are no drivers', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { message: 'Success', List: [] } });
        renderPage();

        expect(await screen.findByText('אין נהגים זמינים')).toBeInTheDocument();
    });

    it('does not crash and falls back to the empty state when the request fails', async () => {
        mockedAxios.post.mockRejectedValueOnce(new Error('network error'));
        renderPage();

        expect(await screen.findByText('אין נהגים זמינים')).toBeInTheDocument();
    });

    it('toggling a driver switch calls ViewOrNotDriver with that driver', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: { message: 'Success', List: [{ nameOfDriver: 'דני', aviable: true }] },
        });
        mockedAxios.post.mockResolvedValueOnce({ data: {} });
        const user = userEvent.setup();
        renderPage();

        const toggle = await screen.findByRole('checkbox');
        expect(toggle).toBeChecked();
        await user.click(toggle);

        expect(mockedAxios.post).toHaveBeenLastCalledWith(
            'http://localhost:2909/MainController/ViewOrNotDriver',
            expect.objectContaining({ driverName: 'דני' })
        );
    });

    it('clicking delete calls DeleteDriver with that driver', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: { message: 'Success', List: [{ nameOfDriver: 'דני', aviable: true }] },
        });
        mockedAxios.post.mockResolvedValueOnce({ data: {} });
        const user = userEvent.setup();
        renderPage();

        await screen.findByText('דני');
        await user.click(screen.getByTestId('DeleteIcon'));

        expect(mockedAxios.post).toHaveBeenLastCalledWith(
            'http://localhost:2909/MainController/DeleteDriver',
            expect.objectContaining({ driverName: 'דני' })
        );
    });
});
