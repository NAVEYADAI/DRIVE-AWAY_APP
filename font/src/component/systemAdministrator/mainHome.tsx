import { useEffect } from "react";
import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import Stack from "@mui/material/Stack";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import ApartmentIcon from "@mui/icons-material/Apartment";
import DomainAddIcon from "@mui/icons-material/DomainAdd";
import BadgeIcon from "@mui/icons-material/Badge";
import AssessmentIcon from "@mui/icons-material/Assessment";
import {useNavigate} from "react-router-dom";
import {useGlobalState} from "../../GlobalState";
import { ActionCard } from "../common/ActionCard";

function SystemAdministrator() {
    const nav = useNavigate()
    const { setNameOfPage } = useGlobalState();
    useEffect(() => {
        setNameOfPage("מנהל חברה");
    }, [setNameOfPage]);
    return (
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Box sx={{ width: '100%', maxWidth: 640 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 4 }}>
                    <ApartmentIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>מנהל חברה</Typography>
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
                    <ActionCard icon={DomainAddIcon} label="רישום חברה" onClick={() => nav('/systemAdministrator/CompanyRegistration')} />
                    <ActionCard icon={BadgeIcon} label="רישום עמדה בחברה קיימת" onClick={() => nav('/systemAdministrator/AddPosition')} />
                    <ActionCard icon={AssessmentIcon} label="דוח מספר שיבוצים" onClick={() => nav('/systemAdministrator/NumberOfInlays')} />
                    <ActionCard icon={ApartmentIcon} label="חברה שלי" onClick={() => nav('/manager')} variant="primary" />
                </Box>
            </Box>
        </Stack>
    );
}

export default systemAdministratorPermissions(SystemAdministrator);
