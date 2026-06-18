import { useState } from "react";
import './modal.css'
import { AddNewService, } from "../../../utils/api";
export default function AddServiceModal({ show, onClose, refreshServices }) {
    const [formData, setFormData] = useState({
        Name: "",
        URL: "",
    });
    function handleChange(e) {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    }
    async function handleSubmit(e) {
        e.preventDefault();
        try {
            const response = await AddNewService(formData);
            alert(response.Message);
            refreshServices();
            onClose();
        } catch (err) {
            console.log(err);
        }
    }
    if (!show) return null;
    return (
        <div className="modal-overlay">
            <div className="modal-box">
                <h2>Add New Service</h2>
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        name="Name"
                        id="name"
                        placeholder="Service Name"
                        value={formData.Name}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="text"
                        name="URL"
                        id="url"
                        placeholder="Service URL"
                        value={formData.URL}
                        onChange={handleChange}
                        required
                    />

                    <div className="modal-buttons">

                        <button
                            type="submit"
                            className="add-btn"
                        >
                            Add Service
                        </button>
                        <button
                            type="button"
                            className="cancel-btn"
                            onClick={onClose}
                        >
                            Cancel
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}