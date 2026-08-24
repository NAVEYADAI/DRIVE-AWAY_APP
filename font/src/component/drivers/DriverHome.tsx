import { driverPermissions } from "./DriverPermission";
import Paper from "@mui/material/Paper";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
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
import LocalShippingIcon from '@mui/icons-material/LocalShipping';
import NavigationIcon from '@mui/icons-material/Navigation';
import { LoadingScreen } from "../common/LoadingScreen";

interface Drive {
    Name: string;
    HToStart: number;
    MToStart: number;
    GpsStart: string;
    GpsEnd: string;
}

interface GetMyDrivesResponse {
    message: string;
    "list of drives": Drive[];
}

async function fetchDrives(userId: number, level: number): Promise<Drive[]> {
    const response = await axios.post<GetMyDrivesResponse>('http://localhost:2909/MainController/GetMyDrives', {
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
    const [drives, setDrives] = React.useState<Drive[]>([]);
    const [loading, setLoading] = React.useState<boolean>(true);
    const { setNameOfPage } = useGlobalState();
    useEffect(() => {
        setNameOfPage("הנסיעות שלי");
    }, [setNameOfPage]);
    useEffect(() => {
        async function getDrives() {
            try {
                const drivesList = await fetchDrives(userId, level);
                setDrives(drivesList);
            } catch (error) {
                console.error('Error fetching drives:', error);
            } finally {
                setLoading(false);
            }
        }

        getDrives();
    }, [userId, level]);

    const handleNavigate = (gpsEnd: string) => {
        const url = `https://www.google.com/maps/dir/?api=1&destination=${gpsEnd}`;
        window.open(url, '_blank');
    };

    return (
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 960, p: { xs: 2, sm: 3 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 2 }}>
                    <LocalShippingIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>הנסיעות שלי</Typography>
                </Stack>
                {loading ? (
                    <LoadingScreen label="טוען נסיעות..." />
                ) : drives.length === 0 ? (
                    <Typography variant="body2" color="text.secondary" align="center" sx={{ py: 2 }}>
                        אין נסיעות זמינות
                    </Typography>
                ) : (
                <TableContainer>
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
                                            startIcon={<NavigationIcon />}
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
                                            startIcon={<NavigationIcon />}
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
                )}
            </Paper>
        </Stack>
    );
}

export default driverPermissions(DriverHome);
