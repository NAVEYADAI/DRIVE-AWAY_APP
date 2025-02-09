import React, { createContext, useContext, useState, useEffect } from 'react';

const GlobalStateContext = createContext();

export const GlobalStateProvider = ({ children }) => {
    const [level, setLevel] = useState(() => {
        const savedLevel = localStorage.getItem('level');
        return savedLevel !== null ? JSON.parse(savedLevel) : 0;
    });

    const [userId, setUserId] = useState(() => {
        const savedUserId = localStorage.getItem('userId');
        return savedUserId !== null ? JSON.parse(savedUserId) : -1;
    });

    const [nameOfPage, setNameOfPage] = useState(() => {
        const savedNameOfPage = localStorage.getItem('nameOfPage');
        return savedNameOfPage !== null ? JSON.parse(savedNameOfPage) : "עמוד ראשי";
    });


    useEffect(() => {
        localStorage.setItem('level', JSON.stringify(level));
    }, [level]);

    useEffect(() => {
        localStorage.setItem('userId', JSON.stringify(userId));
    }, [userId]);

    useEffect(() => {
        localStorage.setItem('nameOfPage', JSON.stringify(nameOfPage));
    }, [nameOfPage]);


    return (
        <GlobalStateContext.Provider value={{ level, setLevel, userId, setUserId, nameOfPage,  setNameOfPage}}>
            {children}
        </GlobalStateContext.Provider>
    );
};

export const useGlobalState = () => {
    const context = useContext(GlobalStateContext);
    if (context === undefined) {
        throw new Error('useGlobalState must be used within a GlobalStateProvider');
    }
    return context;
};
