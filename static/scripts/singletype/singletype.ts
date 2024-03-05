import { modalCloseFunction } from "../functions/modal.js";
import { $, $$, on } from "../utils/aliases.js";
import statemanager, { ComponentStateManager } from "atsmanager";

type Field = {
  name: string;
}[];
export default function SingleTypeField() {
  const manager = new statemanager();
  manager.addComponentState([]);
  const state: ComponentStateManager<Field> =
    manager.getStateManager("component");
  const subscribe = state.subscribe;
  const sfields = state.getState();

  const button = $("#single_type_save_button");
  const closeButton = $(".modal__close");

  function Fields() {
    return `<p>Selected Fields</p>
      ${sfields()
        ?.map((field) => {
          switch (field.name) {
            case "text":
              return `<div class='single_type_field'>
              <label>${field.name}</label>
              </div>`;
            case "bool":
              return `<div><label>${field.name}<input type="checkbox"/></label></div>`;
            case "richtext":
              return `<div contenteditable="true" style="border: 1px solid #ccc; min-height: 100px;">${field.name}</div>`;
            default:
              return `<div>Unsupported type</div>`;
          }
        })
        .join("")}`;
  }

  function updateFields() {
    const singleTypeFields = $("#single_type_fields");
    if (!singleTypeFields) {
      return;
    }
    singleTypeFields.innerHTML = Fields();
  }
  subscribe(updateFields);
  $$(".single_type_field").forEach((field) => {
    on(field, "click", (e) => {
      const target = e.target as HTMLInputElement;
      if (target.checked) {
        state.setState([...(sfields() || []), { name: target.value }]);
      } else {
        state.setState(
          (sfields() || []).filter((field) => field.name !== target.value)
        );
      }
    });
  });

  on(closeButton, "click", () => {
    $$(".single_type_field").forEach((field) => {
      (field as HTMLInputElement).checked = false;
    });
    state.setState([]);
  });

  on(button, "click", () => {
    $$(".single_type_field").forEach((field) => {
      (field as HTMLInputElement).checked = false;
    });
    const data = {
      fields: sfields()?.map((field) => field.name) || [],
    };

    modalCloseFunction();
  });
}
