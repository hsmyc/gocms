import { $, $$, on } from "../utils/aliases";

export function modalOpenFunction(modalId: string | null) {
  const modal = $(`#${modalId}`);

  if (!modal) {
    return;
  }
  modal.classList.add("modal--active");
}

export function modalCloseFunction(modalId: string | undefined) {
  const modal = $(`#${modalId}`);

  if (!modal) {
    return;
  }
  modal.classList.remove("modal--active");
}

export function attachModalEvents() {
  const modalOpens = $$("[data-modal-target]");
  const modalCloses = $$("[data-modal-close]");

  modalOpens.forEach((element: Element) => {
    on(element, "click", () => {
      if (element instanceof HTMLInputElement && !element.checked) {
        return;
      }
      const modalId = element.getAttribute("data-modal-target");
      modalOpenFunction(modalId);
    });
  });

  modalCloses.forEach((button) => {
    on(button, "click", () => {
      const modalId = button.closest(".modal")?.id;

      modalCloseFunction(modalId);
    });
  });
}
