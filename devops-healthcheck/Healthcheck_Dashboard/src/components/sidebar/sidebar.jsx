
import './sidebar.css'

function Sidebar({ openModal, goToTop, goToServices }) {
    return (
        <div
            className="d-flex flex-column sidebar flex-shrink-0 p-3 text-bg-dark"

        >
            <h3 onClick={goToTop} className="mb-4">HealthCheck</h3>
            <ul className="nav nav-pills side-component-li flex-column mb-auto">
                <div onClick={goToTop} className="nav-item side-component mb-2">
                    <a href="#" className="nav-link active">
                        <i className="bi bi-speedometer2 me-2"></i> Dashboard
                    </a>
                </div>
                <li onClick={goToServices} className="mb-2">
                    <a href="#" className="nav-link text-white">
                        <i className="bi bi-link-45deg"></i>&nbsp;&nbsp; All Services
                    </a>
                </li>
                <li onClick={openModal} className="mb-2" >
                    <a href="#" className="nav-link text-white">
                        <i className="bi bi-plus-circle-dotted"></i> &nbsp;&nbsp;Add Service
                    </a>
                </li>

                <li className="mb-2">
                    <a href="#" className="nav-link text-white">
                        <i className="bi bi-pie-chart-fill"></i>&nbsp;&nbsp;&nbsp;Analysis
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link text-white">
                        <i className="bi bi-gear-wide"></i>&nbsp;&nbsp;&nbsp;Settings
                    </a>
                </li>

            </ul>
            {/* <hr /> */}
            <div>
                &nbsp;&nbsp; <strong><i className="bi bi-box-arrow-right"></i>&nbsp;&nbsp;LogOut</strong>
            </div>
        </div>

    );
}

export default Sidebar;