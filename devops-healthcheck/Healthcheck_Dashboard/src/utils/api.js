// for fetching data from the api from the backend
const Base_URL = import.meta.env.VITE_BACKEND_BASE_URL || "http://localhost:8080";
async function GetAllServices() {
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
async function AddNewService(formdata) {
    let response = await fetch(`${Base_URL}/add`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            UserId: Number(formdata.UserId),
            Name: formdata.Name,
            URL: formdata.URL
        })
    });
    let data = await response.json();
    console.log(data)
    return data;
}
async function DeleteService(Id) {
    let response = await fetch(`${Base_URL}/delete`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            Id: Id
        })
    });
    // console.log(response)
    return;
}


async function UpdateService(formdata) {
    let response = await fetch(`${Base_URL}/update`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            ...formdata
        })
    });
    // let data = await response.json()
    // console.log("Update :", data)
    return;
}
export { AddNewService, GetAllServices, DeleteService, UpdateService };