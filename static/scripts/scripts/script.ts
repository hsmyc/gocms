function updateContent() {
  let content = (document.querySelector("[contenteditable]") as HTMLElement)
    .innerText;
  (document.getElementById("hiddenContent") as HTMLInputElement).value =
    content;
}
