import Stack from "@mui/material/Stack";
import CircularProgress from "@mui/material/CircularProgress";
import Typography from "@mui/material/Typography";

interface LoadingScreenProps {
    label?: string;
}

export function LoadingScreen({ label = "טוען..." }: LoadingScreenProps) {
    return (
        <Stack direction="column" alignItems="center" spacing={2} sx={{ py: 6, width: '100%' }}>
            <CircularProgress color="primary" />
            <Typography variant="body2" color="text.secondary">
                {label}
            </Typography>
        </Stack>
    );
}
