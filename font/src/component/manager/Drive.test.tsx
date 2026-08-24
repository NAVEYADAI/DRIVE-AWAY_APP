import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { useState } from 'react';
import { Drive } from './Drive';
import type { DriveFormData } from '../../types';

function Harness() {
    const [data, setData] = useState<DriveFormData>({
        NameValue: '', SelectTimeValue: null, StartGpsValue: '', EndGpsValue: '',
    });
    return <Drive DriveValue={data} setDriveValue={setData} />;
}

describe('Drive', () => {
    it('renders the line name, time picker and GPS fields', () => {
        render(<Harness />);
        expect(screen.getByLabelText('שם הקו')).toBeInTheDocument();
        expect(screen.getByLabelText('שעת יציאה')).toBeInTheDocument();
        expect(screen.getByLabelText('מיקום התחלה')).toBeInTheDocument();
        expect(screen.getByLabelText('מיקום סיום')).toBeInTheDocument();
    });

    it('typing updates the line name and GPS fields', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        await user.type(screen.getByLabelText('שם הקו'), 'קו 5');
        await user.type(screen.getByLabelText('מיקום התחלה'), '32.0,34.8');
        await user.type(screen.getByLabelText('מיקום סיום'), '32.1,34.9');

        expect(screen.getByLabelText('שם הקו')).toHaveValue('קו 5');
        expect(screen.getByLabelText('מיקום התחלה')).toHaveValue('32.0,34.8');
        expect(screen.getByLabelText('מיקום סיום')).toHaveValue('32.1,34.9');
    });
});
