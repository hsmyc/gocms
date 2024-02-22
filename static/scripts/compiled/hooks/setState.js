export default function useState(initialState) {
    let currentState = initialState;
    const listeners = new Set();
    const getState = () => currentState;
    const setState = (newValue) => {
        currentState =
            newValue instanceof Function ? newValue(currentState) : newValue;
        listeners.forEach((listener) => listener());
    };
    const subscribe = (callback) => {
        listeners.add(callback);
        return () => {
            listeners.delete(callback);
        };
    };
    return [getState, setState, subscribe];
}
