import Stack from '@mui/material/Stack';
import Paper from '@mui/material/Paper';
import TextField from '@mui/material/TextField';
import React from "react";
import SendIcon from '@mui/icons-material/Send';
import Button from '@mui/material/Button';
import Chip from "@mui/material/Chip";

export function SingUp() {

    const [UserNameValue, setUserNameValue] = React.useState();
    const [PasswordValue, setPasswordValue] = React.useState();
    const [ReturnPasswordValue, setReturnPasswordValue] = React.useState();
    const [FNameValue, setFNameValue] = React.useState();
    const [LNameValue, setLNameValue] = React.useState();
    const [PhoneValue, setPhoneValue] = React.useState();
    const [EmailValue, setEmailValue] = React.useState();

    return(<div>
        <br/><br/>
        <Stack

            direction="column"
            justifyContent="center"
            alignItems="center"

            spacing={3}
            useFlexGap
        >
           <Paper
               sx = {{marginTop: '15px', width:'75%'}}
           >
               <Chip label="Sing Up" variant="outlined" icon=<SendIcon/>/>
               <br/><br/>
               <TextField
                   id="UserId"
                   label="User Name"
                   value={UserNameValue}
                   onChange={(event) => {
                       setUserNameValue(event.target.value);
                   }}
                   fullWidth
               /><br/><br/>
                <TextField
                   id="PasswordId"
                   label="Password"
                   value={PasswordValue}
                   onChange={(event) => {
                       setPasswordValue(event.target.value);
                   }}
                   fullWidth
               /><br/><br/>
                <TextField
                   id="ReturnPasswordId"
                   label="ReturnPassword"
                   value={ReturnPasswordValue}
                   onChange={(event) => {
                       setReturnPasswordValue(event.target.value);
                   }}
                   fullWidth
               /><br/><br/>
                <TextField
                   id="FNameId"
                   label="First Name"
                   value={FNameValue}
                   onChange={(event) => {
                       setFNameValue(event.target.value);
                   }}
                   fullWidth
               /><br/><br/>
                <TextField
                   id="LNameId"
                   label="Last Name"
                   value={LNameValue}
                   onChange={(event) => {
                       setLNameValue(event.target.value);
                   }}
                   fullWidth
               /><br/><br/>
                <TextField
                   id="PhoneId"
                   label="Phone"
                   value={PhoneValue}
                   onChange={(event) => {
                       setPhoneValue(event.target.value);
                   }}
                   fullWidth
               /><br/><br/>
                <TextField
                   id="EmailId"
                   label="Email"
                   value={EmailValue}
                   onChange={(event) => {
                       setEmailValue(event.target.value);
                   }}
                   fullWidth
               /><br/><br/><br/>
               <Button variant="contained" endIcon={<SendIcon />}>
                   Send
               </Button>
               <br/><br/><br/>
           </Paper>
        </Stack>
    </div>);
}