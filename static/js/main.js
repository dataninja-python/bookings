window.addEventListener("load", function () {
    main();
});
// runs the main function in the file
var main = function () {
    // 👇️ const input: HTMLInputElement | null
    var input = document.getElementById("start_date");
    var value = input === null || input === void 0 ? void 0 : input.value;
    console.log(value); // 👉️ "Initial value"
    if (input != null) {
        console.log(input.value); // 👉️ "Initial value"
    }
    input === null || input === void 0 ? void 0 : input.addEventListener("input", function (event) {
        var target = event.target;
        console.log(target.value);
        displayToScreen(target.value);
    });
    var displayToScreen = function (val) {
        var outputDisplay = document.getElementById("values");
        outputDisplay.textContent = val;
    };
    console.log("Hello from the server.");
};
