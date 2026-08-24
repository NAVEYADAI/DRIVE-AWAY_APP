import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import SystemAdministrator from './mainHome';

function renderAt(path: string) {
    localStorage.setItem('level', JSON.stringify(1));
    return render(
        <GlobalStateProvider>
            <MemoryRouter initialEntries={['/systemAdministrator']}>
                <Routes>
                    <Route path="/systemAdministrator" element={<SystemAdministrator />} />
                    <Route path={path} element={<div>landed on {path}</div>} />
                </Routes>
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('SystemAdministrator (mainHome)', () => {
    beforeEach(() => {
        localStorage.clear();
    });

    it('renders every action card', () => {
        renderAt('/manager');
        expect(screen.getByText('רישום חברה')).toBeInTheDocument();
        expect(screen.getByText('רישום עמדה בחברה קיימת')).toBeInTheDocument();
        expect(screen.getByText('דוח מספר שיבוצים')).toBeInTheDocument();
        expect(screen.getByText('חברה שלי')).toBeInTheDocument();
    });

    it.each([
        ['רישום חברה', '/systemAdministrator/CompanyRegistration'],
        ['רישום עמדה בחברה קיימת', '/systemAdministrator/AddPosition'],
        ['דוח מספר שיבוצים', '/systemAdministrator/NumberOfInlays'],
        ['חברה שלי', '/manager'],
    ])('clicking "%s" navigates to %s', async (label, path) => {
        const user = userEvent.setup();
        renderAt(path);

        await user.click(screen.getByText(label));

        expect(await screen.findByText(`landed on ${path}`)).toBeInTheDocument();
    });
});
