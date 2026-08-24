import React, { useState, useEffect } from 'react';
import { managerPermissions } from './managerPermissions';
import Paper from '@mui/material/Paper';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Switch from '@mui/material/Switch';
import ListItemButton from '@mui/material/ListItemButton';
import DeleteIcon from '@mui/icons-material/Delete';
import GroupIcon from '@mui/icons-material/Group';
import axios from 'axios';
import { useGlobalState } from '../../GlobalState';
import { DriverListItem, GetDriversResponse } from '../../types';
import { LoadingScreen } from '../common/LoadingScreen';

async function fetchDrivers(userId: number, level: number): Promise<DriverListItem[]> {
    const response = await axios.post<GetDriversResponse>('http://localhost:2909/MainController/GetDrivers', {
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
    const [drivers, setDrivers] = useState<DriverListItem[]>([]);
    const [checked, setChecked] = useState<Record<string, boolean>>({}); // Use an object to map driver names to their switch states
    const [loading, setLoading] = useState<boolean>(true);
    const { setNameOfPage } = useGlobalState();
    useEffect(() => {
        setNameOfPage("עריכת נהגים");
    }, [setNameOfPage]);
    useEffect(() => {
        async function getDrivers() {
            try {
                const driversList = await fetchDrivers(userId, level);
                setDrivers(driversList || []);

                // Initialize checked state based on the fetched drivers
                const initialChecked: Record<string, boolean> = {};
                driversList.forEach(driver => {
                    initialChecked[driver.nameOfDriver] = driver.aviable;
                });
                setChecked(initialChecked);
            } catch (error) {
                console.error('Error fetching drivers:', error);
            } finally {
                setLoading(false);
            }
        }

        getDrivers();
    }, [userId, level]);

    const handleToggle = (name: string) => async () => {
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

    const deleteDriver = (name: string) => async () => {
        await axios.post('http://localhost:2909/MainController/DeleteDriver', {
            driverName: name,
            userId: userId,
            level: level,
        });
    };

    return (
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 520, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 2 }}>
                    <GroupIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>עריכת נהגים</Typography>
                </Stack>
                {loading ? (
                    <LoadingScreen label="טוען נהגים..." />
                ) : (
                <List sx={{ width: '100%' }}>
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
                            <ListItemButton onClick={deleteDriver(item.nameOfDriver)} sx={{ flexGrow: 0 }}>
                                <DeleteIcon color="error" />
                            </ListItemButton>
                        </ListItem>
                    )) : (
                        <Typography variant="body2" color="text.secondary" align="center" sx={{ py: 2 }}>
                            אין נהגים זמינים
                        </Typography>
                    )}
                </List>
                )}
            </Paper>
        </Stack>
    );
}

export default managerPermissions(EditDrivers);
