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

async function fetchDrives(userId, level) {
    const response = await axios.post('http://localhost:2909/MainController/GetDrives', {
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
    const [drives, setDrives] = useState([]);
    const [checked, setChecked] = useState({}); // Use an object to map drive names to their switch states
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("עריכת נסיעות");

    useEffect(() => {
        async function getDrives() {
            try {
                const drivesList = await fetchDrives(userId, level); // העברת המשתנים לפונקציה
                setDrives(drivesList || []);

                // Initialize checked state based on the fetched drives
                const initialChecked = {};
                (drivesList || []).forEach(drive => {
                    initialChecked[drive.nameOfDrive] = drive.aviable;
                });
                setChecked(initialChecked);
            } catch (error) {
                console.error('Error fetching drives:', error);
            }
        }

        getDrives();
    }, [userId, level]);

    const handleToggle = (name) => async () => {
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

    const deleteDriver = (name) => async () => {
        try {
            const response = await axios.post('http://localhost:2909/MainController/DeleteDrive', {
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
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={3}
            useFlexGap
        >
            <Paper elevation={3} sx={{ marginTop: '15px', width: '70%' }}>
                <h2>Edit Drives Page</h2>
                <List sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}>
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
                            <ListItemButton onClick={deleteDriver(item.nameOfDrive)}>
                                <DeleteIcon />
                            </ListItemButton>
                        </ListItem>
                    )) : <ListItemText primary="No drives available" />}
                </List>
            </Paper>
        </Stack>
    );
}

export default managerPermissions(EditDrives);
