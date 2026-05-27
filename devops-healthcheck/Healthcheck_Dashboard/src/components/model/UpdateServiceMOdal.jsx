import { useState } from "react";
import './modal.css'
import { UpdateService } from "../../utils/api";
export default function UpdateServiceModal({ formData, setFormData, show, onClose, refreshServices }) {
    function handleChange(e) {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    }
    async function handleSubmit(e) {
        e.preventDefault();
        try {
            // console.log("before Updating",formData);
            const response = await UpdateService(formData);
            refreshServices();
            onClose();
        } catch (err) {
            console.log(err);
        }
    }
    // console.log("from data before update:", formData);

    if (!show) return null;
    return (
        <div className="modal-overlay">
            <div className="modal-box">
                <h2>Update Service</h2>
                <form onSubmit={handleSubmit}>
                    <label htmlFor="name">Update Service Name :</label>
                    <input
                        type="text"
                        id="name"
                        name="Name"
                        placeholder="Update Service Name"
                        value={formData.Name}
                        onChange={handleChange}
                        required
                    />
                    <label htmlFor="url">Update service URL :</label>
                    <input
                        id="url"
                        type="text"
                        name="URL"
                        placeholder="Update Service URL"
                        value={formData.URL}
                        onChange={handleChange}
                        required
                    />

                    <div className="modal-buttons">

                        <button
                            type="submit"
                            className="add-btn"
                        >
                            Update Service
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