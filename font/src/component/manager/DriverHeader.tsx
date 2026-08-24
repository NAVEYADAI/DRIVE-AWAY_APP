import TextField from "@mui/material/TextField";
import React from "react";
import Stack from "@mui/material/Stack";
import Switch from '@mui/material/Switch';
import FormControlLabel from '@mui/material/FormControlLabel';
import { DriverHeaderData } from "../../types";

interface DriverHeaderProps {
    DriverHeaderValue: DriverHeaderData;
    setDriverHeaderValue: React.Dispatch<React.SetStateAction<DriverHeaderData>>;
}

export function DriverHeader({DriverHeaderValue, setDriverHeaderValue}: DriverHeaderProps) {
const {
    AddressValue,
    IsEmployeeValue,
}= DriverHeaderValue
    return(
        <Stack direction="column" spacing={2}>
            <TextField
                id="AddressValue"
                label="כתובת"
                sx={{ width: '100%' }}
                value={AddressValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setDriverHeaderValue({ ...DriverHeaderValue, AddressValue: event.target.value });
                }}
            />
            <FormControlLabel
                control={
                    <Switch
                        checked={IsEmployeeValue}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                            setDriverHeaderValue({ ...DriverHeaderValue, IsEmployeeValue: event.target.checked });
                        }}
                    />
                }
                label="שכיר"
            />
        </Stack>
    );
}
