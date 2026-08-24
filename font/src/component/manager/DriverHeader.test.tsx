import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { useState } from 'react';
import { DriverHeader } from './DriverHeader';
import type { DriverHeaderData } from '../../types';

function Harness() {
    const [data, setData] = useState<DriverHeaderData>({ AddressValue: '', IsEmployeeValue: false });
    return <DriverHeader DriverHeaderValue={data} setDriverHeaderValue={setData} />;
}

describe('DriverHeader', () => {
    it('renders the address field and the "שכיר" switch', () => {
        render(<Harness />);
        expect(screen.getByLabelText('כתובת')).toBeInTheDocument();
        expect(screen.getByRole('checkbox', { name: 'שכיר' })).not.toBeChecked();
    });

    it('typing updates the address field', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        await user.type(screen.getByLabelText('כתובת'), 'רחוב הרצל 1');

        expect(screen.getByLabelText('כתובת')).toHaveValue('רחוב הרצל 1');
    });

    // Regression coverage: the switch used to be wired via `value` instead of
    // `checked`, so it never actually reflected/updated the real toggle
    // state. Confirmed already fixed — this pins the correct behavior.
    it('toggling the switch flips IsEmployeeValue and stays in sync with clicks', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        const toggle = screen.getByRole('checkbox', { name: 'שכיר' });
        await user.click(toggle);
        expect(toggle).toBeChecked();

        await user.click(toggle);
        expect(toggle).not.toBeChecked();
    });
});
