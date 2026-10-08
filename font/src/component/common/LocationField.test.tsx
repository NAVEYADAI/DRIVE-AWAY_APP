import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import axios from 'axios';
import { useState } from 'react';
import { LocationField } from './LocationField';

vi.mock('axios');
const mockedAxios = vi.mocked(axios, true);

// jsdom has no WebGL, so mapbox-gl is replaced by a fake map that lets the test
// "click" a coordinate.
const mapbox = vi.hoisted(() => ({
    clickHandler: null as null | ((e: { lngLat: { lng: number; lat: number } }) => void),
    centers: [] as unknown[],
    flights: [] as unknown[],
}));
vi.mock('mapbox-gl', () => {
    class Map {
        constructor(opts: { center: unknown }) {
            mapbox.centers.push(opts.center);
        }
        addControl() {}
        flyTo(opts: unknown) { mapbox.flights.push(opts); }
        on(event: string, handler: typeof mapbox.clickHandler) {
            if (event === 'click') mapbox.clickHandler = handler;
        }
        remove() {}
    }
    class Marker {
        setLngLat() { return this; }
        addTo() { return this; }
    }
    class NavigationControl {}
    return { default: { Map, Marker, NavigationControl, accessToken: '' } };
});

function Harness({ initial = '' }: { initial?: string }) {
    const [value, setValue] = useState(initial);
    return <LocationField id="loc" label="מיקום התחלה" value={value} onChange={setValue} />;
}

describe('LocationField', () => {
    beforeEach(() => {
        vi.stubEnv('VITE_MAPBOX_TOKEN', 'pk.test');
        mapbox.clickHandler = null;
        mapbox.centers = [];
        mapbox.flights = [];
        vi.clearAllMocks();
    });
    afterEach(() => vi.unstubAllEnvs());

    it('still lets the user type an address by hand', async () => {
        const user = userEvent.setup();
        render(<Harness />);
        await user.type(screen.getByLabelText('מיקום התחלה'), 'דיזנגוף 1');
        expect(screen.getByLabelText('מיקום התחלה')).toHaveValue('דיזנגוף 1');
    });

    it('picking a point on the map fills the field as "lat,lng"', async () => {
        const user = userEvent.setup();
        render(<Harness />);

        await user.click(screen.getByRole('button', { name: 'בחירת מיקום התחלה במפה' }));
        await screen.findByTestId('map-container');
        const confirm = await screen.findByRole('button', { name: 'אישור' });
        expect(confirm).toBeDisabled();

        await vi.waitFor(() => expect(mapbox.clickHandler).not.toBeNull());
        mapbox.clickHandler!({ lngLat: { lng: 34.7818123, lat: 32.0853 } });
        expect(await screen.findByText('מיקום שנבחר: 32.085300,34.781812')).toBeInTheDocument();

        await user.click(screen.getByRole('button', { name: 'אישור' }));
        expect(screen.getByLabelText('מיקום התחלה')).toHaveValue('32.085300,34.781812');
        expect(screen.queryByTestId('map-container')).not.toBeInTheDocument();
    });

    it('searching an address moves the map there and fills the field with its coordinates', async () => {
        const user = userEvent.setup();
        mockedAxios.get.mockResolvedValueOnce({
            data: { features: [
                { properties: { full_address: 'דיזנגוף 1, תל אביב' }, geometry: { coordinates: [34.7743, 32.0781] } },
                { properties: { full_address: 'דיזנגוף 1, חיפה' }, geometry: { coordinates: [34.99, 32.8] } },
            ] },
        });
        render(<Harness />);

        await user.click(screen.getByRole('button', { name: 'בחירת מיקום התחלה במפה' }));
        await user.type(await screen.findByLabelText('חיפוש כתובת'), 'דיזנגוף 1{Enter}');

        expect(mockedAxios.get).toHaveBeenCalledWith(
            'https://api.mapbox.com/search/geocode/v6/forward',
            { params: expect.objectContaining({ q: 'דיזנגוף 1', country: 'il', language: 'he', permanent: true }) },
        );
        await user.click(await screen.findByText('דיזנגוף 1, תל אביב'));
        expect(mapbox.flights).toEqual([{ center: [34.7743, 32.0781], zoom: 16 }]);
        expect(screen.getByText('מיקום שנבחר: 32.078100,34.774300')).toBeInTheDocument();

        await user.click(screen.getByRole('button', { name: 'אישור' }));
        expect(screen.getByLabelText('מיקום התחלה')).toHaveValue('32.078100,34.774300');
    });

    it('tells the user when the search finds nothing or fails', async () => {
        const user = userEvent.setup();
        mockedAxios.get.mockResolvedValueOnce({ data: { features: [] } });
        mockedAxios.get.mockRejectedValueOnce(new Error('network'));
        vi.spyOn(console, 'error').mockImplementation(() => {});
        render(<Harness />);

        await user.click(screen.getByRole('button', { name: 'בחירת מיקום התחלה במפה' }));
        const box = await screen.findByLabelText('חיפוש כתובת');
        await user.type(box, 'xyz{Enter}');
        expect(await screen.findByText('לא נמצאו תוצאות')).toBeInTheDocument();

        await user.click(screen.getByRole('button', { name: 'חיפוש' }));
        expect(await screen.findByText('שגיאה בחיפוש כתובת')).toBeInTheDocument();
        expect(screen.getByRole('button', { name: 'אישור' })).toBeDisabled();
    });

    it('cancel leaves the field unchanged', async () => {
        const user = userEvent.setup();
        render(<Harness initial="דיזנגוף 1" />);

        await user.click(screen.getByRole('button', { name: 'בחירת מיקום התחלה במפה' }));
        await user.click(await screen.findByRole('button', { name: 'ביטול' }));

        expect(screen.getByLabelText('מיקום התחלה')).toHaveValue('דיזנגוף 1');
    });

    it('opens the map centred on an existing "lat,lng" value', async () => {
        const user = userEvent.setup();
        render(<Harness initial="32.0,34.8" />);

        await user.click(screen.getByRole('button', { name: 'בחירת מיקום התחלה במפה' }));
        await screen.findByTestId('map-container');

        await vi.waitFor(() => expect(mapbox.centers).toContainEqual([34.8, 32]));
    });

    it('disables the map button when no Mapbox token is configured', () => {
        vi.stubEnv('VITE_MAPBOX_TOKEN', '');
        render(<Harness />);
        expect(screen.getByRole('button', { name: 'בחירת מיקום התחלה במפה' })).toBeDisabled();
        expect(screen.getByLabelText('מיקום התחלה')).toBeEnabled();
    });
});
