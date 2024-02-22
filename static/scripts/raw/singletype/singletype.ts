import useState from "../hooks/setState.js";

export default function SingleTypeField() {
  const [selectedFields, setSelectedFields, subscribe] = useState([]);
  const button = document.getElementById("single_type_save_button");
  function Fields() {
    return `
        <div
        <p>Selected Fields</p>
        ${selectedFields().map((field) => {
          return `<p>${field.name}</p>`;
        })}
        </div>
        `;
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

  button.addEventListener("click", () => {
    const data = {
      fields: selectedFields(),
    };
    console.log(data);
  });

  updateFields();
}
