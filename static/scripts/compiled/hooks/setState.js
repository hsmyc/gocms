function useState(initialState) {
    let currentState = initialState;
    const setState = (newValue) => {
        currentState =
            newValue instanceof Function ? newValue(currentState) : newValue;
    };
    return [currentState, setState];
}
