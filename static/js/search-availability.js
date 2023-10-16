// main function that runs rest of javascript
var searchAvailabilityPage = function () {
    var form = document.getElementById("check-availability-form");
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
    searchAvailabilityPage();
});
