import './navbar.css'
import ProfileDropdown from "../../ProfileDropdown/profileDropdown"
export default function Navbar() {
    return (
        <>
            <nav className="navbar navbar-expand-lg navbar-dark ">
                <div className="container-fluid ">
                    {/* Logo */}
                    <a className="navbar-brand d-flex align-items-center" href="#">
                        <h1 className="m-0">
                            <i className="bi bi-speedometer2 me-2"></i>
                            Dashboard
                        </h1>
                    </a>
                    {/* Toggle Button */}
                    <button
                        className="navbar-toggler"
                        type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#navbarSupportedContent"
                    >
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    {/* Navbar Content */}
                    <div
                        className="collapse navbar-collapse"
                        id="navbarSupportedContent"
                    >
                        {/* Left Menu */}
                        <ul className="navbar-nav m-0 ms-4">
                            <li className="nav-item">
                                <a className="nav-link active" href="#">
                                    Home
                                </a>
                            </li>
                            <li className="nav-item">
                                <a className="nav-link" href="#">
                                    Link
                                </a>
                            </li>
                        </ul>

                        {/* Right Profile */}
                        <div className=" ms-auto">
                            <ProfileDropdown />
                        </div>
                    </div>
                </div>
                {/* <hr /> */}
            </nav>
            {/* <hr /> */}
        </>
    )
}