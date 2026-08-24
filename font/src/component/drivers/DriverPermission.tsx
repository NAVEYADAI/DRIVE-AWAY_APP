import React from "react";
import { Navigate } from "react-router-dom";
import { useGlobalState } from "../../GlobalState";

export function driverPermissions<P extends object>(Comp: React.ComponentType<P>) {
    return function WrappedComponent(props: P) {
        const { level } = useGlobalState();

        if (level > 0 && level < 6) {
            return <Comp {...props} />;
        }

        return <Navigate to="/LogIn" />;
    }
}
