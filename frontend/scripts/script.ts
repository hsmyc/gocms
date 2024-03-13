import { attachModalEvents } from "./components/modal.js";
import SingleTypeField from "./singletype/singletype.js";

document.addEventListener("DOMContentLoaded", () => {
  attachModalEvents();
  SingleTypeField();
});
