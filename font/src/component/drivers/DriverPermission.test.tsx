import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import { driverPermissions } from './DriverPermission';

function Protected() {
    return <div>protected content</div>;
}
const Wrapped = driverPermissions(Protected);

function renderAtLevel(level: number) {
    localStorage.setItem('level', JSON.stringify(level));
    return render(
        <GlobalStateProvider>
            <MemoryRouter initialEntries={['/driver']}>
                <Routes>
                    <Route path="/driver" element={<Wrapped />} />
                    <Route path="/LogIn" element={<div>login screen</div>} />
                </Routes>
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('driverPermissions', () => {
    beforeEach(() => {
        localStorage.clear();
    });

    it.each([1, 2, 3, 4, 5])('allows any logged-in level %i through', (level) => {
        renderAtLevel(level);
        expect(screen.getByText('protected content')).toBeInTheDocument();
    });

    it('redirects level 0 (logged out) to /LogIn', () => {
        renderAtLevel(0);
        expect(screen.getByText('login screen')).toBeInTheDocument();
    });
});
