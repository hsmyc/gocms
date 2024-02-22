function singleTypeField() {
    const [selectedFields, setSelectedFields] = useState([]);
    const button = document.getElementById("single_type_save_button");
    function Fields() {
        return `
        <div
        ${selectedFields.map((field) => {
            return `<p>${field}</p>`;
        })}
        }
        </div>
        `;
    }
    document.getElementById("single_type_fields").innerHTML = Fields();
    document.querySelectorAll(".field").forEach((field) => {
        field.addEventListener("click", (e) => {
            const target = e.target;
            if (target.checked) {
                setSelectedFields([
                    ...selectedFields,
                    { name: target.value, type: target.type },
                ]);
            }
            else {
                setSelectedFields(selectedFields.filter((field) => field.name !== target.value));
            }
        });
    });
    button.addEventListener("click", () => {
        const data = {
            fields: selectedFields,
        };
        fetch("/api/singletype/", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });
    });
}
