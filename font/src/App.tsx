import './App.css';
import * as React from 'react';
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import LocalShippingIcon from "@mui/icons-material/LocalShipping";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import { LogIn } from "./LogIn/LogIn";
import { BrowserRouter, Navigate, Route, Routes, useNavigate } from "react-router-dom";
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
  const { nameOfPage, level, setLevel, setUserId } = useGlobalState();
  const navigate = useNavigate();

  const [anchorEl, setAnchorEl] = React.useState<HTMLElement | null>(null);
  const open = Boolean(anchorEl);
  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const goHome = () => {
    handleClose();
    if (level === 1 || level === 2) {
      navigate('/systemAdministrator');
    } else if (level === 3 || level === 4) {
      navigate('/manager');
    } else if (level === 5) {
      navigate('/driver');
    } else {
      navigate('/LogIn');
    }
  };

  const goToSort = () => {
    handleClose();
    navigate('/manager/Sort');
  };

  const logOut = () => {
    handleClose();
    setLevel(0);
    setUserId(-1);
    navigate('/LogIn');
  };

  return (
      <Box
          sx={{
            display: 'flex',
            flexDirection: 'column',
            minHeight: '100vh',
            width: '100%',
            bgcolor: 'background.default',
          }}
      >
        <AppBar position="static" color="primary">
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
              <MenuItem onClick={goHome}>דף הבית</MenuItem>
              <MenuItem onClick={goToSort}>שיבוץ נסיעה</MenuItem>
              <MenuItem onClick={logOut}>התנתקות</MenuItem>
            </Menu>
            <LocalShippingIcon sx={{ mr: 1.5 }} />
            <Typography variant="h6" component="div" sx={{ flexGrow: 1, fontWeight: 600 }}>
              {nameOfPage}
            </Typography>
          </Toolbar>
        </AppBar>

        <Box
            component="main"
            sx={{
              flexGrow: 1,
              width: '100%',
              display: 'flex',
              justifyContent: 'center',
              px: 2,
              py: { xs: 3, sm: 5 },
            }}
        >
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
        </Box>

      </Box>
  );
}

function App() {
  return (
      <GlobalStateProvider>
        <BrowserRouter>
          <AppContent />
        </BrowserRouter>
      </GlobalStateProvider>
  );
}

export default App;
