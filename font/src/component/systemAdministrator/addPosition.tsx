import { useEffect } from "react";
import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import BadgeIcon from "@mui/icons-material/Badge";
import * as React from 'react';
import {AddUser} from "../AddUser";
import Button from '@mui/material/Button';
import {useGlobalState} from "../../GlobalState";
import { UserFormData } from "../../types";

interface CompanyOption {
    name: string;
}

function GetAllCompany(): CompanyOption[] {
    // קבלת כל השמות חברות הרשומות
    const listCompany = [{name:"nave"}, {name:"yadai"} ]
    // קבלה מהapi את רשימת כל החברות
    return listCompany
}
function AddPosition (){
    var listCompany = GetAllCompany()
    const [CompanyName, setCompanyName] = React.useState<string>('');

    const handleChange = (event: SelectChangeEvent) => {
        setCompanyName(event.target.value);
    };
    const { setNameOfPage } = useGlobalState();
    useEffect(() => {
        setNameOfPage("הוספת עמדה בחברה קיימת");
    }, [setNameOfPage]);

    // קביעת משתנים עבור נתוני המשתמש
    const [userData, setUserData] = React.useState<UserFormData>({
        FNameValue: '',
        LNameValue: '',
        IDValue: '',
        UserNameValue: '',
        PasswordValue: '',
        Password2Value: '',
        EmailValue: '',
        PhoneNumberValue: ''
    });
    return(
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 520, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <BadgeIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>רישום עמדה בחברה קיימת</Typography>
                </Stack>
                <Stack direction="column" spacing={3}>
                    <FormControl fullWidth>
                        <InputLabel id="CampnyName">שם החברה</InputLabel>
                        <Select
                            labelId="CampnyName"
                            id="demo-simple-select"
                            value={CompanyName}
                            label="שם החברה"
                            onChange={handleChange}
                        >
                            {
                                listCompany.map((item)=>{
                                    return(
                                        <MenuItem key={item.name} value={item.name}>{item.name}</MenuItem>
                                    );
                                })
                            }
                        </Select>
                    </FormControl>
                    <Divider />
                    <AddUser userData={userData} setUserData={setUserData} />
                    <Button variant="contained" size="large" fullWidth>
                        שליחה
                    </Button>
                </Stack>
            </Paper>
        </Stack>
    );
}
export default systemAdministratorPermissions(AddPosition);
