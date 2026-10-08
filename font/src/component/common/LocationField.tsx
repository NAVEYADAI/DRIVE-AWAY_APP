import * as React from "react";
import TextField from "@mui/material/TextField";
import InputAdornment from "@mui/material/InputAdornment";
import IconButton from "@mui/material/IconButton";
import Tooltip from "@mui/material/Tooltip";
import MapIcon from "@mui/icons-material/Map";

const MapPickerDialog = React.lazy(() => import("./MapPickerDialog"));

interface LocationFieldProps {
    id: string;
    label: string;
    value: string;
    onChange: (value: string) => void;
}

// שדה מיקום: אפשר להקליד כתובת או "lat,lng", או לבחור נקודה במפה (נשמרת כ-"lat,lng").
export function LocationField({ id, label, value, onChange }: LocationFieldProps) {
    const [open, setOpen] = React.useState(false);
    const hasToken = Boolean(import.meta.env.VITE_MAPBOX_TOKEN);

    return (
        <>
            <TextField
                id={id}
                label={label}
                sx={{ width: '100%' }}
                value={value}
                onChange={(event: React.ChangeEvent<HTMLInputElement>) => onChange(event.target.value)}
                InputProps={{
                    endAdornment: (
                        <InputAdornment position="end">
                            <Tooltip title={hasToken ? "בחירה במפה" : "בחירה במפה אינה זמינה (חסר VITE_MAPBOX_TOKEN)"}>
                                <span>
                                    <IconButton
                                        aria-label={`בחירת ${label} במפה`}
                                        edge="end"
                                        disabled={!hasToken}
                                        onClick={() => setOpen(true)}
                                    >
                                        <MapIcon />
                                    </IconButton>
                                </span>
                            </Tooltip>
                        </InputAdornment>
                    ),
                }}
            />
            {open && (
                <React.Suspense fallback={null}>
                    <MapPickerDialog
                        open
                        title={`בחירת ${label} במפה`}
                        initialValue={value}
                        onClose={() => setOpen(false)}
                        onConfirm={(picked) => {
                            onChange(picked);
                            setOpen(false);
                        }}
                    />
                </React.Suspense>
            )}
        </>
    );
}
