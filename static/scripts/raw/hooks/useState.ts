export default function useState<T>(initialState: T): useStateReturn<T> {
  let currentState = initialState;
  const listeners = new Set<Function>();
  const getState = () => currentState;
  const setState = (newValue: SetStateAction<T>) => {
    currentState =
      newValue instanceof Function ? newValue(currentState) : newValue;
    listeners.forEach((listener) => listener());
  };

  const subscribe = (callback: Function) => {
    listeners.add(callback);
    return () => {
      listeners.delete(callback);
    };
  };

  return [getState, setState, subscribe];
}
