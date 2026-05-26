import { useState } from "react";
import './modal.css'
import { AddNewService } from "../../utils/api";
export default function AddServiceModal({ show, onClose, refreshServices }) {
    const [formData, setFormData] = useState({
        Name: "",
        URL: "",
        UserId: ""
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
                        type="number"
                        name="UserId"
                        placeholder="User ID"
                        value={formData.UserId}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="text"
                        name="Name"
                        placeholder="Service Name"
                        value={formData.Name}
                        onChange={handleChange}
                        required
                    />
                    <input
                        type="text"
                        name="URL"
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