window.addEventListener("load", () => {
    main()
})

// runs the main function in the file
const main = () => {
    // 👇️ const input: HTMLInputElement | null
    const input = document.getElementById("start_date") as HTMLInputElement | null;

    const value = input?.value;
    console.log(value) // 👉️ "Initial value"

    if (input != null) {
        console.log(input.value); // 👉️ "Initial value"
    }

    input?.addEventListener("input", function (event) {
        const target = event.target as HTMLInputElement;
        console.log(target.value);
        displayToScreen(target.value);
    });

    const displayToScreen = (val) => {
        let outputDisplay = document.getElementById("values");
        outputDisplay.textContent = val;
    }

    console.log("Hello from the server.")
};
