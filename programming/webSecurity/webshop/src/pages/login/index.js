import React,{useContext, useEffect, useState} from 'react';
import UserContext from '../../contexts/userContext';
import {Form, Button, Card} from 'react-bootstrap'
import { Link } from 'react-router-dom';
import ParticlesBg from 'particles-bg';
import { useHistory } from 'react-router-dom';
const LoginPage =() =>{
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const history = useHistory();
    const {login,userState}  = useContext(UserContext);

    const submit = (event)=>{
        event.preventDefault();
        login(username,password);       
    
    }

   const userNamechanged=(change)=>{
      setUsername(change.target.value);
   }
   const passwordChanged = (change)=>{
    setPassword(change.target.value)
   }

   useEffect(()=>{
       if(userState.token){
           history.push('/');
       }
   },[history,userState])


    return (
        <div style={{
            display: "flex",
            marginTop:'3rem',
            justifyContent: "center",
            alignItems: "center",
            flexFlow: "column"
           
        }}>
            <h1>
                Welcome to our shop
            </h1>
        <Card style={{
            width: '40%',
            marginTop:'3rem',
            }}>
            <Card.Body>
            <Form>
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control type="email" placeholder="Enter email" onChange={userNamechanged} />
         
         
        </Form.Group>
      
        <Form.Group className="mb-3" controlId="formBasicPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control type="password" placeholder="Password" onChange={passwordChanged} />
        
        </Form.Group>
        <Form.Group className="mb-3">
         <Link to="/signup" style={{ textDecoration: 'none',fontSize: 13}}>
         Create an account
         </Link>
        </Form.Group> 
        
        <Button variant="primary" type="submit" onClick={submit}>
          Submit
        </Button>
      
      </Form>
     
            </Card.Body>
        </Card>
        <ParticlesBg type="circle" num={10} bg={true} />
        </div>
       
       

    );
   
}
export default LoginPage;