import { $, $$, on } from "../utils/aliases.js";
import statemanager, { ComponentStateManager } from "atsmanager";
import { Field, SingleType } from "../types.ts";
import postSingleType from "../api/postsingletype.ts";

export default function SingleTypeField() {
  const manager = new statemanager();
  manager.addComponentState([]);

  const state: ComponentStateManager<Field> =
    manager.getStateManager("component");

  const subscribe = state.subscribe;
  const sfields = state.getState();

  const apiID = $("#single_type_id") as HTMLInputElement;
  const name = $("#single_type_name") as HTMLInputElement;
  const item_save_button = $("#single_type_item_save_button");
  const closeButton = $(".modal__close");
  const save_button = $("#single_type_save_button");

  function Fields() {
    return `<p>Selected Fields</p>
      ${sfields()
        ?.map((field) => {
          switch (field.name) {
            case "Text":
              return `<div class='single_type_field'>
              <label>${field.name}</label>
              </div>`;
            case "Boolean":
              return `<div>${field.name}</div>`;
            case "RichText":
              return `<div>${field.name}</div>`;
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
        state.setState([
          ...(sfields() || []),
          { name: target.value, type: target.dataset.type },
        ]);
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

  on(item_save_button, "click", () => {
    $$(".single_type_field").forEach((field) => {
      (field as HTMLInputElement).checked = false;
    });
  });

  on(save_button, "click", async () => {
    const data = {
      apiId: apiID?.value,
      name: name?.value,
      fields: sfields(),
    };
    const res = await postSingleType(data as SingleType);
    console.log(res);
  });
}
