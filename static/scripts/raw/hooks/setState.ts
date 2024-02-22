function useState<T>(initialState: T): useStateReturn<T> {
  let currentState = initialState;
  const setState = (newValue: SetStateAction<T>) => {
    currentState =
      newValue instanceof Function ? newValue(currentState) : newValue;
  };

  return [currentState, setState];
}
