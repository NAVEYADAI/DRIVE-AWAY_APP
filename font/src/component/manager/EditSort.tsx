import {managerPermissions} from "./managerPermissions";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import EditNoteIcon from "@mui/icons-material/EditNote";

function EditSort(){
    return(
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 480, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center">
                    <EditNoteIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>עריכת שיבוצים</Typography>
                    <Typography variant="body2" color="text.secondary">
                        מסך זה נמצא בבנייה
                    </Typography>
                </Stack>
            </Paper>
        </Stack>
    );
}
export default managerPermissions(EditSort)
