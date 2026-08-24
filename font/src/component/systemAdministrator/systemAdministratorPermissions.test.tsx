import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import { systemAdministratorPermissions } from './systemAdministratorPermissions';

function Protected() {
    return <div>protected content</div>;
}
const Wrapped = systemAdministratorPermissions(Protected);

function renderAtLevel(level: number) {
    localStorage.setItem('level', JSON.stringify(level));
    return render(
        <GlobalStateProvider>
            <MemoryRouter initialEntries={['/systemAdministrator']}>
                <Routes>
                    <Route path="/systemAdministrator" element={<Wrapped />} />
                    <Route path="/LogIn" element={<div>login screen</div>} />
                </Routes>
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('systemAdministratorPermissions', () => {
    beforeEach(() => {
        localStorage.clear();
    });

    it.each([1, 2])('allows level %i through', (level) => {
        renderAtLevel(level);
        expect(screen.getByText('protected content')).toBeInTheDocument();
    });

    it.each([0, 3, 4, 5])('redirects level %i to /LogIn (only admin levels 1-2 are allowed)', (level) => {
        renderAtLevel(level);
        expect(screen.getByText('login screen')).toBeInTheDocument();
    });
});
