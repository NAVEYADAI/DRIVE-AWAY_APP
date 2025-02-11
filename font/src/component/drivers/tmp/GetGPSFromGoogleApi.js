// GeocodingExample.js
import React, { useState } from 'react';
import axios from 'axios';
import Paper from '@mui/material/Paper';
import InputBase from '@mui/material/InputBase';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import SearchIcon from '@mui/icons-material/Search';

export function GeocodingExample({ onLocationUpdate }) {
    const [address, setAddress] = useState('');
    const [latitude, setLatitude] = useState('');
    const [longitude, setLongitude] = useState('');

    const handleAddressChange = (event) => {
        setAddress(event.target.value);
    };

    const handleGeocode = async () => {
        try {
            const response = await axios.get('https://maps.googleapis.com/maps/api/geocode/json', {
                params: {
                    address: address,
                    key: process.env.MY_GOOGLE_API_KEY // my api code
                }
            });

            if (response.data.status === 'OK') {
                const location = response.data.results[0].geometry.location;
                setLatitude(location.lat);
                setLongitude(location.lng);
                onLocationUpdate(location.lat, location.lng);
            } else {
                console.error('Failed to geocode address');
            }
        } catch (error) {
            console.error('Error fetching geocoding data:', error);
        }
    };

    return (
        <div>
            <Paper
                component="form"
                sx={{
                    p: '2px 4px',
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                    flexWrap: 'wrap',
                }}
            >
                <IconButton sx={{ p: '10px' }} aria-label="menu">
                    <MenuIcon />
                </IconButton>
                <InputBase
                    sx={{ ml: 1, flex: 1 }}
                    placeholder="Search Google Maps"
                    inputProps={{ 'aria-label': 'search google maps' }}
                    value={address}
                    onChange={handleAddressChange}
                />
                <IconButton type="button" sx={{ p: '10px' }} aria-label="search" onClick={handleGeocode}>
                    <SearchIcon />
                </IconButton>
            </Paper>
        </div>
    );
}

export default GeocodingExample;
