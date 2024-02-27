import { $, on } from "../utils/aliases";

const modal = $(".modal");
const modalClose = $(".modal__close");
const modalOpen = $(".modal__open");

export function modalOpenFunction() {
  modal.classList.add("modal--active");
}

export function modalCloseFunction() {
  modal.classList.remove("modal--active");
}

export function modalOpenEvent() {
  on(modalOpen, "click", modalOpenFunction);
}

export function modalCloseEvent() {
  on(modalClose, "click", modalCloseFunction);
}
