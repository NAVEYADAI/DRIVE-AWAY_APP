import {managerPermissions} from "./managerPermissions";
import Paper from "@mui/material/Paper";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import TextField from "@mui/material/TextField";
import ShuffleIcon from "@mui/icons-material/Shuffle";
import React from "react";
import Button from "@mui/material/Button";
import EditDrives from "./EditDrives";
import EditDrivers from "./EditDrivers";
import axios from "axios";
import {useGlobalState} from "../../GlobalState";
import { MessageResponse } from "../../types";

function Sort(){
    const [NameOfSortValue, setNameOfSortValue] = React.useState<string>("");
    const { level, userId } = useGlobalState();
    const { setNameOfPage } = useGlobalState();
    React.useEffect(() => {
        setNameOfPage("יצירת שיבוץ");
    }, [setNameOfPage]);
    const OnClickSend = async () => {
        try {
            const response = await axios.post<MessageResponse>('http://localhost:2909/Sort', {
                nameOfSort: NameOfSortValue,
                userId: userId,
                level: level
            });
            console.log('Data sent successfully:', response.data);
            alert("השיבוץ נוצר בהצלחה")
        } catch (error) {
            console.error('Error sending data:', error);
            alert("שגיאה ביצירת השיבוץ")
        }
    };
    return(
        <Stack direction="column" alignItems="center" spacing={4} sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 520, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <ShuffleIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>יצירת שיבוץ</Typography>
                </Stack>
                <Stack direction="column" spacing={3}>
                    <TextField
                        id="NameOfSort"
                        label="שם השיבוץ"
                        sx={{ width: '100%' }}
                        value={NameOfSortValue}
                        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                            setNameOfSortValue(event.target.value);
                        }}
                    />
                    <Button variant="contained" size="large" fullWidth onClick={OnClickSend}>
                        שליחה
                    </Button>
                </Stack>
            </Paper>
            <Divider sx={{ width: '100%', maxWidth: 520 }} />
            <EditDrives/>
            <EditDrivers/>
        </Stack>
    );
}

export default managerPermissions(Sort)
