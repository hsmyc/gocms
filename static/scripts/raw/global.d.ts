type SetStateAction<T> = T | ((prevState: T) => T);
type useStateReturn<T> = [T, (action: SetStateAction<T>) => void];
