import React from 'react'
import { Helmet } from 'react-helmet';
import { Navigate } from 'react-router-dom';

import LogoSVG from '../../utils/logo';

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

class LoginForm extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            error: true,
            errorMessage: ''
        }

        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit = async (event) => {
        event.preventDefault();

        this.setState({
            error: false,
            errorMessage: ''
        });

        // Login
        let ok, error = await this.props.app.login({
            identifier: this.props.username,
            password: this.props.password
        })
        
        if (!ok) {
            this.setState({
                error: true,
                errorMessage: error
            });
        }

        return;
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
                                value={this.props.username}
                                onChange={ev => this.props.setUsername(ev.target.value)}
                                className={classNames(
                                    this.state.error && this.state.errorMessage === "identifier not found" ? 'border-red-500' : 'border-jt-primary',
                                    "px-3 py-2 | w-full | border-2 | rounded-md text-gray-700 focus:outline-none"
                                )}
                                required
                            />
                        </div>
                        {
                            this.state.error && this.state.errorMessage === "identifier not found" ?
                                (
                                    <div className='w-full h-8 | flex flex-row items-center justify-start'>
                                        <svg className="w-6 h-6 mr-2 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                                        </svg>
                                        <p className="text-sm text-red-500">
                                            Wrong identifier
                                        </p>
                                    </div>
                                ) :
                                (
                                    <div></div>
                                )
                        }
                    </div>

                    {/* Password */}
                    <div className="w-full mb-2 flex flex-col items-start justify-center">
                        <label className="text-xl text-brand-blue text-md font-semibold mb-1 py-2" htmlFor="username">
                            Password
                        </label>
                        <div className="flex items-center justify-center w-full">
                            <input type='password' placeholder=""
                                value={this.props.password}
                                onChange={ev => this.props.setPassword(ev.target.value)}
                                className={classNames(
                                    this.state.error && this.state.errorMessage === "Unauthorized" ? 'border-red-500' : 'border-jt-primary',
                                    "px-3 py-2 | w-full | border-2 | rounded-md text-gray-700 focus:outline-none"
                                )}
                                required
                            />
                        </div>
                        {
                            this.state.error && this.state.errorMessage === "Unauthorized" ?
                            (
                                <div className='w-full h-8 | flex flex-row items-center justify-start'>
                                    <svg className="w-6 h-6 mr-2 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                                    </svg>
                                    <p className="text-sm text-red-500">
                                        Wrong password
                                    </p>
                                </div>
                            ) :
                            (
                                <div></div>
                            )
                        }
                    </div>

                    {/* Login */}
                    <button type="submit"
                        className="py-2 px-2 mt-10 | w-full | rounded-md text-xl font-bold bg-jt-primary text-gray-100 focus:outline-none">Login</button>
                </div>
            </form >
        )
    }

}

class Login extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            loading: false,
            error: false,
            errorMessage: null,
            username: "",
            password: ""
        }

        this.setUsername = this.setUsername.bind(this);
        this.setPassword = this.setPassword.bind(this);
    }

    setUsername = (username) => {
        this.setState({
            username: username
        });
    }

    setPassword = (password) => {
        this.setState({
            password: password
        });
    }

    render() {

        if (this.props.app.state.isLoggedIn) {
            document.location.href = "/";
        }

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
                    <LoginForm
                        app={this.props.app}
                        setUsername={this.setUsername}
                        setPassword={this.setPassword}
                        username={this.state.username}
                        password={this.state.password} />
                    <LogoSVG className="fixed w-25" />
                </div>
            </div>
        )
    }

}

export default Login;