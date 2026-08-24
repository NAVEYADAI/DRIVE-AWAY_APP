import React from "react";
import TextField from "@mui/material/TextField";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Input from "@mui/material/Input";
import InputAdornment from "@mui/material/InputAdornment";
import IconButton from "@mui/material/IconButton";
import Stack from "@mui/material/Stack";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import Visibility from "@mui/icons-material/Visibility";
import { UserFormData } from "../types";

interface AddUserProps {
    userData: UserFormData;
    setUserData: React.Dispatch<React.SetStateAction<UserFormData>>;
}

export function AddUser({ userData, setUserData }: AddUserProps) {
    const {
        FNameValue,
        LNameValue,
        IDValue,
        UserNameValue,
        PasswordValue,
        Password2Value,
        EmailValue,
        PhoneNumberValue
    } = userData;

    const [showPassword, setShowPassword] = React.useState<boolean>(false);
    const [showPassword2, setShowPassword2] = React.useState<boolean>(false);

    const handleClickShowPassword = () => setShowPassword((show) => !show);
    const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    const handleClickShowPassword2 = () => setShowPassword2((show) => !show);
    const handleMouseDownPassword2 = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    return (
        <Stack direction="column" spacing={2}>
            <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
                <TextField
                    id="FirstName"
                    label="שם פרטי"
                    sx={{ width: '100%' }}
                    value={FNameValue}
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        setUserData({ ...userData, FNameValue: event.target.value });
                    }}
                />
                <TextField
                    id="LastName"
                    label="שם משפחה"
                    sx={{ width: '100%' }}
                    value={LNameValue}
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        setUserData({ ...userData, LNameValue: event.target.value });
                    }}
                />
            </Stack>
            <TextField
                id="ID"
                label="תעודת זהות"
                sx={{ width: '100%' }}
                value={IDValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setUserData({ ...userData, IDValue: event.target.value });
                }}
            />
            <TextField
                id="UserName"
                label="שם משתמש"
                sx={{ width: '100%' }}
                value={UserNameValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setUserData({ ...userData, UserNameValue: event.target.value });
                }}
            />
            <FormControl sx={{ width: '100%' }} variant="outlined">
                <InputLabel htmlFor="standard-adornment-password">סיסמה</InputLabel>
                <Input
                    id="standard-adornment-password"
                    type={showPassword ? 'text' : 'password'}
                    endAdornment={
                        <InputAdornment position="end">
                            <IconButton
                                aria-label="toggle password visibility"
                                onClick={handleClickShowPassword}
                                onMouseDown={handleMouseDownPassword}
                            >
                                {showPassword ? <VisibilityOff /> : <Visibility />}
                            </IconButton>
                        </InputAdornment>
                    }
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        setUserData({ ...userData, PasswordValue: event.target.value });
                    }}
                    value={PasswordValue}
                />
            </FormControl>
            <FormControl sx={{ width: '100%' }} variant="outlined">
                <InputLabel htmlFor="standard-adornment-password2">אימות סיסמה</InputLabel>
                <Input
                    id="standard-adornment-password2"
                    type={showPassword2 ? 'text' : 'password'}
                    endAdornment={
                        <InputAdornment position="end">
                            <IconButton
                                aria-label="toggle password visibility"
                                onClick={handleClickShowPassword2}
                                onMouseDown={handleMouseDownPassword2}
                            >
                                {showPassword2 ? <VisibilityOff /> : <Visibility />}
                            </IconButton>
                        </InputAdornment>
                    }
                    onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                        setUserData({ ...userData, Password2Value: event.target.value });
                    }}
                    value={Password2Value}
                />
            </FormControl>
            <TextField
                id="Email"
                label="כתובת מייל"
                sx={{ width: '100%' }}
                value={EmailValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setUserData({ ...userData, EmailValue: event.target.value });
                }}
            />
            <TextField
                id="Phone"
                label="מספר טלפון"
                sx={{ width: '100%' }}
                value={PhoneNumberValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setUserData({ ...userData, PhoneNumberValue: event.target.value });
                }}
            />
        </Stack>
    );
}
