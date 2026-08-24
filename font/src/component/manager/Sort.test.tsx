import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../../GlobalState';
import Sort from './Sort';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function mockEmptyListsAndSortEndpoint() {
    mockedAxios.post.mockImplementation((url: string) => {
        if (url.includes('/GetDrives')) {
            return Promise.resolve({ data: { message: 'Success', List: [] } });
        }
        if (url.includes('/GetDrivers')) {
            return Promise.resolve({ data: { message: 'Success', List: [] } });
        }
        if (url.includes('/Sort')) {
            return Promise.resolve({ data: { message: 'ok' } });
        }
        return Promise.reject(new Error(`unexpected url ${url}`));
    });
}

function renderPage() {
    localStorage.setItem('level', JSON.stringify(3));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <Sort />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('Sort', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('renders the sort-name field, the embedded EditDrives and EditDrivers lists', async () => {
        mockEmptyListsAndSortEndpoint();
        renderPage();

        expect(screen.getByLabelText('שם השיבוץ')).toBeInTheDocument();
        expect(await screen.findByText('אין נסיעות זמינות')).toBeInTheDocument();
        expect(await screen.findByText('אין נהגים זמינים')).toBeInTheDocument();
    });

    it('submits the sort name to the /Sort endpoint and alerts success', async () => {
        mockEmptyListsAndSortEndpoint();
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderPage();

        await user.type(screen.getByLabelText('שם השיבוץ'), 'שיבוץ בוקר');
        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(mockedAxios.post).toHaveBeenCalledWith(
            'http://localhost:2909/Sort',
            expect.objectContaining({ nameOfSort: 'שיבוץ בוקר' })
        );
        expect(alertSpy).toHaveBeenCalledWith('השיבוץ נוצר בהצלחה');
        alertSpy.mockRestore();
    });

    it('alerts failure when creating the sort fails', async () => {
        mockedAxios.post.mockImplementation((url: string) => {
            if (url.includes('/GetDrives') || url.includes('/GetDrivers')) {
                return Promise.resolve({ data: { message: 'Success', List: [] } });
            }
            return Promise.reject(new Error('network error'));
        });
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderPage();

        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(alertSpy).toHaveBeenCalledWith('שגיאה ביצירת השיבוץ');
        alertSpy.mockRestore();
    });
});
