// eslint-disable-next-line no-unused-vars
import React, {useEffect, useState} from 'react';
import { FaEye } from "react-icons/fa6"; //eye icon to hide/show password
import { FaEyeSlash } from "react-icons/fa6";
// import LoginButton from './common/LoginButton'
// import Layout from './components/layout'


 

function Login(){
    //use ternary operator to toggle between states and update eyelash
    const [showPassword, setShowPassword] = useState(""); //initial state(empty string)
    return(
        <div className= "login-main"> 
            <h1>Login Here</h1>
            <p>Please enter your details</p>
        <form>
            <input type="email" placeholder='Email'/>
        <div className= "password-main">
        <input type={showPassword ? "text" : "password"} placeholder="Password" />

        {
  showPassword ? 
    React.createElement(FaEyeSlash, { onClick: function() { setShowPassword(!showPassword); } }) :
    React.createElement(FaEye, { onClick: function() { setShowPassword(!showPassword); } })
} 
        </div>
        <div className = "login-button">
         <button type="button"><a href="#"> Log In </a></button>
        </div>
        </form>
        <p className="register-main">
            No Account. <a href="#">Sign Up! </a>
        </p>
        </div>
        
    )
}


export default Login;