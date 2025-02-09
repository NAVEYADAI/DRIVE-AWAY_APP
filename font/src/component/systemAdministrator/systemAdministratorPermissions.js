import React from "react";
import { Navigate } from "react-router-dom";
import { useGlobalState } from "../../GlobalState";

export function systemAdministratorPermissions(Comp) {
    return function WrappedComponent(props) {
        const { level } = useGlobalState();

        if (level === 1 || level === 2) {
            return <Comp {...props} />;
        }

        return <Navigate to="/LogIn" />;
    }
}
