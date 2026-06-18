import "./navbar.css";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import ProfileDropdown from "../../ProfileDropdown/profileDropdown"
import { useAuth } from "../../../context/context";
export default function Navbar() {
    const { isAuth } = useAuth();
    const navigate = useNavigate();
    function handleDashboard() {
        if (isAuth) {
            navigate("/dashboard");
        } else {
            navigate("/login", {
                state: { from: "/dashboard" }
            });
        }
    }
    return (
        <nav className="navbar">
            {/* Left */}
            <div className="navbar-logo">
                <h2>Health-check</h2>
            </div>

            {/* Center */}
            <ul className="navbar-links">

                <li>
                    <Link onClick={handleDashboard}
                        to="/dashboard"
                        className="dashboard-btn"
                    >
                        <i className="bi bi-speedometer2 me-2"></i>
                        Dashboard
                    </Link>
                </li>
                <li><a href="#home">Home</a></li>
                <li><a href="#services">Services</a></li>
                <li><a href="#about">About</a></li>
                <li><a href="#contact">Contact</a></li>

            </ul>

            {/* Right */}

            <div className="navbar-right">

                {isAuth ?
                    <div className="navbar-profile">
                        <ProfileDropdown />
                    </div>
                    :
                    (
                        <div className="navbar-buttons">
                            <Link to="/login">
                                <button className="login-btn">
                                    Login
                                </button>
                            </Link>
                            <Link to="/signup">
                                <button className="signup-btn">
                                    Signup
                                </button>

                            </Link>
                        </div>
                    )}
            </div>

        </nav>
    );
}