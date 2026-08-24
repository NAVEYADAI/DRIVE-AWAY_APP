import React, { useState } from "react";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import AddRoadIcon from "@mui/icons-material/AddRoad";
import axios from "axios";
import dayjs from "dayjs";
import { useGlobalState } from "../../GlobalState";
import { managerPermissions } from "./managerPermissions";
import { Drive } from "./Drive";
import { DriveFormData, MessageResponse } from "../../types";

function AddDrive() {
    const [DriveValue, setDriveValue] = useState<DriveFormData>({
        NameValue: '',
        SelectTimeValue: null,
        StartGpsValue: '',
        EndGpsValue: '',
    });

    const { level, userId } = useGlobalState();
    const { setNameOfPage } = useGlobalState();
    React.useEffect(() => {
        setNameOfPage("הוספת נסיעה");
    }, [setNameOfPage]);
    const OnClickSend = async () => {
        if (!DriveValue.SelectTimeValue || !DriveValue.SelectTimeValue.isValid()) {
            alert("יש לבחור שעת יציאה");
            return;
        }
        const hours = dayjs(DriveValue.SelectTimeValue).hour();
        const minutes = dayjs(DriveValue.SelectTimeValue).minute();

        try {
            const response = await axios.post<MessageResponse>('http://localhost:2909/MainController/AddDrive', {
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
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 480, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <AddRoadIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>הוספת מסלול</Typography>
                </Stack>
                <Stack direction="column" spacing={3}>
                    <Drive DriveValue={DriveValue} setDriveValue={setDriveValue} />
                    <Button variant="contained" size="large" fullWidth onClick={OnClickSend}>
                        שליחה
                    </Button>
                </Stack>
            </Paper>
        </Stack>
    );
}

export default managerPermissions(AddDrive);
