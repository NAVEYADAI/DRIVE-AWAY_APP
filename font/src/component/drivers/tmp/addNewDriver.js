// Driver.js
import React, { useState } from 'react';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import GeocodingExample from "./GetGPSFromGoogleApi";
import FormControlLabel from '@mui/material/FormControlLabel';
import Switch from '@mui/material/Switch';
import Button from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';
import IconButton from '@mui/material/IconButton';
import Input from '@mui/material/Input';
import InputLabel from '@mui/material/InputLabel';
import InputAdornment from '@mui/material/InputAdornment';
import FormControl from '@mui/material/FormControl';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
export function AddNewDriver() {

    const [driverLatitude, setDriverLatitude] = useState('');
    const [driverLongitude, setDriverLongitude] = useState('');

    const [FNameValue, setFNameValue] = React.useState('');
    const [LNameValue, setLNameValue] = React.useState('');
    const [IDValue, setIDValue] = React.useState('');
    const [UserNameValue, setUserNameValue] = React.useState('');
    const [PasswordValue, setPasswordValue] = React.useState('');
    const [Password2Value,setPassword2Value] = React.useState('');
    const [EmailValue, setEmailValue] = React.useState('');
    const [PhoneNumberValue, setPhoneNumberValue] = React.useState('');
    const [GPSValue, setGPSValue] = React.useState('');
    const [EmployeeValue, setEmployeeValue] = React.useState(false);



    const [showPassword, setShowPassword] = React.useState(false);

    const handleClickShowPassword = () => setShowPassword((show) => !show);

    const handleMouseDownPassword = (event) => {
        event.preventDefault();
    };

    const [auth, setAuth] = React.useState(true);
    const SwitchHandelChange = (event) => {
        setAuth(event.target.checked);
        setEmployeeValue(event.target.checked)
    };

    const SendButton = async () => {
        const userData = {
            UserName: UserNameValue,
            FName: FNameValue,
            LName: LNameValue,
            IdentityCard: IDValue,
            Password: PasswordValue,
            Email: EmailValue,
            Phone: PhoneNumberValue
        };

        try {
            const response = await fetch('http://localhost:2909/login/SingUp', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(userData)
            });

            if (response.ok) {
                // טופס נשלח בהצלחה, ניתן להוסיף לוגיקה נוספת כאן
                // להוסיף הודעה קופצת של הצלחת
                //todo להוסיף רישום של נהג לחברה
                //todo לנקות את הטופס
                console.log("work")
            } else {
                // אירעה שגיאה בשליחת הטופס
                console.error('Error sending form data');

            }
        } catch (error) {
            console.error('Error:', error);
        }
    }

    return (
        <div>
            <Box
                sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                    flexWrap: 'wrap',
                    marginTop: '80px',
                    '& > :not(style)': {
                        width: '65%',
                        height: '75%',
                    },
                }}
            >
                <Paper elevation={3}>
                    <h1>רישום נהג חדש למערכת</h1>
                    <TextField id="FirstName" label="שם פרטי" variant="standard" sx={{m: 1, width: '75%'}} value={FNameValue}
                               onChange={(event) => {
                                   setFNameValue(event.target.value);
                               }}/><br/>
                    <TextField id="LastName" label="שם משפחה" variant="standard" sx={{m: 1, width: '75%'}} value={LNameValue}
                               onChange={(event) => {
                                   setLNameValue(event.target.value);
                               }}/><br/>
                    <TextField id="ID" label="תעודת זהות"
                               variant="standard"
                               sx={{m: 1, width: '75%'}} value={IDValue}
                               onChange={(event) => {
                                   setIDValue(event.target.value);
                               }}/><br/>
                    <TextField id="UserName" label="שם משתמש"
                               variant="standard"
                               sx={{m: 1, width: '75%'}}
                               value={UserNameValue}
                               onChange={(event) => {
                                   setUserNameValue(event.target.value);
                               }}/><br/>
                    <FormControl sx={{m: 1, width: '75%'}} variant="standard">
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
                                        {showPassword ? <VisibilityOff/> : <Visibility/>}
                                    </IconButton>
                                </InputAdornment>
                            }
                        />
                    </FormControl><br/>
                    <FormControl sx={{m: 1, width: '75%'}} variant="standard">
                        <InputLabel htmlFor="standard-adornment-password">אימות סיסמא</InputLabel>
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
                                        {showPassword ? <VisibilityOff/> : <Visibility/>}
                                    </IconButton>
                                </InputAdornment>
                            }
                        />
                    </FormControl><br/>
                    <TextField id="Email" label="כתובת מייל" variant="standard" sx={{m: 1, width: '75%'}} value={EmailValue}
                               onChange={(event) => {
                                   setEmailValue(event.target.value);
                               }}/><br/>
                    <TextField id="Phone" label="מספר טלפון" variant="standard" sx={{m: 1, width: '75%'}} value={PhoneNumberValue}
                               onChange={(event) => {
                                   setPhoneNumberValue(event.target.value);
                               }}/><br/>
                    <Paper
                        elevation={0}
                        sx={{
                            display: 'flex',
                            justifyContent: 'center',
                            alignItems: 'center',
                            textAlign: 'center' // הוספת תכונת textAlign למרכז את התוכן
                        }}>
                        <GeocodingExample
                            onLocationUpdate={(latitude, longitude) => {
                                setDriverLatitude(latitude);
                                setDriverLongitude(longitude);
                            }}
                        />
                    </Paper><br/>
                    <FormControlLabel checked={auth}  control={<Switch defaultChecked/>}
                                      label="שכיר" value={EmployeeValue}
                                      onChange={(event) => {
                                          SwitchHandelChange(event);
                                      }}/><br/><br/>
                    <Button variant="contained" endIcon={<SendIcon/>} onClick={SendButton}>
                        Send
                    </Button><br/><br/>
                </Paper>

            </Box>
            {/*{driverLatitude && driverLongitude && (*/}
            {/*    <div>*/}
            {/*        <p>Driver Latitude: {driverLatitude}</p>*/}
            {/*        <p>Driver Longitude: {driverLongitude}</p>*/}
            {/*    </div>*/}
            {/*)}*/}
        </div>
    );
}
