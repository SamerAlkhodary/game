
import LoginPage from './pages/login';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import { UserContextProvider } from './contexts/userContext';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom';
import SignupPage from './pages/signup';
import ShopPage from './pages/shop';
import { CartContextProvider } from './contexts/cartContext';
import CartPage from './pages/cartPage';
function App() {
  return (
    <Router>
      <UserContextProvider>
        <CartContextProvider>
       <Switch>
         <Route path= "/" exact component= {ShopPage}/>
         <Route path= "/signup" component= {SignupPage}/>
         <Route path= "/login" component= {LoginPage}/>
         <Route path= "/cart" component= {CartPage}/>
     
       </Switch>
       </CartContextProvider>
    </UserContextProvider>

    </Router>
    
  
  );
}

export default App;
