import React from 'react'
import Axios from 'axios';
import { Helmet } from 'react-helmet';

import LogoSVG from '../../utils/logo';

class LoginForm extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            loading: false,
            username: "",
            password: ""
        }

    }

    async userLogin(identifier, password) {
        return
    }

    async handleSubmit(event) {
        event.preventDefault();
        return
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                {/* Login Form */}
                <div className="px-6 pb-10 mt-14">
                    
                    {/* Username */}
                    <div className="w-full mb-2 flex flex-col items-start justify-center">
                        <label className="text-xl text-brand-blue text-md font-semibold mb-1 py-2" htmlFor="username">
                            Identifier
                        </label>
                        <div className="flex items-center justify-center w-full">
                            <input type='text' placeholder="" 
                                value={this.state.username} 
                                onChange={ev => this.setState({username: ev.target.value})}
                                className="px-3 py-2 | w-full | border-2 border-jt-primary | rounded-md text-gray-700 focus:outline-none" />
                        </div>
                    </div>
                    
                    {/* Password */}
                    <div className="w-full mb-2 flex flex-col items-start justify-center">
                        <label className="text-xl text-brand-blue text-md font-semibold mb-1 py-2" htmlFor="username">
                            Password
                        </label>
                        <div className="flex items-center justify-center w-full">
                            <input type='text' placeholder="" 
                                value={this.state.password} 
                                onChange={ev => this.setState({password: ev.target.value})}
                                className="px-3 py-2 | w-full | border-2 border-jt-primary | rounded-md text-gray-700 focus:outline-none" />
                        </div>
                    </div>

                    {/* Login */}
                    <button type="submit"
                        className="py-2 px-2 mt-10 | w-full | rounded-md text-xl font-bold bg-jt-primary text-gray-100 focus:outline-none">Login</button>
                </div>
            </form>
        )
    }

}

class Login extends React.Component {

    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div className="w-full h-screen flex items-center justify-center overflow-hidden">
                <Helmet>
                    <title>Login - JumboTravel</title>
                </Helmet>
                <div className="w-full md:w-1/3">
                    {/* Logo Image */}
                    <div className="flex font-bold justify-center mt-6 mb-5">
                        <img className="h-20 w-20"
                            src="/resources/logo.svg" />
                    </div>
                    {/* Logo Title */}
                    <h2 className="text-5xl text-center font-black">
                        <span className='text-brand-blue'>
                            jumbo
                        </span>
                        <span className='text-brand-green'>
                            travel
                        </span>
                    </h2>
                    {/* Login Form */}
                    <LoginForm env={this.props} />
                    <LogoSVG className="fixed w-25" />
                </div>
            </div>
        )
    }

}

export default Login;