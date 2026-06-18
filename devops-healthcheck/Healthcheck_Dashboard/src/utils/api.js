// for fetching data from the api from the backend
const Base_URL = import.meta.env.VITE_BACKEND_BASE_URL || "http://localhost:8080";


async function GetAllServices() {
    const token = localStorage.getItem("auth_token");
    let response = await fetch(`${Base_URL}/services`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        }
    });
    let data = await response.json();
    console.log( data);
    return data.data;
}


async function AddNewService(formdata) {
    const token = localStorage.getItem("auth_token");
    let response = await fetch(`${Base_URL}/add`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
            Name: formdata.Name,
            URL: formdata.URL
        })
    });
    return await response.json();

}
async function DeleteService(Id) {
    const token = localStorage.getItem("auth_token");
    let response = await fetch(`${Base_URL}/delete`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
            Id: Id
        })
    });
    const data = await response.json();
    return data;
}
async function UpdateService(formdata) {
    const token = localStorage.getItem("auth_token");
    let response = await fetch(`${Base_URL}/update`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
            ...formdata
        })
    });
    const data = await response.json();
    // console.log(data);
    return data;
}
export { AddNewService, GetAllServices, DeleteService, UpdateService };