import React, { Component } from 'react';
import './App.css';
import Authenticate from './Authenticate';
import CreateConsignment from './CreateConsignment';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      err: null,
      authenticated: false,
      token: this.getToken(),
    };
  }

  getToken = () => {
    return localStorage.getItem('token') || false;
  }

  isAuthenticated = () => {
    return this.state.authenticated || this.getToken() || false;
  }

  onAuth = token => {
    this.setState({
      authenticated: true,
      token,
    });
  }

  renderLogin = () => {
    return (
      <Authenticate onAuth={this.onAuth} />
    );
  }

  renderAuthenticated = () => {
    return (
      <CreateConsignment token={this.state.token} />
    );
  }

  render() {
    const authenticated = this.isAuthenticated();
    return (
      <div className="App">
        <div className="App-header">
          <h2>Shippy</h2>
        </div>
        <div className='App-intro container'>
          {(authenticated ? this.renderAuthenticated() : this.renderLogin())}
        </div>
      </div>
    );
  }
}

export default App;
