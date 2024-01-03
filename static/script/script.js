function updateContent() {
    var content = document.querySelector("[contenteditable]")
        .innerText;
    document.getElementById("hiddenContent").value =
        content;
}
