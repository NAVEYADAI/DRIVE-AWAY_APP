import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import ButtonBase from "@mui/material/ButtonBase";
import Box from "@mui/material/Box";
import { SvgIconComponent } from "@mui/icons-material";
import { alpha } from "@mui/material/styles";

interface ActionCardProps {
    icon: SvgIconComponent;
    label: string;
    onClick: () => void;
    variant?: "default" | "primary";
}

export function ActionCard({ icon: Icon, label, onClick, variant = "default" }: ActionCardProps) {
    const isPrimary = variant === "primary";
    return (
        <ButtonBase
            onClick={onClick}
            sx={{ borderRadius: 3, width: '100%', display: 'block', textAlign: 'inherit' }}
        >
            <Paper
                elevation={isPrimary ? 3 : 1}
                sx={{
                    p: 2.5,
                    borderRadius: 3,
                    width: '100%',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                    gap: 1,
                    border: '1px solid',
                    borderColor: isPrimary ? 'primary.main' : 'divider',
                    bgcolor: isPrimary ? 'primary.main' : 'background.paper',
                    color: isPrimary ? 'primary.contrastText' : 'text.primary',
                    transition: 'transform 0.15s ease, box-shadow 0.15s ease',
                    '&:hover': {
                        transform: 'translateY(-2px)',
                        boxShadow: 4,
                    },
                }}
            >
                <Box
                    sx={{
                        width: 48,
                        height: 48,
                        borderRadius: '50%',
                        display: 'flex',
                        alignItems: 'center',
                        justifyContent: 'center',
                        bgcolor: isPrimary ? 'rgba(255,255,255,0.15)' : (theme) => alpha(theme.palette.primary.main, 0.1),
                    }}
                >
                    <Icon sx={{ fontSize: 26, color: isPrimary ? 'inherit' : 'primary.main' }} />
                </Box>
                <Typography variant="body2" fontWeight={600} textAlign="center">
                    {label}
                </Typography>
            </Paper>
        </ButtonBase>
    );
}
