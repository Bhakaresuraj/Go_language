import './table.css';
import { useState } from 'react';
import { DeleteService } from '../../../utils/api'
export default function ServiceTable({ setShowUpdate, setFormData, data }) {

    async function HandleDelete(Id) {
        const confirmDelete = window.confirm(
            "Are you sure you want to delete this service?"
        );
        if (!confirmDelete) {
            return;
        }
        let response = await DeleteService(Id);
        alert(response.Message);
    }

    const [search, setSearch] = useState("");
    const filteredServices = data.filter((service) =>
        service.Name.toLowerCase().includes(search.toLowerCase())
    );

    async function HandleUpdate(service) {
        // console.log("from table :", service);
        setFormData({
            ...service
        });
        setShowUpdate(true);
    }

    return (
        <div className="table-container">
            <div className="table-header">
                <h2 className="table-title">
                    📋 Monitored Services
                </h2>
                <div className="table-actions">
                    <input
                        type="text"
                        placeholder="Search services..."
                        value={search}
                        onChange={(e) => setSearch(e.target.value)}
                    />
                    <select>
                        <option>All Status</option>
                        <option>Healthy</option>
                        <option>Down</option>
                    </select>
                </div>
            </div>
            <table className="custom-table">
                <thead>
                    <tr>
                        <th>SERVICE NAME</th>
                        <th>TYPE</th>
                        <th>STATUS</th>
                        <th>RESPONSE</th>
                        <th>LAST CHECKED</th>
                        <th>ACTIONS</th>
                    </tr>
                </thead>
                <tbody>
                    {filteredServices.map((service) => (
                        <tr key={service.ID}>
                            <td>
                                <div className="service-name">
                                    {service.Name}
                                </div>
                                <div className="service-url">
                                    {service.URL}
                                </div>
                            </td>
                            <td>
                                <span className="type-badge">
                                    HTTP
                                </span>
                            </td>
                            <td>
                                <span className={`status ${service.Healthy
                                    ? "healthy"
                                    : "down"
                                    }`}>
                                    {service.Healthy
                                        ? "Healthy"
                                        : "Down"}
                                </span>
                            </td>
                            <td>
                                {
                                    service.Healthy ? (
                                        <span
                                            className={
                                                service.Response_time < 200
                                                    ? "response-good"
                                                    : service.Response_time < 500
                                                        ? "response-medium"
                                                        : "response-bad"
                                            }
                                        >
                                            {service.Response_time} ms
                                        </span>
                                    ) : (
                                        <span className="down timeout">
                                            Timeout
                                        </span>
                                    )
                                }
                            </td>
                            <td>
                                <div className="service-name">

                                    {service.Checked_at.slice(11, 19)}
                                </div>
                                <div className="service-url">

                                    {service.Checked_at.slice(0, 10)}
                                </div>
                            </td>
                            <td>
                                <div className="action-buttons">
                                    <button onClick={() => HandleUpdate(service)} className="action-btn">
                                        ✏
                                    </button>
                                    <button onClick={() => HandleDelete(service.ID)} className="action-btn">
                                        🗑
                                    </button>
                                </div>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div >
    );
}