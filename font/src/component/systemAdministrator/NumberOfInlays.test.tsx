import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import NumberOfInlays from './NumberOfInlays';

function renderPage() {
    localStorage.setItem('level', JSON.stringify(1));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <NumberOfInlays />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

describe('NumberOfInlays', () => {
    it('renders the report heading and the list of companies with their sort counts', () => {
        renderPage();
        expect(screen.getByText('דוח מספר שיבוצים')).toBeInTheDocument();
        expect(screen.getByText('name')).toBeInTheDocument();
        expect(screen.getByText('4')).toBeInTheDocument();
        expect(screen.getByText('yadai')).toBeInTheDocument();
        expect(screen.getByText('3')).toBeInTheDocument();
    });

    // Known gap, not fixed here (missing feature, not a regression):
    // GetValueOnCompany() returns a hardcoded fake list instead of fetching
    // real sort counts from the API.
    it.skip('known gap: company sort counts are hardcoded, not fetched from the API', () => {});
});
