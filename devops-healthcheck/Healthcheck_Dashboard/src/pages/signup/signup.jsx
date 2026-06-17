import "./signup.css";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
export default function Signup() {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        username: "",
        email: "",
        password: "",
        confirmPassword: ""

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
        /* PASSWORD CHECK */
        if (
            formData.password !==
            formData.confirmPassword
        ) {
            alert("Passwords do not match");
            return;
        }
        try {

            const response = await fetch(
                "http://localhost:8080/register",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json"
                    },
                    body: JSON.stringify({
                        username: formData.username,
                        email: formData.email,
                        password: formData.password
                    })
                }
            );
            const data = await response.json();
            // console.log(data);
            if (data.Success) {
                alert(data.Message);
                /* REDIRECT */
                navigate("/login");
            } else {
                alert(data.Message)
                navigate("/signup")
            }


        } catch (error) {
            console.log(error);
        }
    }
    return (
        <div className="signup-container">
            <div className="signup-card">
                <h1>
                    Create Account
                </h1>
                <p>
                    Start monitoring your infrastructure
                </p>
                {/* FORM */}
                <form
                    onSubmit={handleSubmit}
                    className="signup-form"
                >
                    {/* NAME */}
                    <div className="input-group">
                        <label htmlFor="username">
                            Name
                        </label>
                        <input
                            type="text"
                            id="username"
                            name="username"
                            placeholder="Enter username"
                            value={formData.username}
                            onChange={handleChange}
                            required
                        />
                    </div>

                    {/* EMAIL */}
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
                    {/* PASSWORD */}
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
                    {/* CONFIRM PASSWORD */}
                    <div className="input-group">
                        <label htmlFor="confirmPassword">
                            Confirm Password
                        </label>
                        <input
                            type="password"
                            id="confirmPassword"
                            name="confirmPassword"
                            placeholder="Confirm Password"
                            value={formData.confirmPassword}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    {/* BUTTON */}
                    <button type="submit">
                        Signup
                    </button>
                </form>
                {/* FOOTER */}
                <div className="signup-footer">
                    <p>
                        Already have an account?
                    </p>
                    <span
                        onClick={() => navigate("/login")}
                    >
                        Login
                    </span>
                </div>
            </div>
        </div>
    );
}