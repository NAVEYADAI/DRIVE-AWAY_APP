import * as React from "react";
import axios from "axios";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemText from "@mui/material/ListItemText";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import mapboxgl from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";
import { formatLatLng, parseLatLng } from "./latLng";

// מרכז ישראל, כשעדיין אין מיקום שנבחר
const ISRAEL_CENTER: [number, number] = [34.85, 31.5];

interface SearchResult {
    label: string;
    lat: number;
    lng: number;
}

interface MapPickerDialogProps {
    open: boolean;
    title: string;
    initialValue: string;
    onClose: () => void;
    onConfirm: (value: string) => void;
}

// נטען עצמאית (React.lazy) כדי ש-mapbox-gl לא יכנס לבאנדל הראשי.
export default function MapPickerDialog({ open, title, initialValue, onClose, onConfirm }: MapPickerDialogProps) {
    const [container, setContainer] = React.useState<HTMLDivElement | null>(null);
    const [picked, setPicked] = React.useState<string>('');
    const [query, setQuery] = React.useState('');
    const [results, setResults] = React.useState<SearchResult[]>([]);
    const [searchMessage, setSearchMessage] = React.useState('');
    const mapRef = React.useRef<mapboxgl.Map | null>(null);
    const markerRef = React.useRef<mapboxgl.Marker | null>(null);

    const pick = React.useCallback((lat: number, lng: number) => {
        const map = mapRef.current;
        if (!map) return;
        if (!markerRef.current) markerRef.current = new mapboxgl.Marker().setLngLat([lng, lat]).addTo(map);
        else markerRef.current.setLngLat([lng, lat]);
        setPicked(formatLatLng(lat, lng));
    }, []);

    React.useEffect(() => {
        if (!open || !container) return;
        const start = parseLatLng(initialValue);
        setPicked('');

        mapboxgl.accessToken = import.meta.env.VITE_MAPBOX_TOKEN ?? '';
        const map = new mapboxgl.Map({
            container,
            style: 'mapbox://styles/mapbox/streets-v12',
            center: start ? [start.lng, start.lat] : ISRAEL_CENTER,
            zoom: start ? 14 : 7,
        });
        mapRef.current = map;
        map.addControl(new mapboxgl.NavigationControl(), 'top-left');
        if (start) pick(start.lat, start.lng);
        map.on('click', (e) => pick(e.lngLat.lat, e.lngLat.lng));
        return () => {
            map.remove();
            mapRef.current = null;
            markerRef.current = null;
        };
    }, [open, container, initialValue, pick]);

    const search = async () => {
        const q = query.trim();
        if (!q) return;
        setSearchMessage('מחפש...');
        setResults([]);
        try {
            // permanent=true: הקואורדינטות נשמרות ב-DB, ותנאי Mapbox מתירים שמירה רק עם הפרמטר הזה.
            const response = await axios.get('https://api.mapbox.com/search/geocode/v6/forward', {
                params: {
                    q,
                    country: 'il',
                    language: 'he',
                    limit: 5,
                    permanent: true,
                    access_token: import.meta.env.VITE_MAPBOX_TOKEN,
                },
            });
            const found: SearchResult[] = (response.data?.features ?? []).map((f: any) => ({
                label: f.properties?.full_address || f.properties?.name || q,
                lng: f.geometry.coordinates[0],
                lat: f.geometry.coordinates[1],
            }));
            setResults(found);
            setSearchMessage(found.length ? '' : 'לא נמצאו תוצאות');
        } catch (error) {
            console.error('Error searching address:', error);
            setSearchMessage('שגיאה בחיפוש כתובת');
        }
    };

    const choose = (r: SearchResult) => {
        mapRef.current?.flyTo({ center: [r.lng, r.lat], zoom: 16 });
        pick(r.lat, r.lng);
        setResults([]);
    };

    return (
        <Dialog open={open} onClose={onClose} fullWidth maxWidth="md">
            <DialogTitle>{title}</DialogTitle>
            <DialogContent>
                <Stack direction="row" spacing={1} sx={{ mb: 1, mt: 1 }}>
                    <TextField
                        size="small"
                        fullWidth
                        label="חיפוש כתובת"
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        onKeyDown={(e) => {
                            if (e.key === 'Enter') {
                                e.preventDefault();
                                search();
                            }
                        }}
                    />
                    <Button variant="outlined" onClick={search}>חיפוש</Button>
                </Stack>
                {searchMessage && (
                    <Typography variant="body2" color="text.secondary" sx={{ mb: 1 }}>{searchMessage}</Typography>
                )}
                {results.length > 0 && (
                    <List dense disablePadding sx={{ mb: 1, border: 1, borderColor: 'divider', borderRadius: 2 }}>
                        {results.map((r, i) => (
                            <ListItemButton key={`${r.lat},${r.lng},${i}`} onClick={() => choose(r)}>
                                <ListItemText primary={r.label} />
                            </ListItemButton>
                        ))}
                    </List>
                )}
                <Box ref={setContainer} data-testid="map-container" sx={{ height: 420, width: '100%', borderRadius: 2 }} />
                <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
                    {picked ? `מיקום שנבחר: ${picked}` : 'חפשו כתובת או לחצו על המפה כדי לבחור מיקום'}
                </Typography>
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose}>ביטול</Button>
                <Button variant="contained" disabled={!picked} onClick={() => onConfirm(picked)}>
                    אישור
                </Button>
            </DialogActions>
        </Dialog>
    );
}
