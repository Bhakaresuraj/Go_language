
import "./dropdown.css"
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../context/context";
export default function ProfileDropdown() {
    const navigate = useNavigate();
    const { logout } = useAuth();
    function handleLogout() {
        navigate("/", { replace: true });

        setTimeout(() => {
            logout();
        }, 100);
    }
    // const user = JSON.parse(localStorage.getItem("user")) || {};
    const { user } = useAuth();
    return (
        <>
            <div className="profile-section dropdown">

                <button
                    className="profile-btn dropdown-toggle"
                    data-bs-toggle="dropdown"
                >

                    <div className="avatar">

                        {user?.username?.charAt(0).toUpperCase()}

                    </div>

                    <div className="profile-info">

                        <span className="username">
                            {user?.username}
                        </span>

                        <span className="role">
                            Administrator
                        </span>

                    </div>

                </button>

                <ul className="dropdown-menu dropdown-menu-end">

                    <li className="dropdown-header">

                        <strong>{user?.username}</strong>

                        <br />

                        <small>{user?.email}</small>

                    </li>

                    <li><hr className="dropdown-divider" /></li>

                    <li>

                        <button className="dropdown-item">

                            <i className="bi bi-person me-2"></i>

                            My Profile

                        </button>

                    </li>

                    <li>

                        <button className="dropdown-item">

                            <i className="bi bi-key me-2"></i>

                            Change Password

                        </button>

                    </li>

                    <li>

                        <button
                            className="dropdown-item text-danger"
                            onClick={handleLogout}
                        >

                            <i className="bi bi-box-arrow-right me-2"></i>

                            Logout

                        </button>

                    </li>

                </ul>

            </div>



        </>
    )
}