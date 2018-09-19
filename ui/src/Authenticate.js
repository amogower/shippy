import React from 'react';

class Authenticate extends React.Component {
  state = {
    authenticated: false,
    email: '',
    password: '',
    err: '',
  }

  login = () => {
    fetch(`http://localhost:8080/rpc`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        method: 'AuthService.Auth',
        service: 'shippy.auth',
        request: {
          email: this.state.email,
          password: this.state.password,
        },
      }),
    })
    .then(res => res.json())
    .then(res => {
      if (res.token) {
        this.props.onAuth(res.token);
        localStorage.setItem('token', res.token);
      }
    })
    .catch(err => this.setState({ err }));
  }

  signup = () => {
    fetch(`http://localhost:8080/rpc`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        method: 'AuthService.Create',
        service: 'shippy.auth',
        request: {
          email: this.state.email,
          password: this.state.password,
          name: this.state.name,
        },
      }),
    })
    .then(res => res.json())
    .then(res => {
      if (res.token) {
        this.props.onAuth(res.token);
        localStorage.setItem('token', res.token);
      }
    })
    .catch(err => this.setState({ err }));
  }

  setEmail = e => {
    this.setState({
      email: e.target.value,
    });
  }

  setPassword = e => {
    this.setState({
      password: e.target.value,
    });
  }

  setName = e => {
    this.setState({
      name: e.target.value,
    });
  }

  render() {
    return (
      <div className='Authenticate'>
        <div className='Login'>
          <div className='form-group'>
            <input
              type="email"
              onChange={this.setEmail}
              placeholder='E-Mail'
              className='form-control' />
          </div>
          <div className='form-group'>
            <input
              type="password"
              onChange={this.setPassword}
              placeholder='Password'
              className='form-control' />
          </div>
          <button className='btn btn-primary' onClick={this.login}>Login</button>
          <br /><br />
        </div>
        <div className='Sign-up'>
          <div className='form-group'>
            <input
              type='input'
              onChange={this.setName}
              placeholder='Name'
              className='form-control' />
          </div>
          <div className='form-group'>
            <input
              type='email'
              onChange={this.setEmail}
              placeholder='E-Mail'
              className='form-control' />
          </div>
          <div className='form-group'>
            <input
              type='password'
              onChange={this.setPassword}
              placeholder='Password'
              className='form-control' />
          </div>
          <button className='btn btn-primary' onClick={this.signup}>Sign-up</button>
        </div>
      </div>
    );
  }
}

export default Authenticate;
