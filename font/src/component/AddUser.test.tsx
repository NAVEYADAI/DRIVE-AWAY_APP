import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { useState } from 'react';
import { AddUser } from './AddUser';
import type { UserFormData } from '../types';

const emptyUser: UserFormData = {
    FNameValue: '', LNameValue: '', IDValue: '', UserNameValue: '',
    PasswordValue: '', Password2Value: '', EmailValue: '', PhoneNumberValue: '',
};

function Harness() {
    const [userData, setUserData] = useState<UserFormData>(emptyUser);
    return <AddUser userData={userData} setUserData={setUserData} />;
}

describe('AddUser', () => {
    it('renders every field', () => {
        render(<Harness />);
        expect(screen.getByLabelText('שם פרטי')).toBeInTheDocument();
        expect(screen.getByLabelText('שם משפחה')).toBeInTheDocument();
        expect(screen.getByLabelText('תעודת זהות')).toBeInTheDocument();
        expect(screen.getByLabelText('שם משתמש')).toBeInTheDocument();
        expect(screen.getByLabelText('סיסמה')).toBeInTheDocument();
        expect(screen.getByLabelText('אימות סיסמה')).toBeInTheDocument();
        expect(screen.getByLabelText('כתובת מייל')).toBeInTheDocument();
        expect(screen.getByLabelText('מספר טלפון')).toBeInTheDocument();
    });

    it('updates the shared userData object as fields are typed into', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        await user.type(screen.getByLabelText('שם פרטי'), 'ישראל');
        await user.type(screen.getByLabelText('תעודת זהות'), '123456789');

        expect(screen.getByLabelText('שם פרטי')).toHaveValue('ישראל');
        expect(screen.getByLabelText('תעודת זהות')).toHaveValue('123456789');
    });

    it('password fields are masked by default and toggle to visible text', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        const passwordInput = screen.getByLabelText('סיסמה');
        expect(passwordInput).toHaveAttribute('type', 'password');

        const [toggle] = screen.getAllByLabelText('toggle password visibility');
        await user.click(toggle);

        expect(passwordInput).toHaveAttribute('type', 'text');
    });

    it('confirm-password field toggles independently from the password field', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        const confirmInput = screen.getByLabelText('אימות סיסמה');
        expect(confirmInput).toHaveAttribute('type', 'password');

        const [, confirmToggle] = screen.getAllByLabelText('toggle password visibility');
        await user.click(confirmToggle);

        expect(confirmInput).toHaveAttribute('type', 'text');
        expect(screen.getByLabelText('סיסמה')).toHaveAttribute('type', 'password');
    });
});
