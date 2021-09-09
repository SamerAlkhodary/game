
import LoginPage from './pages/login';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import { UserContextProvider } from './contexts/userContext';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom';
import SignupPage from './pages/signup';
import shopPage from './pages/shop';
function App() {
  return (
    <Router>
      <UserContextProvider>
       <Switch>
         <Route path= "/" exact component= {LoginPage}/>
         <Route path= "/signup" component= {SignupPage}/>
         <Route path= "/shop" component= {shopPage}/>
     
       </Switch>

    </UserContextProvider>

    </Router>
    
  
  );
}

export default App;
