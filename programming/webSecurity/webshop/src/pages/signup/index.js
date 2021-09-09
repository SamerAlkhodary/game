import React,{useContext, useEffect, useState} from 'react';
import UserContext from '../../contexts/userContext';
import {Form, Button, Card, Col, Row} from 'react-bootstrap'
import { Link } from 'react-router-dom';
import ParticlesBg from 'particles-bg'

const SignupPage=()=>{
    const [usernameError, setUsernameError] = useState("");
    const [passwordError, setPasswordError] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [rePasswordError, setrePasswordError] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [rePassword, setrePassword] = useState("");
    
    const {signup,userState}  = useContext(UserContext);
    const isFormValid=()=>{
        return validatePassword(password)&& validateEmail(username)&& rePassword === password && firstName!=="" && lastName!=="";
    }
    const submit = (event)=>{
        event.preventDefault();
        setrePasswordError(password === rePassword? "":"Passwords do not match" )
        if(isFormValid()) {
            signup(username,password);       
        }else{
            console.log("invalid email or password format");
        }

    }
    useEffect(()=>{
        console.log(userState);
    },[userState]);
   const fisrtNameChanged=(change)=>{
    setFirstName(change.target.value);

   }
   const lastNameChanged=(change)=>{
    setLastName(change.target.value);
   }
   const userNamechanged=(change)=>{
      setUsername(change.target.value);
      setUsernameError(validateEmail(change.target.value)|| change.target.value===""? "": "invalid email format");
   }
   const passwordChanged = (change)=>{
    setPassword(change.target.value)
    setPasswordError(validatePassword(change.target.value) ||change.target.value===""? "": "Password should contain lowercase, upper case letters, numbers and at least 10 charachters")
   }
   const rePasswordChanged = (change)=>{
    setrePassword(change.target.value)
    
   }

   const validateEmail = (email)=>{
    const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

    return re.test(String(email).toLowerCase());
    }
    const validatePassword = (password)=>{
        var lowerCaseLetters = /[a-z]/g;
        var upperCaseLetters = /[A-Z]/g;
        var numbers = /[0-9]/g;
        return password.match(lowerCaseLetters) && 
               password.match(upperCaseLetters) && 
               password.match(numbers) && 
               password.length >= 10;

    }
    return(
        <div style={{
            display: "flex",
            marginTop:'3rem',
            justifyContent: "center",
            alignItems: "center",
            flexFlow: "column"
           
        }}>
       <Card style={{
        width: '40%',
        marginTop:'3rem',
    
        
        }}>
           <Card.Body>
           <Form>
        <Row className="mb-3">
          <Form.Group as={Col} >
            <Form.Control type="text" placeholder="First Name" onChange={fisrtNameChanged} />
          </Form.Group>
      
          <Form.Group as={Col} >
            <Form.Control type="text" placeholder="Last Name"  onChange={lastNameChanged}/>
          </Form.Group>
        </Row>
      
        <Form.Group className="mb-3" >
          <Form.Label>Email</Form.Label>
          <Form.Control placeholder="Email" type="email" onChange={userNamechanged}/>
          <Form.Text className="text-muted">
            {usernameError}
        </Form.Text>
        </Form.Group>
      
        <Form.Group className="mb-3" >
          <Form.Label>Password</Form.Label>
          <Form.Control placeholder="Password" type="password" onChange={passwordChanged}/>
          <Form.Text className="text-muted">
            {passwordError}
        </Form.Text>
        </Form.Group>
        <Form.Group className="mb-3" >
          <Form.Label>Repeat password</Form.Label>
          <Form.Control placeholder="Repeat password" type="password" onChange={rePasswordChanged}/>
          <Form.Text className="text-muted">
            {rePasswordError}
        </Form.Text>
        </Form.Group>
      
        <Button variant="primary" type="submit" onClick={submit}>
          Create
        </Button>
      </Form>
           </Card.Body>
       </Card>
       <ParticlesBg type="circle" num={10} bg={true} />
       </div>
    )
}
export default SignupPage;