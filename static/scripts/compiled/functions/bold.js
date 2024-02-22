export default function Bold() {
    const boldButton = document.getElementById("bold");
    boldButton &&
        boldButton.addEventListener("click", () => {
            document.execCommand("bold", false, null);
        });
    console.log("bold function");
}
