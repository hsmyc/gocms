const modal = document.querySelector(".modal");
const modalClose = document.querySelector(".modal__close");
const modalOpen = document.querySelector(".modal__open");

export function modalOpenFunction() {
  modal.classList.add("modal--active");
}

export function modalCloseFunction() {
  modal.classList.remove("modal--active");
}

export function modalOpenEvent() {
  modalOpen.addEventListener("click", modalOpenFunction);
}

export function modalCloseEvent() {
  modalClose.addEventListener("click", modalCloseFunction);
}
