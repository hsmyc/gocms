import { $, on } from "../utils/aliases";

const modal = $(".modal");
const modalClose = $(".modal__close");
const modalOpen = $(".modal__open");

export function modalOpenFunction() {
  if (!modal) {
    return;
  }
  modal.classList.add("modal--active");
}

export function modalCloseFunction() {
  if (!modal) {
    return;
  }
  modal.classList.remove("modal--active");
}

export function modalOpenEvent() {
  if (!modalOpen) {
    return;
  }
  on(modalOpen, "click", modalOpenFunction);
}

export function modalCloseEvent() {
  if (!modalClose) {
    return;
  }
  on(modalClose, "click", modalCloseFunction);
}
