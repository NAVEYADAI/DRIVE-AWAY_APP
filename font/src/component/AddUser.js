import React from "react";
import TextField from "@mui/material/TextField";
import FormControl from "@mui/material/FormControl";
import InputLabel from "@mui/material/InputLabel";
import Input from "@mui/material/Input";
import InputAdornment from "@mui/material/InputAdornment";
import IconButton from "@mui/material/IconButton";
import VisibilityOff from "@mui/icons-material/VisibilityOff";
import Visibility from "@mui/icons-material/Visibility";

export function AddUser({ userData, setUserData }) {
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

    const [showPassword, setShowPassword] = React.useState(false);
    const [showPassword2, setShowPassword2] = React.useState(false);

    const handleClickShowPassword = () => setShowPassword((show) => !show);
    const handleMouseDownPassword = (event) => {
        event.preventDefault();
    };

    const handleClickShowPassword2 = () => setShowPassword2((show) => !show);
    const handleMouseDownPassword2 = (event) => {
        event.preventDefault();
    };

    return (
        <>
            <br />
            <TextField
                id="FirstName"
                label="שם פרטי"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={FNameValue}
                onChange={(event) => {
                    setUserData({ ...userData, FNameValue: event.target.value });
                }}
            /><br />
            <TextField
                id="LastName"
                label="שם משפחה"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={LNameValue}
                onChange={(event) => {
                    setUserData({ ...userData, LNameValue: event.target.value });
                }}
            /><br />
            <TextField
                id="ID"
                label="תעודת זהות"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={IDValue}
                onChange={(event) => {
                    setUserData({ ...userData, IDValue: event.target.value });
                }}
            /><br />
            <TextField
                id="UserName"
                label="שם משתמש"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={UserNameValue}
                onChange={(event) => {
                    setUserData({ ...userData, UserNameValue: event.target.value });
                }}
            /><br />
            <FormControl sx={{ m: 1, width: '75%' }} variant="standard">
                <InputLabel htmlFor="standard-adornment-password">סיסמא</InputLabel>
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
                    onChange={(event) => {
                        setUserData({ ...userData, PasswordValue: event.target.value });
                    }}
                    value={PasswordValue}
                />
            </FormControl><br />
            <FormControl sx={{ m: 1, width: '75%' }} variant="standard">
                <InputLabel htmlFor="standard-adornment-password">אימות סיסמא</InputLabel>
                <Input
                    id="standard-adornment-password"
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
                    onChange={(event) => {
                        setUserData({ ...userData, Password2Value: event.target.value });
                    }}
                    value={Password2Value}
                />
            </FormControl><br />
            <TextField
                id="Email"
                label="כתובת מייל"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={EmailValue}
                onChange={(event) => {
                    setUserData({ ...userData, EmailValue: event.target.value });
                }}
            /><br />
            <TextField
                id="Phone"
                label="מספר טלפון"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={PhoneNumberValue}
                onChange={(event) => {
                    setUserData({ ...userData, PhoneNumberValue: event.target.value });
                }}
            /><br />
        </>
    );
}
