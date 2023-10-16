// page function that runs rest of javascript
const generalsPage = () => {
    console.log("Hello from generals page");
    let form :HTMLElement = document.getElementById("form-modal-general");
    let formData :FormData = new FormData(form as HTMLFormElement);
    formData.append("csrf_token", "{{.CSRFToken}}");

    fetch("/search-availability-json", {
        method: "post",
        body: formData,
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            console.log(data.ok);
            console.log(data.message);
        });
};

// runs page functions after load
window.addEventListener("load", () => {
    generalsPage();
});