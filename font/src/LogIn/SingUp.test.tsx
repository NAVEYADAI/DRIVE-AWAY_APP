import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { SingUp } from './SingUp';

describe('SingUp', () => {
    it('renders every registration field', () => {
        render(<SingUp />);
        expect(screen.getByLabelText('שם משתמש')).toBeInTheDocument();
        expect(screen.getByLabelText('סיסמה')).toBeInTheDocument();
        expect(screen.getByLabelText('אימות סיסמה')).toBeInTheDocument();
        expect(screen.getByLabelText('שם פרטי')).toBeInTheDocument();
        expect(screen.getByLabelText('שם משפחה')).toBeInTheDocument();
        expect(screen.getByLabelText('טלפון')).toBeInTheDocument();
        expect(screen.getByLabelText('דוא״ל')).toBeInTheDocument();
        expect(screen.getByRole('button', { name: 'הרשמה' })).toBeInTheDocument();
    });

    it('updates each field as the user types', async () => {
        const user = userEvent.setup();
        render(<SingUp />);

        await user.type(screen.getByLabelText('שם משתמש'), 'newuser');
        await user.type(screen.getByLabelText('דוא״ל'), 'x@example.com');

        expect(screen.getByLabelText('שם משתמש')).toHaveValue('newuser');
        expect(screen.getByLabelText('דוא״ל')).toHaveValue('x@example.com');
    });

    // Known gap, not fixed here (out of scope for a testing pass — this is a
    // missing feature, not a regression): the "הרשמה" button has no onClick
    // handler at all, so filling out this form and submitting currently does
    // nothing. Flagging via a skipped test rather than silently ignoring it.
    it.skip('known gap: submitting the form does not send any request (no onClick handler exists)', () => {
        // Intentionally left unimplemented — see comment above.
    });
});
