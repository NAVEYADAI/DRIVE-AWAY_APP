import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter, Routes, Route } from 'react-router-dom';
import { NotFound } from './NotFound';

function renderNotFound() {
    return render(
        <MemoryRouter initialEntries={['/does-not-exist']}>
            <Routes>
                <Route path="/does-not-exist" element={<NotFound />} />
                <Route path="/LogIn" element={<div>login screen</div>} />
            </Routes>
        </MemoryRouter>
    );
}

describe('NotFound', () => {
    it('renders the not-found message', () => {
        renderNotFound();
        expect(screen.getByText('הדף לא נמצא')).toBeInTheDocument();
    });

    it('navigates back to /LogIn when the button is clicked', async () => {
        const user = userEvent.setup();
        renderNotFound();

        await user.click(screen.getByRole('button', { name: 'חזרה לדף הבית' }));

        expect(await screen.findByText('login screen')).toBeInTheDocument();
    });
});
