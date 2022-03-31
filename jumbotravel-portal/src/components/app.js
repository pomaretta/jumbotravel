import React, { Component } from "react";
import { loadAnimation } from 'lottie-web';
import { defineLordIconElement } from 'lord-icon-element';

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

// Context
import AppContext from "./context/app";
import App from "./modules/app";

defineLordIconElement(loadAnimation);

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

            // =====================
            // FEATURES
            // =====================
            intervals: {},
            agent: null,

            // =====================
            // Authentication
            // =====================
            isLoggedIn: false,
            token: null,
            
            // =====================
            // Notifications
            // =====================
            notificationsIsOpen: false,
            notifications: null,
            hasNotifications: false,

        }
    }

    componentDidMount() {
        this.loginWithToken();
    }

    // ==================
    // FEATURES
    // ==================

    clearIntervals() {
        Object.keys(this.state.intervals).forEach(key => {
            clearInterval(this.state.intervals[key]);
        })
    }

    addInterval(interval) {
        if (this.state.intervals[interval]) {
            throw new Error(`Interval with id ${interval} already exists`);
        }
        this.state.intervals[interval] = interval;
    }

    removeInterval(id) {
        if (this.state.intervals[id]) {
            clearInterval(this.state.intervals[id]);
            delete this.state.intervals[id];
        }
    }

    // ==================
    // AUTHENTICATION
    // ==================

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
        this.clearIntervals();
        this.setState({
            isLoggedIn: false,
            token: null,
            intervals: {},
            notificationsIsOpen: false,
            notifications: [],
            hasNotifications: false
        })
    }

    // END AUTHENTICATION
    // ==================

    // ==================
    // NOTIFICATIONS
    // ==================

    async getNotifications() {
        
        let notifications = null;
        try {
            notifications = await this.api.getNotifications({
                token:  this.state.token
            });
        } catch (e) {
            // TODO: Handle error
            console.error(e);
            return;
        }

        let hasNotifications = false;
        if (notifications && notifications.getAll().length > 0) {
            hasNotifications = true;
        }

        this.setState({
            notifications: notifications,
            hasNotifications: hasNotifications
        });

        return this.state.notifications;
    }

    isNotificationsOpen() {
        return this.state.notificationsIsOpen;
    }

    setNotificationsOpen() {
        this.setState({
            notificationsIsOpen: !this.state.notificationsIsOpen
        })
    }

    // END NOTIFICATIONS
    // ==================

    // ==================
    // AGENT
    // ==================

    getAgent() {
        return this.state.agent;
    }

    async updateAgent() {

        let agent = null;
        try {
            agent = await this.api.getAgentData({
                token:  this.state.token
            });
        } catch (e) {
            // TODO: Handle error
            console.error(e);
            return;
        }

        this.setState({
            agent: agent
        });
    }

    // END AGENT
    // ==================

    render() {
        if (this.hasToLogIn()) {
            return <LoginModule app={this} config={this.props.config} />
        }
        return <AppContext.Provider value={{
            // =====================
            // FEATURES 
            // =====================
            intervals: this.state.intervals,
            clearIntervals: this.clearIntervals.bind(this),
            addInterval: this.addInterval.bind(this),
            removeInterval: this.removeInterval.bind(this),

            // =====================
            // AGENT
            // =====================
            agent: this.state.agent,
            getAgent: this.getAgent.bind(this),
            updateAgent: this.updateAgent.bind(this),

            // =====================
            // Notifications
            // =====================
            notificationsIsOpen: this.state.notificationsIsOpen,
            notifications: this.state.notifications,
            hasNotifications: this.state.hasNotifications,
            getNotifications: this.getNotifications.bind(this),
            isNotificationsOpen: this.isNotificationsOpen.bind(this),
            setNotificationsOpen: this.setNotificationsOpen.bind(this),
            
            // =====================
            // Auth
            // =====================
            isLoggedIn: this.state.isLoggedIn,
            token: this.state.token,
            hasToLogIn: this.hasToLogIn.bind(this),
            logout: this.logout.bind(this),
            login: this.login.bind(this),
        }}>
            <AppRouter config={this.config} />
        </AppContext.Provider>
    }

}

export default AppWrapper;