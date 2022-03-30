import React, { Component } from "react";

// ==================
// RestClient
// ==================
import RestClient from "../api/client";
import JWTToken from "./utils/token";

// ==================
// Routes
// ==================
import AppRouter from './router';

// Login
import LoginModule from "./modules/login/Module";

class AppWrapper extends Component {

    constructor(props) {
        super(props);

        // Routes && Config
        this.config = this.props.config;

        // API Client
        this.api = new RestClient({
            environment: this.props.config.environment,
            config: this.props.config,
            app: this
        });

        this.state = {
            isLoggedIn: false,
            token: null
        }
    }

    componentDidMount() {
        this.loginWithToken();
    }

    getToken() {
        const tokenString = sessionStorage.getItem('auth_token');
        if (!tokenString) {
            return null;
        }

        const userToken = JSON.parse(tokenString);
        return new JWTToken({
            jti: userToken.jti,
            exp: userToken.exp,
            iat: userToken.iat,
            token: userToken.token,
            agentId: userToken.id
        })
    }

    setToken(token) {
        // Save to sessionStorage
        sessionStorage.setItem('auth_token', token.stringify());
        
        this.setState({
            isLoggedIn: true,
            token: token
        })
    }

    login({ identifier, password }) {

        // Login
        let ok = this.api.authorize({
            identifier: identifier,
            password: password
        })

        if (!ok) {
            return false;
        }

        return true;
    }

    loginWithToken() {

        const token = this.getToken();
        if (token && !token.isValid() || !token) {
            return;
        }

        // Login
        this.setState({
            isLoggedIn: true,
            token: token
        })
    }

    hasToLogIn() {
        return !this.state.isLoggedIn || !this.state.token || this.state.token && !this.state.token.isValid();
    }

    logout() {
        sessionStorage.removeItem('auth_token');
        this.setState({
            isLoggedIn: false,
            token: null
        })
    }

    render() {
        if (this.hasToLogIn()) {
            return <LoginModule app={this} config={this.props.config} />
        }
        return <AppRouter app={this} />
    }

}

export default AppWrapper;