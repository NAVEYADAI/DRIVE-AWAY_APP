import * as React from 'react';
import Switch from '@mui/material/Switch';
import {LogIn} from "./LogIn";
import {SingUp} from "./SingUp";


export function MainLogIn() {

    const [auth, setAuth] = React.useState(true);
    const handleChange = (event) => {
        setAuth(event.target.checked);
    };
    return(<div>
        <Switch checked={auth} onChange={handleChange} aria-label="login switch"/>
        {auth ? <SingUp/> : <LogIn/>}
    </div>);
}