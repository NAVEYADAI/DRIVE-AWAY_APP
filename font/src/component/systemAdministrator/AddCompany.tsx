import TextField from "@mui/material/TextField";
import Stack from "@mui/material/Stack";
import React from "react";
import { CompanyFormData } from "../../types";

interface AddCompanyProps {
    companyData: CompanyFormData;
    setCompanyData: React.Dispatch<React.SetStateAction<CompanyFormData>>;
}

export function AddCompany({companyData, setCompanyData}: AddCompanyProps) {
    const {
        CompanyNameValue,
        CompanyIDValue,
        AddressValue,
    } = companyData


    return(
        <Stack direction="column" spacing={2}>
            <TextField
                id="CompanyName"
                label="שם החברה"
                sx={{ width: '100%' }}
                value={CompanyNameValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setCompanyData({ ...companyData, CompanyNameValue: event.target.value });
                }}
            />
            <TextField
                id="CompanyID"
                label="ח.פ. / ת״ז"
                sx={{ width: '100%' }}
                value={CompanyIDValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setCompanyData({ ...companyData, CompanyIDValue: event.target.value });
                }}
            />
            <TextField
                id="Address"
                label="כתובת"
                sx={{ width: '100%' }}
                value={AddressValue}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setCompanyData({ ...companyData, AddressValue: event.target.value });
                }}
            />
        </Stack>
    );
}
