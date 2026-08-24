import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import AssessmentIcon from "@mui/icons-material/Assessment";
import * as React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import {useGlobalState} from "../../GlobalState";

interface CompanySortCount {
    nameOfCompany: string;
    numOfSort: string;
}

function GetValueOnCompany(): CompanySortCount[] {
    var listOfCompany = [{nameOfCompany:"name", numOfSort: "4" },{nameOfCompany:"yadai", numOfSort: "3" } ]
    return listOfCompany
}
function NumberOfInlays (){
    var listOfCompany = GetValueOnCompany()
    const { setNameOfPage } = useGlobalState();
    React.useEffect(() => {
        setNameOfPage("קבלת מספר השיבוצים שבצעה כל חברה");
    }, [setNameOfPage]);
    return(
        <Stack direction="column" alignItems="center" sx={{ width: '100%' }}>
            <Paper elevation={3} sx={{ mt: { xs: 2, sm: 6 }, width: '100%', maxWidth: 480, p: { xs: 3, sm: 4 }, borderRadius: 3 }}>
                <Stack direction="column" spacing={1} alignItems="center" sx={{ mb: 3 }}>
                    <AssessmentIcon color="primary" sx={{ fontSize: 36 }} />
                    <Typography variant="h6" fontWeight={600}>דוח מספר שיבוצים</Typography>
                </Stack>
                <List sx={{ width: '100%' }}>
                    {listOfCompany.map((item, index)=>(
                        <React.Fragment key={item.nameOfCompany}>
                            {index > 0 && <Divider component="li" />}
                            <ListItem>
                                <ListItemText primary={item.nameOfCompany} />
                                <ListItemText secondary={item.numOfSort} sx={{ textAlign: 'left' }} />
                            </ListItem>
                        </React.Fragment>
                    ))}
                </List>
            </Paper>
        </Stack>
    );
}
export default systemAdministratorPermissions(NumberOfInlays);
