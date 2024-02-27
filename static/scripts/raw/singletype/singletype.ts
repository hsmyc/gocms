import useState from "../hooks/setState.js";
import { modalCloseFunction } from "../functions/modal.js";
import { $, $$, on } from "../utils/aliases.js";
export default function SingleTypeField() {
  const [selectedFields, setSelectedFields, subscribe] = useState([]);
  const button = $("#single_type_save_button");
  const closeButton = $(".modal__close");

  function Fields() {
    return `<p>Selected Fields</p>
        ${selectedFields().map((field) => {
          return `<div>${field.name}</div>`;
        })}`;
  }
  function updateFields() {
    $("#single_type_fields").innerHTML = Fields();
  }

  subscribe(updateFields);

  $$(".single_type_field").forEach((field) => {
    on(field, "click", (e) => {
      const target = e.target as HTMLInputElement;
      if (target.checked) {
        setSelectedFields([...selectedFields(), { name: target.value }]);
      } else {
        setSelectedFields(
          selectedFields().filter((field) => field.name !== target.value)
        );
      }
    });
  });

  on(closeButton, "click", () => {
    $$(".single_type_field").forEach((field: HTMLInputElement) => {
      field.checked = false;
    });
    setSelectedFields([]);
  });

  on(button, "click", () => {
    $$(".single_type_field").forEach((field: HTMLInputElement) => {
      field.checked = false;
    });
    const data = {
      fields: selectedFields(),
    };
    console.log(data);
    modalCloseFunction();
  });
}
