import './navbar.css'

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
                        {/* <ul className="navbar-nav m-0 ms-4">
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
                        </ul> */}
                        {/* Right Profile */}
                        <ul className="navbar-nav ms-auto">
                            <li className="nav-item dropdown">
                                <a
                                    className="nav-link dropdown-toggle"
                                    href="#"
                                    id="navbarDropdown"
                                    role="button"
                                    data-bs-toggle="dropdown"
                                >
                                    Profile
                                </a>

                                <ul className="dropdown-menu dropdown-menu-end">

                                    <li>
                                        <a className="dropdown-item" href="#">
                                            Profile
                                        </a>
                                    </li>

                                    <li>
                                        <a className="dropdown-item" href="#">
                                            Settings
                                        </a>
                                    </li>

                                    <li><hr className="dropdown-divider" /></li>

                                    <li>
                                        <a className="dropdown-item text-danger" href="#">
                                            <i className="bi bi-box-arrow-right me-2"></i>
                                            Logout
                                        </a>
                                    </li>

                                </ul>

                            </li>

                        </ul>

                    </div>

                </div>


            {/* <hr /> */}
            </nav>
            {/* <hr /> */}
        </>
    )
}