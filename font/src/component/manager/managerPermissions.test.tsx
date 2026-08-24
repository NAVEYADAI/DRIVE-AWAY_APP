import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import { managerPermissions } from './managerPermissions';

function Protected() {
    return <div>protected content</div>;
}
const Wrapped = managerPermissions(Protected);

function renderAtLevel(level: number) {
    localStorage.setItem('level', JSON.stringify(level));
    return render(
        <GlobalStateProvider>
            <MemoryRouter initialEntries={['/manager']}>
                <Routes>
                    <Route path="/manager" element={<Wrapped />} />
                    <Route path="/LogIn" element={<div>login screen</div>} />
                </Routes>
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('managerPermissions', () => {
    beforeEach(() => {
        localStorage.clear();
    });

    it.each([1, 2, 3, 4])('allows level %i through', (level) => {
        renderAtLevel(level);
        expect(screen.getByText('protected content')).toBeInTheDocument();
    });

    it('redirects level 0 (logged out) to /LogIn', () => {
        renderAtLevel(0);
        expect(screen.getByText('login screen')).toBeInTheDocument();
    });

    it('redirects level 5 (driver) to /LogIn', () => {
        renderAtLevel(5);
        expect(screen.getByText('login screen')).toBeInTheDocument();
    });
});
