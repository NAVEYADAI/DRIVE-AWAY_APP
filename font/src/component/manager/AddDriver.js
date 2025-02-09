import {managerPermissions} from "./managerPermissions";
import * as React from "react";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import {AddUser} from "../AddUser";
import {DriverHeader} from "./DriverHeader";
import SendIcon from "@mui/icons-material/Send";
import Button from "@mui/material/Button";
import axios from "axios";
import {useGlobalState} from "../../GlobalState";

function AddDriver(){
    const { level, userId } = useGlobalState();
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("הוספת נהג")
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
    // קביעת משתנים עבור נתוני המשתמש
    const [DriverHeaderValue, setDriverHeaderValue] = React.useState({
        AddressValue:'',
        IsEmployeeValue:false,
    });
    const OnClickSend = async () => {
        // if ( DriverHeaderValue.IsEmployeeValue && DriverHeaderValue.AddressValue &&userData.FNameValue &&
        //     userData.LNameValue && userData.EmailValue && userData.IDValue && userData.UserNameValue &&
        //     userData.PasswordValue && userData.PhoneNumberValue){
            try {
                const response = await axios.post('http://localhost:2909/MainController/AddDriver', {
                    address: DriverHeaderValue.AddressValue,
                    isEmployee: DriverHeaderValue.IsEmployeeValue,
                    fName: userData.FNameValue,
                    lName: userData.LNameValue,
                    id: userData.IDValue,
                    userName:userData.UserNameValue,
                    password: userData.PasswordValue,
                    email: userData.EmailValue,
                    phoneNumber: userData.PhoneNumberValue,
                    userId: userId,
                    level: level
                });
                //todo
                console.log('Data sent successfully:', response.data);
                alert("נהג נוסף הבצלחה")
            } catch (error) {
                console.error('Error sending data:', error);
                alert("שגיאה בהוספת נהג")
            }
        // }
    };
    return(
        <>
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
                    AddDriver page
                    <AddUser userData={userData} setUserData={setUserData} />
                    <DriverHeader DriverHeaderValue={DriverHeaderValue} setDriverHeaderValue={setDriverHeaderValue} />
                    שכיר? כן
                    <br/>
                    <Button variant="contained" endIcon={<SendIcon />} onClick={OnClickSend}>
                        Send
                    </Button>
                </Paper>
            </Stack>

        </>
    );
}
export default managerPermissions(AddDriver)