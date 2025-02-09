import React from "react";
import axios from "axios";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import SendIcon from "@mui/icons-material/Send";
import Chip from '@mui/material/Chip';
import LoginIcon from '@mui/icons-material/Login';
import { useNavigate } from "react-router-dom";
import { useGlobalState } from "../GlobalState";

export function LogIn() {
    const [UserNameValue, setUserNameValue] = React.useState();
    const [PasswordValue, setPasswordValue] = React.useState();
    const nav = useNavigate();
    const { setLevel, setUserId } = useGlobalState();
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("התחברות")
    const onCLickLogInSendButton = async () => {
        console.log(UserNameValue);
        console.log(PasswordValue);
        if (UserNameValue && PasswordValue) {
            try {
                const response = await axios.post('http://localhost:2909/login', {
                    UserName: UserNameValue,
                    Password: PasswordValue,
                    IP: UserNameValue,
                    MacAddress: PasswordValue
                });
                setLevel(response.data['level']);
                setUserId(response.data['UserId']);

                console.log('Data sent successfully:', response.data);
                alert("התחברת")
                if (response.data['level'] === 1 || response.data['level'] === 2) {
                    nav('/systemAdministrator');
                } else if (response.data['level'] === 3 || response.data['level'] === 4) {
                    nav('/manager');
                }
            } catch (error) {
                alert("שגיאת התחברות")
                console.error('Error sending data:', error);
            }
        }
    };

    return (
        <div>
            <br /><br />
            <Stack
                direction="column"
                justifyContent="center"
                alignItems="center"
                spacing={3}
                useFlexGap
                fullWidth
            >
                <Paper sx={{ marginTop: '15px', width: '75%' }}>
                    <Chip label="Log In" variant="outlined" icon={<LoginIcon />} />
                    <br /><br />
                    <TextField
                        id="UserId"
                        label="User Name"
                        value={UserNameValue}
                        onChange={(event) => setUserNameValue(event.target.value)}
                        fullWidth
                    /><br /><br />
                    <TextField
                        id="PasswordId"
                        label="Password"
                        value={PasswordValue}
                        onChange={(event) => setPasswordValue(event.target.value)}
                        fullWidth
                    /><br /><br />
                    <Button variant="contained" endIcon={<SendIcon />} onClick={onCLickLogInSendButton}>
                        Send
                    </Button>
                    <br /><br /><br />
                </Paper>
            </Stack>
        </div>
    );
}
