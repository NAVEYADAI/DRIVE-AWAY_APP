import {managerPermissions} from "./managerPermissions";
import Paper from "@mui/material/Paper";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import React from "react";
import SendIcon from "@mui/icons-material/Send";
import Button from "@mui/material/Button";
import EditDrives from "./EditDrives";
import EditDrivers from "./EditDrivers";
import axios from "axios";
import {useGlobalState} from "../../GlobalState";

function Sort(){
    const [NameOfSortValue, setNameOfSortValue] = React.useState("");
    const { level, userId } = useGlobalState();
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("יצירת שיבוץ")
    const OnClickSend = async () => {
        // if ( DriverHeaderValue.IsEmployeeValue && DriverHeaderValue.AddressValue &&userData.FNameValue &&
        //     userData.LNameValue && userData.EmailValue && userData.IDValue && userData.UserNameValue &&
        //     userData.PasswordValue && userData.PhoneNumberValue){
        try {
            const response = await axios.post('http://localhost:2909/Sort', {
                nameOfSort: NameOfSortValue,
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
                       sx={{marginTop: '15px', width: '70%'}}
                >
                    <h1>Sort page</h1>
                    <TextField
                        id="NameOfSort"
                        label="שם השיבוץ"
                        variant="standard"
                        sx={{m: 1, width: '75%'}}
                        value={NameOfSortValue}
                        onChange={(event) => {
                            setNameOfSortValue(event.target.value);
                        }}
                    /><br/>

                    <EditDrives/><br/>
                    <EditDrivers/><br/>

                    <Button variant="contained" endIcon={<SendIcon/>} onClick={OnClickSend}>
                        Send
                    </Button>
                </Paper>
            </Stack>
        </>
    );
}

export default managerPermissions(Sort)