import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import { GlobalStateProvider } from '../../GlobalState';
import EditSort from './EditSort';

describe('EditSort', () => {
    it('renders the "under construction" placeholder', () => {
        localStorage.setItem('level', JSON.stringify(3));
        render(
            <GlobalStateProvider>
                <MemoryRouter>
                    <EditSort />
                </MemoryRouter>
            </GlobalStateProvider>
        );

        expect(screen.getByText('עריכת שיבוצים')).toBeInTheDocument();
        expect(screen.getByText('מסך זה נמצא בבנייה')).toBeInTheDocument();
    });
});
