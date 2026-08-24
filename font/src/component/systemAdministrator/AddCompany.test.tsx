import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { useState } from 'react';
import { AddCompany } from './AddCompany';
import type { CompanyFormData } from '../../types';

function Harness() {
    const [data, setData] = useState<CompanyFormData>({ CompanyNameValue: '', CompanyIDValue: '', AddressValue: '' });
    return <AddCompany companyData={data} setCompanyData={setData} />;
}

describe('AddCompany', () => {
    it('renders every field', () => {
        render(<Harness />);
        expect(screen.getByLabelText('שם החברה')).toBeInTheDocument();
        expect(screen.getByLabelText('ח.פ. / ת״ז')).toBeInTheDocument();
        expect(screen.getByLabelText('כתובת')).toBeInTheDocument();
    });

    it('typing updates each field', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        await user.type(screen.getByLabelText('שם החברה'), 'חברת בדיקה');
        await user.type(screen.getByLabelText('ח.פ. / ת״ז'), '123456789');
        await user.type(screen.getByLabelText('כתובת'), 'רחוב הרצל 1');

        expect(screen.getByLabelText('שם החברה')).toHaveValue('חברת בדיקה');
        expect(screen.getByLabelText('ח.פ. / ת״ז')).toHaveValue('123456789');
        expect(screen.getByLabelText('כתובת')).toHaveValue('רחוב הרצל 1');
    });
});
