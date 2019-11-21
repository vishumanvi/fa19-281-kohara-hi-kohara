import React, {Component} from "react";
import axios from "axios";
import { Redirect } from "react-router";
import './styles/Login.css';

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: {
        firstname: "",
        lastname: "",
        email: "",
        password: ""
      },
      errors: {},
      dbErrors: ""
    };
  }

  doSubmit = () => {

  };


  render() {
    return (
      <React.Fragment>
          <section id="w3hubs">
              <div class="container ex">
              
                <h1>
                  <img className="login-logo" src="https://i.pinimg.com/originals/a2/5f/4f/a25f4f58938bbe61357ebca42d23866f.png" />
                  <img className="login-image-text" src="https://www.edigitalagency.com.au/wp-content/uploads/instagram-logo-text-blue-png.png"/>
                  

                </h1>
                <form action="#" method="POST">
                <div class="form-group">
                  
                    <input type="name" class="form-control" placeholder="Enter username..." />
                  </div>
                  <div class="form-group">
                  
                    <input type="password" class="form-control" placeholder="Password" />
                  </div>
                  <button type="button" class="btn instagradient logbtn">Log In</button>
                </form>
                <h4>OR</h4>
                
                
          </div>
                <div class="box">
                  <p className="no-pad">Don't have account? <a href="/signup">Sign up</a></p>
                </div>
              
            
            </section>
       

      </React.Fragment>
    );
  }
}

export default Login;