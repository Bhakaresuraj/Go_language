import { createContext, useContext, useState } from "react";

const AuthContext = createContext();
export function AuthProvider({ children }) {
    const [user, setUser] = useState(
        JSON.parse(localStorage.getItem("user"))
    );
    const [isAuth, setIsAuth] = useState(
        !!localStorage.getItem("auth_token")
    );
    function login(user, token) {
        localStorage.setItem("auth_token", token);
        localStorage.setItem(
            "user",
            JSON.stringify(user)
        );
        setUser(user);
        setIsAuth(true);
    }
    function logout() {
        localStorage.removeItem("auth_token");
        localStorage.removeItem("user");
        setUser(null);
        setIsAuth(false);
    }
    return (
        <AuthContext.Provider
            value={{
                user,   
                isAuth,
                login,
                logout
            }}
        >
            {children}
        </AuthContext.Provider>
    );
}
export function useAuth() {
    return useContext(AuthContext);
}