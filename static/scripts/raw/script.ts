import Bold from "./functions/bold.js";
import { modalCloseEvent, modalOpenEvent } from "./functions/modal.js";
import SingleTypeField from "./singletype/singletype.js";

document.addEventListener("DOMContentLoaded", () => {
  Bold();
  modalCloseEvent();
  modalOpenEvent();
  SingleTypeField();
});
