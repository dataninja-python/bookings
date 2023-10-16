// main function that runs rest of javascript
const searchAvailabilityPage = () => {

    let form :HTMLElement = document.getElementById("check-availability-form");
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
    searchAvailabilityPage();
});