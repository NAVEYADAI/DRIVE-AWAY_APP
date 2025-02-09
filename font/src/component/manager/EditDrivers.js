import React, { useState, useEffect } from 'react';
import { managerPermissions } from './managerPermissions';
import Paper from '@mui/material/Paper';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import Stack from '@mui/material/Stack';
import Switch from '@mui/material/Switch';
import ListItemButton from '@mui/material/ListItemButton';
import DeleteIcon from '@mui/icons-material/Delete';
import axios from 'axios';
import { useGlobalState } from '../../GlobalState';

async function fetchDrivers(userId, level) {
    const response = await axios.post('http://localhost:2909/MainController/GetDrivers', {
        userId: userId,
        level: level,
    });
    const data = response.data;
    if (data.message === 'Success') {
        return data.List;
    }
    throw new Error('Failed to fetch drivers');
}

function EditDrivers() {
    const { level, userId } = useGlobalState();
    const [drivers, setDrivers] = useState([]);
    const [checked, setChecked] = useState({}); // Use an object to map driver names to their switch states
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("עריכת נהגים")
    useEffect(() => {
        async function getDrivers() {
            try {
                const driversList = await fetchDrivers(userId, level);
                setDrivers(driversList || []);

                // Initialize checked state based on the fetched drivers
                const initialChecked = {};
                driversList.forEach(driver => {
                    initialChecked[driver.nameOfDriver] = driver.aviable;
                });
                setChecked(initialChecked);
            } catch (error) {
                console.error('Error fetching drivers:', error);
            }
        }

        getDrivers();
    }, [userId, level]);

    const handleToggle = (name) => async () => {
        setChecked(prevChecked => ({
            ...prevChecked,
            [name]: !prevChecked[name]
        }));
        await axios.post('http://localhost:2909/MainController/ViewOrNotDriver', {
            driverName:name,
            userId: userId,
            level: level,
        });
    };

    const deleteDriver = (name) => async () => {
        await axios.post('http://localhost:2909/MainController/DeleteDriver', {
            driverName: name,
            userId: userId,
            level: level,
        });
    };

    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={3}
            useFlexGap
        >
            <Paper elevation={3} sx={{ marginTop: '15px', width: '70%' }}>
                <h2>EditDrivers page</h2>
                <List sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}>
                    {drivers.length > 0 ? drivers.map((item) => (
                        <ListItem key={item.nameOfDriver}>
                            <ListItemText primary={item.nameOfDriver} />
                            <Switch
                                edge="end"
                                onChange={handleToggle(item.nameOfDriver)}
                                checked={checked[item.nameOfDriver] || false}
                                inputProps={{
                                    'aria-labelledby': `switch-list-label-${item.nameOfDriver}`,
                                }}
                            />
                            <ListItemButton onClick={deleteDriver(item.nameOfDriver)}>
                                <DeleteIcon />
                            </ListItemButton>
                        </ListItem>
                    )) : <ListItemText primary="No drivers available" />}
                </List>

            </Paper>
        </Stack>
    );
}

export default managerPermissions(EditDrivers);
