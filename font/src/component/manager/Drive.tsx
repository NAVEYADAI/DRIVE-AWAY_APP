import * as React from 'react';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { TimePicker } from '@mui/x-date-pickers/TimePicker';
import TextField from "@mui/material/TextField";
import Stack from "@mui/material/Stack";
import type { Dayjs } from 'dayjs';
import { DriveFormData } from "../../types";

interface DriveProps {
    DriveValue: DriveFormData;
    setDriveValue: React.Dispatch<React.SetStateAction<DriveFormData>>;
}

export function Drive({DriveValue,setDriveValue}: DriveProps) {
    const {
        NameValue,
        SelectTimeValue,
        StartGpsValue,
        EndGpsValue,
    } = DriveValue
    return(
        <Stack direction="column" spacing={2}>
            <TextField
                id="StartGps"
                label="שם הקו"
                sx={{ width: '100%' }}
                value={NameValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setDriveValue({ ...DriveValue, NameValue: event.target.value });
                }}
            />
            <LocalizationProvider dateAdapter={AdapterDayjs}>
                <TimePicker
                    label="שעת יציאה"
                    value={SelectTimeValue}
                    onChange={(value: Dayjs | null) => {
                        setDriveValue({...DriveValue, SelectTimeValue: value});
                    }}
                    slotProps={{ textField: { fullWidth: true } }}
                />
            </LocalizationProvider>
            <TextField
                id="StartGpsLocation"
                label="מיקום התחלה"
                sx={{ width: '100%' }}
                value={StartGpsValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setDriveValue({ ...DriveValue, StartGpsValue: event.target.value });
                }}
            />
            <TextField
                id="EndGps"
                label="מיקום סיום"
                sx={{ width: '100%' }}
                value={EndGpsValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setDriveValue({ ...DriveValue, EndGpsValue: event.target.value });
                }}
            />
        </Stack>
    );
}
