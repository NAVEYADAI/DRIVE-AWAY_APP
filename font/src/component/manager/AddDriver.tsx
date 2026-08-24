import {managerPermissions} from "./managerPermissions";
import * as React from "react";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import PersonAddIcon from "@mui/icons-material/PersonAdd";
import {AddUser} from "../AddUser";
import {DriverHeader} from "./DriverHeader";
import Button from "@mui/material/Button";
import axios from "axios";
import {useGlobalState} from "../../GlobalState";
import { UserFormData, DriverHeaderData, MessageResponse } from "../../types";

function AddDriver(){
    const { level, userId } = useGlobalState();
    const { setNameOfPage } = useGlobalState();
    React.useEffect(() => {
        setNameOfPage("הוספת נהג");
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
    // קביעת משתנים עבור נתוני המשתמש
    const [DriverHeaderValue, setDriverHeaderValue] = React.useState<DriverHeaderData>({
        AddressValue:'',
        IsEmployeeValue:false,
    });
    const OnClickSend = async () => {
        try {
            const response = await axios.post<MessageResponse>('http://localhost:2909/MainController/AddDriver', {
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
            console.log('Data sent successfully:', response.data);
            alert("נהג נוסף בהצלחה")
        } catch (error) {
            console.error('Error sending data:', error);
            alert("שגיאה בהוספת נהג")
        }
    };
    return(
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 520, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <PersonAddIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>הוספת נהג</Typography>
                </Stack>
                <Stack direction="column" spacing={3}>
                    <AddUser userData={userData} setUserData={setUserData} />
                    <Divider />
                    <DriverHeader DriverHeaderValue={DriverHeaderValue} setDriverHeaderValue={setDriverHeaderValue} />
                    <Button variant="contained" size="large" fullWidth onClick={OnClickSend}>
                        שליחה
                    </Button>
                </Stack>
            </Paper>
        </Stack>
    );
}
export default managerPermissions(AddDriver)
