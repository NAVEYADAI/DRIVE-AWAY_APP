import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';

interface GlobalStateContextValue {
    level: number;
    setLevel: React.Dispatch<React.SetStateAction<number>>;
    userId: number;
    setUserId: React.Dispatch<React.SetStateAction<number>>;
    nameOfPage: string;
    setNameOfPage: React.Dispatch<React.SetStateAction<string>>;
}

const GlobalStateContext = createContext<GlobalStateContextValue | undefined>(undefined);

export const GlobalStateProvider = ({ children }: { children: ReactNode }) => {
    const [level, setLevel] = useState<number>(() => {
        const savedLevel = localStorage.getItem('level');
        return savedLevel !== null ? JSON.parse(savedLevel) : 0;
    });

    const [userId, setUserId] = useState<number>(() => {
        const savedUserId = localStorage.getItem('userId');
        return savedUserId !== null ? JSON.parse(savedUserId) : -1;
    });

    const [nameOfPage, setNameOfPage] = useState<string>(() => {
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
        <GlobalStateContext.Provider value={{ level, setLevel, userId, setUserId, nameOfPage, setNameOfPage }}>
            {children}
        </GlobalStateContext.Provider>
    );
};

export const useGlobalState = (): GlobalStateContextValue => {
    const context = useContext(GlobalStateContext);
    if (context === undefined) {
        throw new Error('useGlobalState must be used within a GlobalStateProvider');
    }
    return context;
};
