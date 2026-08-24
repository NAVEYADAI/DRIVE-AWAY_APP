import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import ManagerHome from './managerHome';

function renderAt(path: string) {
    localStorage.setItem('level', JSON.stringify(3));
    return render(
        <GlobalStateProvider>
            <MemoryRouter initialEntries={['/manager']}>
                <Routes>
                    <Route path="/manager" element={<ManagerHome />} />
                    <Route path={path} element={<div>landed on {path}</div>} />
                </Routes>
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('ManagerHome', () => {
    beforeEach(() => {
        localStorage.clear();
    });

    it('renders every action card', () => {
        renderAt('/manager/AddDriver');
        expect(screen.getByText('הוספת נהג')).toBeInTheDocument();
        expect(screen.getByText('עריכת נהגים')).toBeInTheDocument();
        expect(screen.getByText('הוספת נסיעה')).toBeInTheDocument();
        expect(screen.getByText('עריכת נסיעות')).toBeInTheDocument();
        expect(screen.getByText('יצירת שיבוץ')).toBeInTheDocument();
        expect(screen.getByText('עריכת שיבוצים')).toBeInTheDocument();
        expect(screen.getByText('הנסיעות שלי')).toBeInTheDocument();
    });

    it.each([
        ['הוספת נהג', '/manager/AddDriver'],
        ['עריכת נהגים', '/manager/EditDrivers'],
        ['הוספת נסיעה', '/manager/AddDrive'],
        ['עריכת נסיעות', '/manager/EditDrives'],
        ['יצירת שיבוץ', '/manager/Sort'],
        ['עריכת שיבוצים', '/manager/EditSort'],
        ['הנסיעות שלי', '/driver'],
    ])('clicking "%s" navigates to %s', async (label, path) => {
        const user = userEvent.setup();
        renderAt(path);

        await user.click(screen.getByText(label));

        expect(await screen.findByText(`landed on ${path}`)).toBeInTheDocument();
    });
});
