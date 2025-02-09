import React from "react";
import { Navigate } from "react-router-dom";
import { useGlobalState } from "../../GlobalState";

export function driverPermissions(Comp) {
    return function WrappedComponent(props) {
        const { level } = useGlobalState();

        if (level > 0 && level < 6) {
            return <Comp {...props} />;
        }

        return <Navigate to="/LogIn" />;
    }
}
