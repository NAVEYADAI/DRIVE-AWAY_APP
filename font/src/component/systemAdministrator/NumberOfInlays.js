import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import * as React from 'react';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import {useGlobalState} from "../../GlobalState";


function GetValueOnCompany (){
    var listOfCompany = [{nameOfCompany:"name", numOfSort: "4" },{nameOfCompany:"yadai", numOfSort: "3" } ]
    return listOfCompany
}
function NumberOfInlays (){
    var listOfCompany = GetValueOnCompany()
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("קבלת מספר השיבוצים שבצעה כל חברה")
    return(
        <>
            דוח מספר שיבוצים
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
                    <List sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}>

                    {listOfCompany.map((item)=>{
                        return(
                            <>
                                <ListItem>
                                    <ListItemText primary={item.nameOfCompany} />
                                    <ListItemText  secondary={item.numOfSort} />

                                </ListItem>
                            </>
                        );
                    })}
                    </List>
                </Paper>
            </Stack>
        </>
    );
}
export default systemAdministratorPermissions(NumberOfInlays);