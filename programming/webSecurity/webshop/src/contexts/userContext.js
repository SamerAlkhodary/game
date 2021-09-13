import React, { createContext, useCallback, useReducer } from 'react';
import source from '../dataSources';
const initialState = {
    user: null,
    token: null,
};
const actions = {
    Login: 'Login',
    Logout: 'Logout',
    Signup: 'Signup',
    Load: 'Load'
}
const reducer = (state, action) => {
    // eslint-disable-next-line default-case
    switch (action.type) {
        case actions.Login: {
            return { ...state, user: action.user, token: action.token }
        }
        case actions.Signup: {
            return { ...state, user: action.user, token: action.token }
        }
        case actions.Logout: {
            return initialState
        }
        case actions.Load: {
            return { ...state, user: action.user, token: action.token }
        }
    }
    return state;
}
const UserContext = createContext(initialState);
export const UserContextProvider = ({ children }) => {
    const [userState, userDispatch] = useReducer(reducer, initialState);
    const load = useCallback(
        () => {
            const user = localStorage.getItem('user');
            const token = localStorage.getItem('token');

            if (user && token) {
                userDispatch({
                    type: actions.Load,
                    user,
                    token
                });
            }
        }, [userDispatch]
    );
    const signup = useCallback(
        async (user) => {
            try {
                const result = await source.signup(user);
                if (result) {
                    userDispatch({
                        type: actions.Signup,
                        user: result.user,
                        token: result.token,
                    });
                    localStorage.setItem("user", String(result.user));
                    localStorage.setItem("token", String(result.token));

                }
            } catch (err) {
                console.warn("Signup faild", err);
                userDispatch({
                    type: actions.Signup,
                    user: {},
                    token: undefined,
                    error: err
                });

            }

        }, [userDispatch]
    );
    const logout = useCallback(
        () => {
      
            localStorage.removeItem("user" );
            localStorage.removeItem("token");
            userDispatch({
                type: actions.Logout,
            });



        }, [userDispatch]
    );

    const login = useCallback(
        async (username, password) => {
            try {
                const result = await source.login(username, password);
                if (result) {
                    userDispatch({
                        type: actions.Login,
                        user: result.user,
                        token: result.token,
                    });
                    localStorage.setItem("user", String(result.user));
                    localStorage.setItem("token", String(result.token));

                }
            } catch (err) {
                console.warn("Login faild", err);
                userDispatch({
                    type: actions.Login,
                    user: {},
                    token: undefined,
                    error: err
                });

            }

        }, [userDispatch]
    );
    return (
        <UserContext.Provider value={{
            userState,
            login,
            signup,
            load,
            logout
        }}
        >
            {children}
        </UserContext.Provider>
    );

}
export default UserContext;