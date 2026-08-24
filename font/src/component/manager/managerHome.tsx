import {managerPermissions} from "./managerPermissions";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import GroupIcon from "@mui/icons-material/Group";
import PersonAddIcon from "@mui/icons-material/PersonAdd";
import AddRoadIcon from "@mui/icons-material/AddRoad";
import RouteIcon from "@mui/icons-material/Route";
import ShuffleIcon from "@mui/icons-material/Shuffle";
import EditNoteIcon from "@mui/icons-material/EditNote";
import LocalShippingIcon from "@mui/icons-material/LocalShipping";
import {useNavigate} from "react-router-dom";
import { useEffect } from "react";
import { ActionCard } from "../common/ActionCard";
import { useGlobalState } from "../../GlobalState";

function ManagerHome(){
    const nav = useNavigate()
    const { setNameOfPage } = useGlobalState();
    useEffect(() => {
        setNameOfPage("ניהול חברה");
    }, [setNameOfPage]);

    return(
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Box sx={{ width: '100%', maxWidth: 720 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 4 }}>
                    <GroupIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>ניהול חברה</Typography>
                    <Typography variant="body2" color="text.secondary">
                        מה תרצי לעשות?
                    </Typography>
                </Stack>
                <Box
                    sx={{
                        display: 'grid',
                        gridTemplateColumns: 'repeat(auto-fit, minmax(140px, 1fr))',
                        gap: 2,
                    }}
                >
                    <ActionCard icon={PersonAddIcon} label="הוספת נהג" onClick={() => nav('/manager/AddDriver')} />
                    <ActionCard icon={GroupIcon} label="עריכת נהגים" onClick={() => nav('/manager/EditDrivers')} />
                    <ActionCard icon={AddRoadIcon} label="הוספת נסיעה" onClick={() => nav('/manager/AddDrive')} />
                    <ActionCard icon={RouteIcon} label="עריכת נסיעות" onClick={() => nav('/manager/EditDrives')} />
                    <ActionCard icon={ShuffleIcon} label="יצירת שיבוץ" onClick={() => nav('/manager/Sort')} />
                    <ActionCard icon={EditNoteIcon} label="עריכת שיבוצים" onClick={() => nav('/manager/EditSort')} />
                    <ActionCard icon={LocalShippingIcon} label="הנסיעות שלי" onClick={() => nav('/driver')} variant="primary" />
                </Box>
            </Box>
        </Stack>
    );
}
export default managerPermissions(ManagerHome)
