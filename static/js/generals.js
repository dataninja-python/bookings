// page function that runs rest of javascript
var generalsPage = function () {
    console.log("Hello from generals page");
    var form = document.getElementById("form-modal-general");
    var formData = new FormData(form);
    formData.append("csrf_token", "{{.CSRFToken}}");
    fetch("/search-availability-json", {
        method: "post",
        body: formData,
    })
        .then(function (response) { return response.json(); })
        .then(function (data) {
        console.log(data);
        console.log(data.ok);
        console.log(data.message);
    });
};
// runs page functions after load
window.addEventListener("load", function () {
    generalsPage();
});
