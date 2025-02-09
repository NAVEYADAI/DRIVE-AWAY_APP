import * as React from 'react';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { TimePicker } from '@mui/x-date-pickers/TimePicker';
import TextField from "@mui/material/TextField";



// import dayjs from 'dayjs';
// import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
//


export function Drive({DriveValue,setDriveValue}) {
    const {
        NameValue,
        SelectTimeValue,
        StartGpsValue,
        EndGpsValue,

    } = DriveValue
    return(
        <>

            <TextField
                id="StartGps"
                label="שם הקן"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={NameValue}
                onChange={(event) => {
                    setDriveValue({ ...DriveValue, NameValue: event.target.value });
                }}
            /><br /><br />
            <LocalizationProvider dateAdapter={AdapterDayjs} sx={{ m: 1, width: '75%' }}>
                <TimePicker
                    label="Date Picker"
                    variant="standard"
                    renderInput={(params)=><TextField {...params}/>}
                    value={SelectTimeValue}
                    onChange={(event) => {
                        setDriveValue({...DriveValue, SelectTimeValue:event});
                        // setSelectTime(event);
                        console.log(event)
                    }}
                />
            </LocalizationProvider>
            <br/>
            <TextField
                id="StartGps"
                label="מיקום התחלה"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={StartGpsValue}
                onChange={(event) => {
                    setDriveValue({ ...DriveValue, StartGpsValue: event.target.value });
                }}
            /><br />
            <TextField
                id="EndGps"
                label="מיקום סיום"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={EndGpsValue}
                onChange={(event) => {
                    setDriveValue({ ...DriveValue, EndGpsValue: event.target.value });
                }}
            /><br />
        </>
    );
}