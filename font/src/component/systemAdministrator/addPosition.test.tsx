import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import AddPosition from './addPosition';

function renderPage() {
    localStorage.setItem('level', JSON.stringify(1));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <AddPosition />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('AddPosition', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('renders the company select and the user sub-form', () => {
        renderPage();
        expect(screen.getByLabelText('שם החברה')).toBeInTheDocument();
        expect(screen.getByLabelText('שם פרטי')).toBeInTheDocument();
    });

    // Regression test: the InputLabel's id ("CampnyNam") didn't match the
    // Select's labelId ("CampnyName") due to a typo, breaking the
    // aria-labelledby association between the visible label and the select.
    it('the company select is correctly associated with its visible label', () => {
        renderPage();
        const select = screen.getByLabelText('שם החברה');
        expect(select).toBeInTheDocument();
        expect(select).toHaveAttribute('aria-labelledby', expect.stringContaining('CampnyName'));
    });

    it('lists the available companies as options', async () => {
        const user = userEvent.setup();
        renderPage();

        await user.click(screen.getByLabelText('שם החברה'));

        expect(await screen.findByRole('option', { name: 'nave' })).toBeInTheDocument();
        expect(screen.getByRole('option', { name: 'yadai' })).toBeInTheDocument();
    });

    // Known gaps, not fixed here (missing features, not regressions):
    // GetAllCompany() returns a hardcoded fake list instead of calling the
    // real API, and the "שליחה" button has no onClick handler at all.
    it.skip('known gap: company list is hardcoded, not fetched from the API', () => {});
    it.skip('known gap: submitting the form does not send any request (no onClick handler exists)', () => {});
});
