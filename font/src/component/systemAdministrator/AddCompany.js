import TextField from "@mui/material/TextField";
import React from "react";
export function AddCompany({companyData, setCompanyData}) {
    const {
        CompanyNameValue,
        CompanyIDValue,
        AddressValue,
    } = companyData


    return(
        <>
            <TextField
                id="FirstName"
                label="שם החברה"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={CompanyNameValue}
                onChange={(event) => {
                    setCompanyData({ ...companyData, CompanyNameValue: event.target.value });
                }}
            /><br />
            <TextField
                id="FirstName"
                label="ח.פ. / ת״ז"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={CompanyIDValue}
                onChange={(event) => {
                    setCompanyData({ ...companyData, CompanyIDValue: event.target.value });
                }}
            /><br />

            <TextField
                id="FirstName"
                label="כתובת"
                variant="standard"
                sx={{ m: 1, width: '75%' }}
                value={AddressValue}
                onChange={(event) => {
                    setCompanyData({ ...companyData, AddressValue: event.target.value });
                }}
            />

        </>
    );
}