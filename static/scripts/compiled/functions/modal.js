export default function Modal() {
    const modal = document.querySelector(".modal");
    const modalClose = document.querySelector(".modal__close");
    const modalOpen = document.querySelector(".modal__open");
    modalOpen.addEventListener("click", () => {
        modal.classList.add("modal--active");
    });
    modalClose.addEventListener("click", () => {
        modal.classList.remove("modal--active");
    });
}
