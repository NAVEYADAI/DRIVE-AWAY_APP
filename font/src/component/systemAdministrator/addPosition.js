import {systemAdministratorPermissions} from "./systemAdministratorPermissions";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import * as React from 'react';
import {AddUser} from "../AddUser";
import Button from '@mui/material/Button';
import SendIcon from '@mui/icons-material/Send';
import {useGlobalState} from "../../GlobalState";


function GetAllCompany (){
    // קבלת כל השמות חברות הרשומות
    const listCompany = [{name:"nave"}, {name:"yadai"} ]
    // קבלה מהapi את רשימת כל החברות
    return listCompany
}
function AddPosition (){
    var listCompany = GetAllCompany()
    const [CompanyName, setCompanyName] = React.useState('');

    const handleChange = (event) => {
        setCompanyName(event.target.value);
    };
    const { setNameOfPage } = useGlobalState();
    setNameOfPage("הוספת עמדה בחברה קיימת")

    // קביעת משתנים עבור נתוני המשתמש
    const [userData, setUserData] = React.useState({
        FNameValue: '',
        LNameValue: '',
        IDValue: '',
        UserNameValue: '',
        PasswordValue: '',
        Password2Value: '',
        EmailValue: '',
        PhoneNumberValue: ''
    });
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
            רישום עמדה בחרבה קיימת
                    <br/>

                    <FormControl    sx = {{marginTop: '15px', width:'75%'}}>
                        <InputLabel id="CampnyNam">שם החברה</InputLabel>
                        <Select
                            labelId="CampnyName"
                            id="demo-simple-select"
                            value={CompanyName}
                            label="שם החברה"
                            onChange={handleChange}
                        >
                            {
                                listCompany.map((item)=>{
                                    return(
                                        <MenuItem value={item.name}>{item.name}</MenuItem>
                                    );
                                })
                            }

                    </Select>
                    </FormControl>
                    <AddUser userData={userData} setUserData={setUserData} />
                    <Button variant="contained" endIcon={<SendIcon />}>
                        Send
                    </Button>
                </Paper>
            </Stack>
        </>
    );
}
export default systemAdministratorPermissions(AddPosition);