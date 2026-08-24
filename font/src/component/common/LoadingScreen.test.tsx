import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import { LoadingScreen } from './LoadingScreen';

describe('LoadingScreen', () => {
    it('renders the default label', () => {
        render(<LoadingScreen />);
        expect(screen.getByText('טוען...')).toBeInTheDocument();
    });

    it('renders a custom label when provided', () => {
        render(<LoadingScreen label="טוען נהגים..." />);
        expect(screen.getByText('טוען נהגים...')).toBeInTheDocument();
        expect(screen.queryByText('טוען...')).not.toBeInTheDocument();
    });
});
