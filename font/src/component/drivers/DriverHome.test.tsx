import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../../GlobalState';
import DriverHome from './DriverHome';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function renderPage() {
    localStorage.setItem('level', JSON.stringify(5));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <DriverHome />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('DriverHome', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('shows a loading state before the drives resolve', () => {
        mockedAxios.post.mockReturnValueOnce(new Promise(() => {}));
        renderPage();
        expect(screen.getByText('טוען נסיעות...')).toBeInTheDocument();
    });

    it('renders the fetched drives in a table', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: {
                message: 'work',
                'list of drives': [
                    { Name: 'קו 5', HToStart: 8, MToStart: 30, GpsStart: '32.0,34.8', GpsEnd: '32.1,34.9' },
                ],
            },
        });
        renderPage();

        expect(await screen.findByText('קו 5')).toBeInTheDocument();
        expect(screen.getByText('8:30')).toBeInTheDocument();
        expect(screen.getByText('32.0,34.8')).toBeInTheDocument();
    });

    it('shows an empty-state message when there are no drives', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { message: 'work', 'list of drives': [] } });
        renderPage();

        expect(await screen.findByText('אין נסיעות זמינות')).toBeInTheDocument();
    });

    it('does not crash and falls back to the empty state when the request fails', async () => {
        mockedAxios.post.mockRejectedValueOnce(new Error('network error'));
        renderPage();

        expect(await screen.findByText('אין נסיעות זמינות')).toBeInTheDocument();
    });

    it('clicking "ניווט" opens Google Maps directions to that GPS point', async () => {
        mockedAxios.post.mockResolvedValueOnce({
            data: {
                message: 'work',
                'list of drives': [
                    { Name: 'קו 5', HToStart: 8, MToStart: 30, GpsStart: '32.0,34.8', GpsEnd: '32.1,34.9' },
                ],
            },
        });
        const openSpy = vi.spyOn(window, 'open').mockImplementation(() => null);
        const user = userEvent.setup();
        renderPage();

        const [navigateToEnd] = await screen.findAllByRole('button', { name: 'ניווט' });
        await user.click(navigateToEnd);

        expect(openSpy).toHaveBeenCalledWith(
            'https://www.google.com/maps/dir/?api=1&destination=32.1,34.9',
            '_blank'
        );
        openSpy.mockRestore();
    });
});
