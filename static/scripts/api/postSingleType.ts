import { SingleType } from "../types";
export default async function postSingleType(body: SingleType) {
  const response = await fetch("/api/singletype/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });

  return response.json();
}
