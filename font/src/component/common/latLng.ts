// הפורמט שהשרת מצפה לו (back/internal/sorting/MapboxApi.go::parseLatLng): "lat,lng".
// מיקום בפורמט הזה עובר ישירות ל-Mapbox בלי geocoding.

export function formatLatLng(lat: number, lng: number): string {
    return `${lat.toFixed(6)},${lng.toFixed(6)}`;
}

export function parseLatLng(value: string): { lat: number; lng: number } | null {
    const parts = value.split(',');
    if (parts.length !== 2) return null;
    const lat = Number(parts[0].trim());
    const lng = Number(parts[1].trim());
    if (parts[0].trim() === '' || parts[1].trim() === '') return null;
    if (!Number.isFinite(lat) || !Number.isFinite(lng)) return null;
    if (lat < -90 || lat > 90 || lng < -180 || lng > 180) return null;
    return { lat, lng };
}
