import './App.css';
import * as React from 'react';
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import { LogIn } from "./LogIn/LogIn";
import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import SystemAdministrator from "./component/systemAdministrator/mainHome";
import CompanyRegistration from "./component/systemAdministrator/CompanyRegistration";
import AddPosition from "./component/systemAdministrator/addPosition";
import ManagerHome from "./component/manager/managerHome";
import AddDriver from "./component/manager/AddDriver";
import EditDrivers from "./component/manager/EditDrivers";
import NumberOfInlays from "./component/systemAdministrator/NumberOfInlays";
import AddDrive from "./component/manager/AddDrive";
import EditDrives from "./component/manager/EditDrives";
import Sort from "./component/manager/Sort";
import EditSort from "./component/manager/EditSort";
import { GlobalStateProvider, useGlobalState } from "./GlobalState";
import DriverHome from "./component/drivers/DriverHome";

function AppContent() {
  const { nameOfPage } = useGlobalState();

  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const [LogInAnchorEl, setLogInAnchorEl] = React.useState(null);
  const LogInOpen = Boolean(LogInAnchorEl);
  const logInMenu = (event) => {
    setLogInAnchorEl(event.currentTarget);
  };
  const LogInHandleClose = () => {
    setLogInAnchorEl(null);
    handleClose();
  };

  return (
      <Box
          width={'100%'}
          height={'100%'}
          sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            // justifyContent: 'center',
            minHeight: '100vh', // ensures full height of viewport
            textAlign: 'center' // centers text content
          }}
      >
        <AppBar position="static">
          <Toolbar>
            <IconButton
                size="large"
                edge="start"
                color="inherit"
                aria-label="menu"
                sx={{ mr: 2 }}
                aria-controls={open ? 'basic-menu' : undefined}
                aria-haspopup="true"
                aria-expanded={open ? 'true' : undefined}
                onClick={handleClick}
            >
              <MenuIcon />
            </IconButton>
            <Menu
                id="basic-menu"
                anchorEl={anchorEl}
                open={open}
                onClose={handleClose}
                MenuListProps={{
                  'aria-labelledby': 'basic-button',
                }}
            >
              <MenuItem onClick={() => handleClose()}>דף הבית</MenuItem>
              <MenuItem onClick={handleClose}>שיבוץ נסיעה</MenuItem>
              <MenuItem
                  aria-controls={LogInOpen ? 'basic-menu' : undefined}
                  aria-haspopup="true"
                  aria-expanded={LogInOpen ? 'true' : undefined}
                  onClick={logInMenu}
              >
                Logout
              </MenuItem>
              <Menu
                  open={LogInOpen}
                  anchorEl={LogInAnchorEl}
                  onClose={LogInHandleClose}
                  MenuListProps={{
                    'aria-labelledby': 'basic-button',
                  }}
              >
                <MenuItem onClick={() => LogInHandleClose()}>Log In</MenuItem>
                <MenuItem onClick={() => LogInHandleClose()}>Sing Up</MenuItem>
              </Menu>
            </Menu>
            <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
              {nameOfPage}
            </Typography>
          </Toolbar>
        </AppBar>

        <Box
            width={'100%'}
            height={'100%'}>
        <BrowserRouter fullWidth>
          <Routes>
            <Route path="/LogIn" element={<LogIn />} />
            <Route path="/" element={<Navigate to="/LogIn" />} />
            <Route path="/systemAdministrator" element={<SystemAdministrator />} />
            <Route path="/systemAdministrator/CompanyRegistration" element={<CompanyRegistration />} />
            <Route path="/systemAdministrator/AddPosition" element={<AddPosition />} />
            <Route path="/systemAdministrator/NumberOfInlays" element={<NumberOfInlays />} />
            <Route path="/manager" element={<ManagerHome />} />
            <Route path="/manager/AddDriver" element={<AddDriver />} />
            <Route path="/manager/EditDrivers" element={<EditDrivers />} />
            <Route path="/manager/AddDrive" element={<AddDrive />} />
            <Route path="/manager/EditDrives" element={<EditDrives />} />
            <Route path="/manager/Sort" element={<Sort />} />
            <Route path="/manager/EditSort" element={<EditSort />} />
            <Route path="/driver" element={<DriverHome />} />
          </Routes>
        </BrowserRouter>
        </Box>

      </Box>
  );
}

function App() {
  return (
      <GlobalStateProvider>
        <AppContent />
      </GlobalStateProvider>
  );
}

export default App;
