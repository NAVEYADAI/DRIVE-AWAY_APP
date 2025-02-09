import { driverPermissions } from "./DriverPermission";
import Paper from "@mui/material/Paper";
import Stack from "@mui/material/Stack";
import React, { useEffect } from "react";
import { useGlobalState } from "../../GlobalState";
import axios from "axios";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Button from '@mui/material/Button';


async function fetchDrives(userId, level) {
    const response = await axios.post('http://localhost:2909/MainController/GetMyDrives', {
        userId: userId,
        level: level,
    });
    const data = response.data;
    if (data.message === 'work') {
        return data["list of drives"];
    }
    throw new Error('Failed to fetch drives');
}

function DriverHome() {
    const { level, userId } = useGlobalState();
    const [drives, setDrives] = React.useState([]);
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("הנסיעות שלי")
    useEffect(() => {
        async function getDrives() {
            try {
                const drivesList = await fetchDrives(userId, level);
                setDrives(drivesList);
            } catch (error) {
                console.error('Error fetching drives:', error);
            }
        }

        getDrives();
    }, [userId, level]);

    const handleNavigate = (gpsEnd) => {
        const url = `https://www.google.com/maps/dir/?api=1&destination=${gpsEnd}`;
        window.open(url, '_blank');
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
                <TableContainer component={Paper}>
                    <Table sx={{ minWidth: 650 }} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">ניווט לנק סיום</TableCell>
                                <TableCell align="center">GPS סיום</TableCell>
                                <TableCell align="center">ניווט לנק התחלה</TableCell>
                                <TableCell align="center">GPS התחלה</TableCell>
                                <TableCell align="center">שעת התחלה</TableCell>
                                <TableCell align="center">שם</TableCell>

                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {drives.map((drive, index) => (
                                <TableRow key={index}>
                                    <TableCell align="center">
                                        <Button
                                            variant="contained"
                                            color="primary"
                                            onClick={() => handleNavigate(drive.GpsEnd)}
                                        >
                                            ניווט
                                        </Button>
                                    </TableCell>
                                    <TableCell align="center">{drive.GpsEnd}</TableCell>
                                    <TableCell align="center">
                                        <Button
                                            variant="contained"
                                            color="primary"
                                            onClick={() => handleNavigate(drive.GpsStart)}
                                        >
                                            ניווט
                                        </Button>
                                    </TableCell>
                                    <TableCell align="center">{drive.GpsStart}</TableCell>
                                    <TableCell align="center">{`${drive.HToStart}:${drive.MToStart}`}</TableCell>
                                    <TableCell align="center">{drive.Name}</TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Paper>
        </Stack>
    );
}

export default driverPermissions(DriverHome);
