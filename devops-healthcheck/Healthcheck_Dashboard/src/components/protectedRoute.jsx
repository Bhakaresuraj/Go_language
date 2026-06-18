import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../context/context";

export default function ProtectedRoute({ children }) {
    const location = useLocation();
    const { isAuth } = useAuth();

    if (!isAuth) {

        return <Navigate to="/login" replace state={{
            from: location.pathname
        }} />;

    }

    return children;
}
