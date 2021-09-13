import React, { createContext, useCallback, useReducer } from 'react';
const initialState = {
    items: []
};
const actions = {
    Add: 'Add',
    Remove: 'Remove',
    Load: 'Load'
}
const reducer = (state, action) => {
    // eslint-disable-next-line default-case
    switch (action.type) {
        case actions.Add: {
            return { ...state, items: [...state.items, action.item] }
        }
        case actions.Remove: {
            const index = state.item.indexof(action.item);
            return {
                ...state.items.slice(0, index),
                ...state.items.slice(index + 1)
            }
        }
        case actions.Load: {
            return { ...state,items: action.items }
        }
    }
    return state;
}

const CartContext = createContext(initialState);
export const CartContextProvider = ({ children }) => {
    const [cartState, cartDispatch] = useReducer(reducer, initialState);

    const add= useCallback(
        (elem) => {
            
            cartDispatch({
                type: actions.Add,
                item:elem

            });
            localStorage.setItem("cart",cartState.items)
        },
        [cartDispatch,cartState.items],
    ) ;
    const remove= useCallback(
        (elem) => {
            cartDispatch({
                type: actions.Remove,
                elem
            })
            localStorage.setItem("cart",cartState.items);
        },
        [cartDispatch, cartState.items],
    ); 
    return (
        <CartContext.Provider value={{
            cartState,
           add,
           remove,
        }}
        >
            {children}
        </CartContext.Provider>
    );

}
export default CartContext;