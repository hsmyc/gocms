export const $ = (selector: string) => document.querySelector(selector);
export const $$ = (selector: string) => document.querySelectorAll(selector);
export const on = (
  element: string | Element,
  event: string,
  callback: EventListenerOrEventListenerObject
): void => {
  if (typeof element === "string") {
    const queryElement = document.querySelector(element);
    if (queryElement) {
      queryElement.addEventListener(event, callback);
    } else {
      console.error(`Element not found for selector: "${element}"`);
    }
  } else if (element instanceof Element) {
    element.addEventListener(event, callback);
  } else {
    console.error('Invalid element provided to "on" function');
  }
};
export const off = (
  element: string | Element,
  event: string,
  callback: EventListenerOrEventListenerObject
): void => {
  if (typeof element === "string") {
    const queryElement = document.querySelector(element);
    if (queryElement) {
      queryElement.removeEventListener(event, callback);
    } else {
      console.error(`Element not found for selector: "${element}"`);
    }
  } else if (element instanceof Element) {
    element.removeEventListener(event, callback);
  } else {
    console.error('Invalid element provided to "off" function');
  }
};
