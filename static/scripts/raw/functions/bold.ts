import { $ } from "../utils/aliases";

export default function Bold() {
  const boldButton = $("#bold");
  boldButton &&
    on(boldButton, "click", () => {
      document.execCommand("bold", false, null);
    });

  console.log("bold function");
}
