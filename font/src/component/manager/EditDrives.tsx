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
import RouteIcon from '@mui/icons-material/Route';
import axios from 'axios';
import { useGlobalState } from '../../GlobalState';
import { DriveListItem, GetDrivesResponse, MessageResponse } from '../../types';
import { LoadingScreen } from '../common/LoadingScreen';

async function fetchDrives(userId: number, level: number): Promise<DriveListItem[]> {
    const response = await axios.post<GetDrivesResponse>('http://localhost:2909/MainController/GetDrives', {
        userId: userId,
        level: level,
    });
    const data = response.data;
    if (data.message === 'Success') {
        return data.List;
    }
    throw new Error('Failed to fetch drives');
}

function EditDrives() {
    const { level, userId } = useGlobalState();
    const [drives, setDrives] = useState<DriveListItem[]>([]);
    const [checked, setChecked] = useState<Record<string, boolean>>({}); // Use an object to map drive names to their switch states
    const [loading, setLoading] = useState<boolean>(true);
    const { setNameOfPage } = useGlobalState();
    useEffect(() => {
        setNameOfPage("עריכת נסיעות");
    }, [setNameOfPage]);

    useEffect(() => {
        async function getDrives() {
            try {
                const drivesList = await fetchDrives(userId, level); // העברת המשתנים לפונקציה
                setDrives(drivesList || []);

                // Initialize checked state based on the fetched drives
                const initialChecked: Record<string, boolean> = {};
                (drivesList || []).forEach(drive => {
                    initialChecked[drive.nameOfDrive] = drive.aviable;
                });
                setChecked(initialChecked);
            } catch (error) {
                console.error('Error fetching drives:', error);
            } finally {
                setLoading(false);
            }
        }

        getDrives();
    }, [userId, level]);

    const handleToggle = (name: string) => async () => {
        const newChecked = !checked[name];
        setChecked(prevChecked => ({
            ...prevChecked,
            [name]: newChecked
        }));

        try {
            await axios.post('http://localhost:2909/MainController/ViewOrNotDrive', {
                driveName: name,
                userId: userId,
                level: level,
            });
            console.log(`Drive ${name} availability updated successfully`);
        } catch (error) {
            console.error('Error updating drive availability:', error);
        }
    };

    const deleteDriver = (name: string) => async () => {
        try {
            const response = await axios.post<MessageResponse>('http://localhost:2909/MainController/DeleteDrive', {
                driveName: name,
                userId: userId,
                level: level,
            });
            console.log('Data sent successfully:', response.data);
            alert("הנסיעה נמחקה בהצלחה");
        } catch (error) {
            console.error('Error deleting drive:', error);
            alert("שגיאה במחיקת נסיעה");
        }
    };

    return (
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 520, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 2 }}>
                    <RouteIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>עריכת נסיעות</Typography>
                </Stack>
                {loading ? (
                    <LoadingScreen label="טוען נסיעות..." />
                ) : (
                <List sx={{ width: '100%' }}>
                    {drives.length > 0 ? drives.map((item) => (
                        <ListItem key={item.nameOfDrive}>
                            <ListItemText primary={item.nameOfDrive} />
                            <Switch
                                edge="end"
                                onChange={handleToggle(item.nameOfDrive)}
                                checked={checked[item.nameOfDrive] || false}
                                inputProps={{
                                    'aria-labelledby': `switch-list-label-${item.nameOfDrive}`,
                                }}
                            />
                            <ListItemButton onClick={deleteDriver(item.nameOfDrive)} sx={{ flexGrow: 0 }}>
                                <DeleteIcon color="error" />
                            </ListItemButton>
                        </ListItem>
                    )) : (
                        <Typography variant="body2" color="text.secondary" align="center" sx={{ py: 2 }}>
                            אין נסיעות זמינות
                        </Typography>
                    )}
                </List>
                )}
            </Paper>
        </Stack>
    );
}

export default managerPermissions(EditDrives);
