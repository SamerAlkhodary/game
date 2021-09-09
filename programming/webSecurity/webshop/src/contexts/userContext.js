import React,{createContext,useCallback, useReducer} from 'react';
import source from '../dataSources';
const initialState = {
    user: null,
    token: null,
};
const actions = {
    Login: 'Login',
    Logout: 'Logout',
    Signup: 'Signup'
}
const reducer = (state, action)=>{
    switch (action.type){
        case actions.Login:{
            return {...state,user: action.user, token: action.token}
        }
        case actions.Signup:{
            return {...state,user: action.user, token: action.token}
        }
        case actions.Logout:{
            return {}
        }
    }
    return state;
}
const UserContext = createContext(initialState);
export const UserContextProvider = ({children}) =>{
    const [userState, userDispatch] = useReducer(reducer, initialState);
    
    const signup= useCallback(
        async(user)=>{
            try{
                const result = await source.signup(user);
                if(result){
                    userDispatch({
                        type: actions.Signup,
                        user: result.user,
                        token: result.token,
                    });
    
                }
            }catch(err){
                console.warn("Signup faild",err);
                userDispatch({
                    type: actions.Signup,
                    user: {},
                    token: undefined,
                    error: err
                });

            }

        }
    );

    const login = useCallback(
        async (username, password) => {
            try{
                const result = await source.login(username, password);
                if(result){
                    userDispatch({
                        type: actions.Login,
                        user: result.user,
                        token: result.token,
                    });
    
                }
            }catch(err){
                console.warn("Login faild",err);
                userDispatch({
                    type: actions.Login,
                    user: {},
                    token: undefined,
                    error: err
                });

            }
           
        },[userDispatch,source.login]
    );
    return (
        <UserContext.Provider value ={{
            userState,
            login,
            signup
        }}
        >
            {children}
        </UserContext.Provider>
    );

}
export default UserContext;