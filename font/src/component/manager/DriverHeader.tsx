import { LocationField } from "../common/LocationField";
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
            <LocationField
                id="AddressValue"
                label="כתובת"
                value={AddressValue}
                onChange={(value) => setDriverHeaderValue({ ...DriverHeaderValue, AddressValue: value })}
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
