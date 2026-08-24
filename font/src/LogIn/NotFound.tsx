import Stack from '@mui/material/Stack';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import SearchOffIcon from '@mui/icons-material/SearchOff';
import { useNavigate } from 'react-router-dom';

export function NotFound() {
    const nav = useNavigate();
    return (
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 4, sm: 10 }, width: '100%', maxWidth: 420, p: 4, borderRadius: 3, textAlign: 'center' }}>
                <Stack direction="column" spacing={2} alignItems="center">
                    <SearchOffIcon color="disabled" sx={{ fontSize: 48 }} />
                    <Typography variant="h6" fontWeight={600}>הדף לא נמצא</Typography>
                    <Button variant="contained" onClick={() => nav('/LogIn')}>חזרה לדף הבית</Button>
                </Stack>
            </Paper>
        </Stack>
    );
}
