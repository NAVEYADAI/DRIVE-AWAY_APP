import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import DomainAddIcon from "@mui/icons-material/DomainAdd";
import {AddUser} from "../AddUser";
import React from "react";
import {AddCompany} from "./AddCompany";
import Button from '@mui/material/Button';
import axios from "axios";
import {useGlobalState} from "../../GlobalState";
import { CompanyFormData, UserFormData, MessageResponse } from "../../types";


function CompanyRegistration(){
    const [companyData, setCompanyData] = React.useState<CompanyFormData>({
        CompanyNameValue: '',
        CompanyIDValue: '',
        AddressValue: '',
    })
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
    const send = async () => {
        if (companyData.CompanyNameValue && companyData.CompanyIDValue && companyData.AddressValue && userData.FNameValue && userData.LNameValue &&
            userData.EmailValue && userData.IDValue && userData.UserNameValue && userData.PasswordValue && userData.PhoneNumberValue){
            try {
                const response = await axios.post<MessageResponse>('http://localhost:2909/MainController/AddCompany', {
                    companyName: companyData.CompanyNameValue,
                    companyID: companyData.CompanyIDValue,
                    address: companyData.AddressValue,
                    fName: userData.FNameValue,
                    lName: userData.LNameValue,
                    id: userData.IDValue,
                    userName:userData.UserNameValue,
                    password: userData.PasswordValue,
                    email: userData.EmailValue,
                    phoneNumber: userData.PhoneNumberValue,
                    });
                console.log('Data sent successfully:', response.data);
                alert("החברה נוספה")
            } catch (error) {
                console.error('Error sending data:', error);
                alert("שגיאה בהוספת חברה")
            }
        }
    };
    const { setNameOfPage } = useGlobalState();
    React.useEffect(() => {
        setNameOfPage("הוספת חברה");
    }, [setNameOfPage]);
    return(
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 560, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <DomainAddIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>רישום חברה חדשה</Typography>
                </Stack>
                <Stack direction="column" spacing={3}>
                    <AddCompany companyData={companyData} setCompanyData={setCompanyData}/>
                    <Divider />
                    <AddUser userData={userData} setUserData={setUserData} />
                    <Button variant="contained" size="large" fullWidth onClick={send}>
                        שליחה
                    </Button>
                </Stack>
            </Paper>
        </Stack>
    );
}
export default systemAdministratorPermissions(CompanyRegistration)
