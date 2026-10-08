import { describe, it, expect } from 'vitest';
import { formatLatLng, parseLatLng } from './latLng';

describe('latLng', () => {
    it('formats as "lat,lng" (latitude first, the order the backend parses)', () => {
        expect(formatLatLng(32.0853, 34.7818)).toBe('32.085300,34.781800');
    });

    it('parses "lat,lng" with spaces', () => {
        expect(parseLatLng(' 32.0 , 34.8 ')).toEqual({ lat: 32, lng: 34.8 });
    });

    it('rejects free-text addresses and out-of-range values', () => {
        for (const bad of ['', 'דיזנגוף 1, תל אביב', '91,34', '32,181', '1,2,3', ',5', '5,']) {
            expect(parseLatLng(bad), bad).toBeNull();
        }
    });
});
