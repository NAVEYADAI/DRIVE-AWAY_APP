import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import axios from 'axios';
import { GlobalStateProvider } from '../../GlobalState';
import CompanyRegistration from './CompanyRegistration';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

function renderPage() {
    localStorage.setItem('level', JSON.stringify(1));
    return render(
        <GlobalStateProvider>
            <MemoryRouter>
                <CompanyRegistration />
            </MemoryRouter>
        </GlobalStateProvider>
    );
}

async function fillAllRequiredFields(user: ReturnType<typeof userEvent.setup>) {
    await user.type(screen.getByLabelText('שם החברה'), 'חברת בדיקה');
    await user.type(screen.getByLabelText('ח.פ. / ת״ז'), '123456789');
    await user.type(screen.getByLabelText('כתובת'), 'רחוב הרצל 1');
    await user.type(screen.getByLabelText('שם פרטי'), 'ישראל');
    await user.type(screen.getByLabelText('שם משפחה'), 'ישראלי');
    await user.type(screen.getByLabelText('תעודת זהות'), '987654321');
    await user.type(screen.getByLabelText('שם משתמש'), 'israel');
    await user.type(screen.getByLabelText('סיסמה'), 'Secret123!');
    await user.type(screen.getByLabelText('כתובת מייל'), 'x@example.com');
    await user.type(screen.getByLabelText('מספר טלפון'), '0501234567');
}

describe('CompanyRegistration', () => {
    beforeEach(() => {
        localStorage.clear();
        vi.clearAllMocks();
    });

    it('renders both the company and the admin-user sub-forms', () => {
        renderPage();
        expect(screen.getByLabelText('שם החברה')).toBeInTheDocument();
        expect(screen.getByLabelText('שם פרטי')).toBeInTheDocument();
    });

    // Regression test: AddCompany's "שם החברה" field and AddUser's "שם פרטי"
    // field both used id="FirstName", so on this page (which renders both
    // sub-forms together) their <label for="FirstName"> both resolved to the
    // same underlying input — typing into one silently typed into the other.
    // Fixed by renaming AddCompany's id to "CompanyName".
    it('the company-name and first-name fields are independent inputs', async () => {
        const user = userEvent.setup();
        renderPage();

        await user.type(screen.getByLabelText('שם החברה'), 'חברת בדיקה');
        await user.type(screen.getByLabelText('שם פרטי'), 'ישראל');

        expect(screen.getByLabelText('שם החברה')).toHaveValue('חברת בדיקה');
        expect(screen.getByLabelText('שם פרטי')).toHaveValue('ישראל');
    });

    it('does not submit when required fields are missing', async () => {
        const user = userEvent.setup();
        renderPage();

        await user.type(screen.getByLabelText('שם החברה'), 'חברת בדיקה');
        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(mockedAxios.post).not.toHaveBeenCalled();
    });

    // Longer timeout: this fills 10 fields sequentially via userEvent, which
    // can exceed the default 5s budget under load (e.g. the full suite
    // running many jsdom environments in parallel).
    it('submits combined company + admin-user data once every required field is filled', async () => {
        mockedAxios.post.mockResolvedValueOnce({ data: { message: 'ok' } });
        const alertSpy = vi.spyOn(window, 'alert').mockImplementation(() => {});
        const user = userEvent.setup();
        renderPage();

        await fillAllRequiredFields(user);
        await user.click(screen.getByRole('button', { name: 'שליחה' }));

        expect(mockedAxios.post).toHaveBeenCalledWith(
            'http://localhost:2909/MainController/AddCompany',
            expect.objectContaining({ companyName: 'חברת בדיקה', fName: 'ישראל' })
        );
        expect(alertSpy).toHaveBeenCalledWith('החברה נוספה');
        alertSpy.mockRestore();
    }, 15000);
});
