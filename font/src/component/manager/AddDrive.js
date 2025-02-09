import React, { useState } from "react";
import SendIcon from "@mui/icons-material/Send";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import axios from "axios";
import dayjs from "dayjs";
import { useGlobalState } from "../../GlobalState";
import { managerPermissions } from "./managerPermissions";
import { Drive } from "./Drive";

function AddDrive() {
    const [DriveValue, setDriveValue] = useState({
        NameValue: '',
        SelectTimeValue: null,
        StartGpsValue: '',
        EndGpsValue: '',
    });

    const { level, userId } = useGlobalState();
    // console.log(level)
    // console.log(userId)
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("הוספת נסיעה")
    const OnClickSend = async () => {
        const hours = dayjs(DriveValue.SelectTimeValue).hour();
        const minutes = dayjs(DriveValue.SelectTimeValue).minute();

        try {
            const response = await axios.post('http://localhost:2909/MainController/AddDrive', {
                nameOfDriver: DriveValue.NameValue,
                timeToStartH: hours,
                timeToStartM: minutes,
                gpsStart: DriveValue.StartGpsValue,
                gpsEnd: DriveValue.EndGpsValue,
                userId: userId,
                level: level
            });
            console.log('Data sent successfully:', response.data);
            alert("הנסיעה בוצעה בהצלחה");
        } catch (error) {
            console.error('Error sending data:', error);
            alert("שגיאה בהוספת נסיעה");
        }
    };

    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={3}
            useFlexGap
        >
            <Paper elevation={3} sx={{ marginTop: '15px', width: '70%' }}>
                <br /><br />
                <h2>הוספת מסלול</h2>
                <br />
                <Drive DriveValue={DriveValue} setDriveValue={setDriveValue} />
                <br /><br />
                <Button variant="contained" endIcon={<SendIcon />} onClick={OnClickSend}>
                    Send
                </Button>
                <br /><br />
            </Paper>
        </Stack>
    );
}

export default managerPermissions(AddDrive);
