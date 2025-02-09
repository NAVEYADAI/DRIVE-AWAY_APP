import Button from '@mui/material/Button';
import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import Paper from '@mui/material/Paper';
import Stack from "@mui/material/Stack";
import {useNavigate} from "react-router-dom";
import {useGlobalState} from "../../GlobalState";

function SystemAdministrator() {
    const nav = useNavigate()
    const CompanyRegistrationClick = () => {
        nav('/systemAdministrator/CompanyRegistration')
    }
    const AddPositionClick = () => {
        nav('/systemAdministrator/AddPosition')
    }
    const NumberOfInlaysClick = () => {
        nav('/systemAdministrator/NumberOfInlays')
    }
    const MyCompany = () => {
        nav('/manager')
    }
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("מנהל חברה")
    return (<>
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
            <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={CompanyRegistrationClick}>רישום חברה</Button><br/>
            <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={AddPositionClick}>רישום עמדה בחרבה קיימת</Button><br/>
            <Button variant="outlined" sx = {{marginTop: '20px', width:'60%'}} onClick={NumberOfInlaysClick}>דוח מספר שיבוצים</Button><br/>
            <Button variant="outlined" sx = {{marginTop: '20px', width:'60%', marginBlockEnd:'20px'}} onClick={MyCompany}>חברה שלי</Button><br/><br/>
        </Paper>
        </Stack>
        <br/>

    </>);
}

export default systemAdministratorPermissions(SystemAdministrator);