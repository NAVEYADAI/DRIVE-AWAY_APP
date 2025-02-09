import {managerPermissions} from "./managerPermissions";
import Paper from "@mui/material/Paper";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import {useNavigate} from "react-router-dom";

function ManagerHome(){
    const nav = useNavigate()
    const AddDriverClick = () => {
        nav('/manager/AddDriver')
    }
    const EditDriversClick = () => {
        nav('/manager/EditDrivers')
    }
    const AddDriveClick = () => {
        nav('/manager/AddDrive')
    }
    const EditDrivesClick = () => {
        nav('/manager/EditDrives')
    }
    const SortClick = () => {
        nav('/manager/Sort')
    }
    const EditSortsClick = () => {
        nav('/manager/EditSort')
    }
    const MyDrives = () => {
        nav('/driver')
    }

    return(
        <>
            <Stack
                direction="column"
                justifyContent="center"
                alignItems="center"
                spacing={3}
                useFlexGap
            >
                <Paper elevation={3}
                       sx = {{marginTop: '15px', width:'70%'}}
                >
                    <br/>
                    <Button variant="outlined" sx = {{marginTop: '20px', width:'90%'}} onClick={AddDriverClick}>הוספת נהג</Button><br/>
                    <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={EditDriversClick}>עריכת נהגים</Button><br/>
                    <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={AddDriveClick}>הוספת נסיעה</Button><br/>
                    <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={EditDrivesClick}>עריכת נסיעות</Button><br/>
                    <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={SortClick}>יצירת שיבוץ</Button><br/>
                    <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={EditSortsClick}>עריכת שיבוצים</Button><br/>
                    <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={MyDrives}>הנסיעות שלי</Button><br/>

                    <br/><br/>
                </Paper>
            </Stack>
        </>
    );
}
export default managerPermissions(ManagerHome)