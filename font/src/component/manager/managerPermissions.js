import React from "react";
import { Navigate } from "react-router-dom";
import { useGlobalState } from "../../GlobalState";

export function managerPermissions(Comp) {
    return function WrappedComponent(props) {
        const { level } = useGlobalState();

        if (level > 0 && level < 5) {
            return <Comp {...props} />;
        }

        return <Navigate to="/LogIn" />;
    }
}
