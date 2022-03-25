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
import ROUTES from '../routes';
import LoginModule from "./modules/login/Module";

class App extends Component {

    constructor(props) {
        super(props);

        // Routes && Config
        this.routes = ROUTES;
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
            token: userToken.token
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
        if (!token) {
            return false;
        }

        // Login
        this.setState({
            isLoggedIn: true,
            token: token
        })
    }

    componentDidMount() {
        if (!this.state.isLoggedIn) {
            this.loginWithToken();
            return <LoginModule app={this} config={this.props.config} />
        }
    }

    render() {
        return (
            <AppRouter app={this} />
        );
    }

}

export default App;