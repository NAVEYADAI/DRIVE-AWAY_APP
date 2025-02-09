import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import {AddUser} from "../AddUser";
import React from "react";
import {AddCompany} from "./AddCompany";
import Button from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';
import axios from "axios";
import {useGlobalState} from "../../GlobalState";
// import {useGlobalState} from "../../GlobalState";


function CompanyRegistration(){
    // const { level, userId } = useGlobalState();

    //רישום חברה
    // name = "רישום חברה חדשה"
    // קביעת משתנים עבור נתונים על חברה
    const [companyData, setCompanyData] = React.useState({
        CompanyNameValue: '',
        CompanyIDValue: '',
        AddressValue: '',
    })
    // קביעת משתנים עבור נתוני המשתמש
    const [userData, setUserData] = React.useState({
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
                const response = await axios.post('http://localhost:2909/MainController/AddCompany', {
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
                    // userId: userId,
                    // level: level
                    });
                //todo
                console.log('Data sent successfully:', response.data);
                alert("החברה נוספה")
            } catch (error) {
                console.error('Error sending data:', error);
                alert("שגיאה בהוספת חברה")
            }
        }
    };
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("הוספת חברה")
    return(<>
        <Stack

            direction="column"
            justifyContent="center"
            alignItems="center"

            spacing={3}
            useFlexGap
        >
            <Paper elevation={3}
                   sx = {{marginTop: '15px', width:'70%'}}
            >

                <AddCompany companyData={companyData} setCompanyData={setCompanyData}/>
                <AddUser userData={userData} setUserData={setUserData} />
                <Button variant="contained" endIcon={<SendIcon />} onClick={send}>
                    Send
                </Button>
            </Paper>
        </Stack>
    </>);
}
export default systemAdministratorPermissions(CompanyRegistration)