import TextField from "@mui/material/TextField";
import React from "react";
import Switch from '@mui/material/Switch';

export function DriverHeader({DriverHeaderValue, setDriverHeaderValue}) {
const {
    AddressValue,
    IsEmployeeValue,

}= DriverHeaderValue
    return(
        <>
            <TextField
                id="AddressValue"
                label="כתובת"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={AddressValue}
                onChange={(event) => {
                    setDriverHeaderValue({ ...DriverHeaderValue, AddressValue: event.target.value });
                }}
            /><br />
            <Switch  defaultChecked
                     value={IsEmployeeValue}
                     onChange={(event) => {
                setDriverHeaderValue({ ...DriverHeaderValue, IsEmployeeValue: event.target.value });
            }}/>

        </>
    );
}