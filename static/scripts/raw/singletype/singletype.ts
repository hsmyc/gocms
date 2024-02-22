import useState from "../hooks/setState.js";
import { modalCloseFunction } from "../functions/modal.js";
export default function SingleTypeField() {
  const [selectedFields, setSelectedFields, subscribe] = useState([]);
  const button = document.getElementById("single_type_save_button");
  const closeButton = document.querySelector(".modal__close");

  function Fields() {
    return `<p>Selected Fields</p>
        ${selectedFields().map((field) => {
          return `<div>${field.name}</div>`;
        })}`;
  }
  function updateFields() {
    document.getElementById("single_type_fields").innerHTML = Fields();
  }

  subscribe(updateFields);

  document.querySelectorAll(".single_type_field").forEach((field) => {
    field.addEventListener("click", (e) => {
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

  closeButton.addEventListener("click", () => {
    document
      .querySelectorAll(".single_type_field")
      .forEach((field: HTMLInputElement) => {
        field.checked = false;
      });
    setSelectedFields([]);
  });

  button.addEventListener("click", () => {
    document
      .querySelectorAll(".single_type_field")
      .forEach((field: HTMLInputElement) => {
        field.checked = false;
      });
    const data = {
      fields: selectedFields(),
    };
    console.log(data);
    modalCloseFunction();
  });
}
