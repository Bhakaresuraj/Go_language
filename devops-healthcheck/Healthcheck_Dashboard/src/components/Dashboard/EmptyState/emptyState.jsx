import "./emptyState.css";

export default function EmptyState({ openModal }) {
    return (
        <div className="empty-state-container">
            <div className="empty-card">
                <div className="empty-icon">
                    <i className="bi bi-cloud-slash"></i>
                </div>
                <h2>No Services Being Monitored</h2>
                <p>
                    Your dashboard is currently empty.
                    <br />
                    Register your first website or API endpoint
                    to start monitoring uptime and health.
                </p>
                <button
                    className="empty-btn"
                    onClick={openModal}
                >
                    <i className="bi bi-plus-circle me-2"></i>
                    Add First Service
                </button>
                <hr className="empty-divider" />
                <div className="quick-start">
                    <h5>
                        <i className="bi bi-lightning-charge-fill me-2"></i>
                        Quick Start
                    </h5>
                    <div className="steps">
                        <div className="step completed">
                            <i className="bi bi-check-circle-fill"></i>
                            Dashboard Ready
                        </div>
                        <div className="step">
                            <i className="bi bi-circle"></i>
                            Add First Service
                        </div>
                        <div className="step">
                            <i className="bi bi-circle"></i>
                            Start Monitoring
                        </div>

                        <div className="step">
                            <i className="bi bi-circle"></i>
                            View Analytics
                        </div>
                    </div>

                </div>

            </div>

        </div>
    );
}