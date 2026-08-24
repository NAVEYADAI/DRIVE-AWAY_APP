import * as React from 'react';
import Stack from '@mui/material/Stack';
import FormControlLabel from '@mui/material/FormControlLabel';
import Switch from '@mui/material/Switch';
import {LogIn} from "./LogIn";
import {SingUp} from "./SingUp";

export function MainLogIn() {

    const [auth, setAuth] = React.useState<boolean>(true);
    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setAuth(event.target.checked);
    };
    return (
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <FormControlLabel
                sx={{ mt: 2 }}
                control={<Switch checked={auth} onChange={handleChange} aria-label="login switch" />}
                label={auth ? "הרשמה" : "התחברות"}
            />
            {auth ? <SingUp/> : <LogIn/>}
        </Stack>
    );
}
