type SetStateAction<T> = T | ((prevState: T) => T);
type useStateReturn<T> = [
  getState: () => T,
  (action: SetStateAction<T>) => void,
  subscribe: (callback: Function) => void
];
