import "./login.css";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../context/context";
import { useLocation } from "react-router-dom";
export default function Login() {
    const { login } = useAuth();
    const location = useLocation();
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        email: "",
        password: ""
    });
    function handleChange(event) {
        const { name, value } = event.target;
        setFormData((prev) => ({
            ...prev,
            [name]: value
        }));
    }
    async function handleSubmit(event) {
        event.preventDefault();
        try {
            // console.log(formData);
            const response = await fetch(
                "http://localhost:8080/login",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify(formData)
                }
            );
            const data = await response.json();
            // console.log(data);
            /* STORE TOKEN */
            if (data.Success) {
                const redirectTo = location.state?.from || "/";
                login(data.data, data.token);
                navigate(redirectTo, { replace: true });
            } else {
                alert(data.Message)
            }
        } catch (error) {
            console.log(error);
        }
    }
    return (
        <div className="login-container">
            <div className="login-card">
                <h1>
                    Welcome Back
                </h1>
                <p>
                    Login to access your dashboard
                </p>
                {/* FORM */}
                <form
                    onSubmit={handleSubmit}
                    className="login-form"
                >
                    <div className="input-group">
                        <label htmlFor="email">
                            Email
                        </label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            placeholder="Enter Email"
                            value={formData.email}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    <div className="input-group">
                        <label htmlFor="password">
                            Password
                        </label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            placeholder="Enter Password"
                            value={formData.password}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    <button type="submit">
                        Login
                    </button>
                </form>
                {/* FOOTER */}
                <div className="login-footer">
                    <p>
                        Don't have an account?
                    </p>
                    <span onClick={() => navigate("/signup")}>
                        Signup
                    </span>
                </div>
            </div>
        </div>
    );
}
