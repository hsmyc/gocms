function updateContent() {
    let content = document.querySelector("[contenteditable]")
        .innerText;
    document.getElementById("hiddenContent").value =
        content;
}
