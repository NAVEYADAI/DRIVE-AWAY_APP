import React from "react";
import axios from "axios";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import LoginIcon from '@mui/icons-material/Login';
import Typography from "@mui/material/Typography";
import { useNavigate } from "react-router-dom";
import { useGlobalState } from "../GlobalState";

interface LoginResponse {
    messege: string;
    UserId: number;
    level: number;
}

export function LogIn() {
    const [UserNameValue, setUserNameValue] = React.useState<string>('');
    const [PasswordValue, setPasswordValue] = React.useState<string>('');
    const nav = useNavigate();
    const { setLevel, setUserId } = useGlobalState();
    const { setNameOfPage } = useGlobalState();
    React.useEffect(() => {
        setNameOfPage("התחברות");
    }, [setNameOfPage]);
    const onCLickLogInSendButton = async () => {
        if (UserNameValue && PasswordValue) {
            try {
                const response = await axios.post<LoginResponse>('http://localhost:2909/login', {
                    UserName: UserNameValue,
                    Password: PasswordValue,
                    IP: UserNameValue,
                    MacAddress: PasswordValue
                });
                setLevel(response.data.level);
                setUserId(response.data.UserId);

                if (response.data.level === 1 || response.data.level === 2) {
                    nav('/systemAdministrator');
                } else if (response.data.level === 3 || response.data.level === 4) {
                    nav('/manager');
                } else if (response.data.level === 5) {
                    nav('/driver');
                }
            } catch (error) {
                alert("שגיאת התחברות")
                console.error('Error sending data:', error);
            }
        }
    };

    return (
        <Stack
            direction="column"
            alignItems="center"
            sx={{ width: '100%' }}
        >
            <Paper
                elevation={3}
                sx={{
                    mt: { xs: 2, sm: 6 },
                    width: '100%',
                    maxWidth: 420,
                    p: { xs: 3, sm: 4 },
                    borderRadius: 3,
                }}
            >
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <LoginIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>כניסה למערכת</Typography>
                </Stack>
                <Stack direction="column" spacing={2.5}>
                    <TextField
                        id="UserId"
                        label="שם משתמש"
                        value={UserNameValue}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => setUserNameValue(event.target.value)}
                        fullWidth
                    />
                    <TextField
                        id="PasswordId"
                        label="סיסמה"
                        type="password"
                        value={PasswordValue}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => setPasswordValue(event.target.value)}
                        fullWidth
                    />
                    <Button
                        variant="contained"
                        size="large"
                        onClick={onCLickLogInSendButton}
                        sx={{ mt: 1 }}
                    >
                        כניסה
                    </Button>
                </Stack>
            </Paper>
        </Stack>
    );
}
