// for fetching data from the api from the backend


const Base_URL = import.meta.env.VITE_BACKEND_BASE_URL || "http://localhost:8080";
export default async function GetAllServices() {
    let response = await fetch(`${Base_URL}/runall`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            UserId: 1
        })
    });
    let data = await response.json();
    return data;

}
// GetAllServices();

