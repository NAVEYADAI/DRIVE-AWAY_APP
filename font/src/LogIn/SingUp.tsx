import Stack from '@mui/material/Stack';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import React from "react";
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import PersonAddIcon from '@mui/icons-material/PersonAdd';

export function SingUp() {

    const [UserNameValue, setUserNameValue] = React.useState<string>('');
    const [PasswordValue, setPasswordValue] = React.useState<string>('');
    const [ReturnPasswordValue, setReturnPasswordValue] = React.useState<string>('');
    const [FNameValue, setFNameValue] = React.useState<string>('');
    const [LNameValue, setLNameValue] = React.useState<string>('');
    const [PhoneValue, setPhoneValue] = React.useState<string>('');
    const [EmailValue, setEmailValue] = React.useState<string>('');

    return (
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper
                elevation={3}
                sx={{
                    mt: { xs: 2, sm: 6 },
                    width: '100%',
                    maxWidth: 460,
                    p: { xs: 3, sm: 4 },
                    borderRadius: 3,
                }}
            >
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <PersonAddIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>הרשמה למערכת</Typography>
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
                    <TextField
                        id="ReturnPasswordId"
                        label="אימות סיסמה"
                        type="password"
                        value={ReturnPasswordValue}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => setReturnPasswordValue(event.target.value)}
                        fullWidth
                    />
                    <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
                        <TextField
                            id="FNameId"
                            label="שם פרטי"
                            value={FNameValue}
                            onChange={(event: React.ChangeEvent<HTMLInputElement>) => setFNameValue(event.target.value)}
                            fullWidth
                        />
                        <TextField
                            id="LNameId"
                            label="שם משפחה"
                            value={LNameValue}
                            onChange={(event: React.ChangeEvent<HTMLInputElement>) => setLNameValue(event.target.value)}
                            fullWidth
                        />
                    </Stack>
                    <TextField
                        id="PhoneId"
                        label="טלפון"
                        value={PhoneValue}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => setPhoneValue(event.target.value)}
                        fullWidth
                    />
                    <TextField
                        id="EmailId"
                        label="דוא״ל"
                        value={EmailValue}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => setEmailValue(event.target.value)}
                        fullWidth
                    />
                    <Button variant="contained" size="large" sx={{ mt: 1 }}>
                        הרשמה
                    </Button>
                </Stack>
            </Paper>
        </Stack>
    );
}
