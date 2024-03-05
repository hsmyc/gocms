type SetStateAction<T> = T | ((prevState: T) => T);
type useStateReturn<T> = [
  getState: () => T,
  (action: SetStateAction<T>) => void,
  subscribe: (callback: (value: T) => void) => void
];

declare var $: (selector: string) => HTMLElement | null;
declare var $$: (selector: string) => NodeListOf<HTMLElement>;
declare var on: (
  element: string | Element | Document,
  event: string,
  callback: EventListenerOrEventListenerObject
) => void;
declare var off: (
  element: string | Element | Document,
  event: string,
  callback: EventListenerOrEventListenerObject
) => void;
