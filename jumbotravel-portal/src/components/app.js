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
import NotificationCollection from "../api/collection/notification";
import APIError from "../api/error";

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
            newNotifications: false,

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
        this.setState({
            intervals: {
                ...this.state.intervals,
                [interval]: interval
            }
        })
        // this.state.intervals[interval] = interval;
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
        // const tokenString = sessionStorage.getItem('auth_token');
        const tokenString = this.getCookie('auth_token');
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

    getCookie(name) {
        var value = "; " + document.cookie;
        var parts = value.split("; " + name + "=");
        if (parts.length === 2) return parts.pop().split(";").shift();
    }

    addCookie({
        token,
        expires
    }) {
        document.cookie = `auth_token=${token}; expires=${expires}; path=/`;
    }

    removeCookie(name) {
        document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:01 GMT; path=/;`;
    }

    setToken(token) {
        // Save to sessionStorage
        // sessionStorage.setItem('auth_token', token.stringify());
        this.addCookie({
            token: token.stringify(),
            expires: new Date(Date.now() + (1000 * 60 * 60 * 24 * 7))
        })
        
        this.setState({
            isLoggedIn: true,
            token: token
        })
    }

    login({ identifier, password }) {

        // TODO: Handle error
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
        // sessionStorage.removeItem('auth_token');
        this.removeCookie('auth_token');
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

    async validateSession({token}) {

        // Validate token
        if (!token || !token.isValid()) {
            this.logout();
        }

        // Validate agent
        let validate = null;
        try {
            validate = await this.api.validate({
                token: token
            });
        } catch (e) {
            this.logout();
        }

        if (!validate) {
            this.logout();
        }

        return;
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
            if (e instanceof APIError) {
                if (e.getStatus() === 401) {
                    this.logout();
                }
            }
            return;
        }

        let newNotifications = false;
        let rawLocalNotifications = JSON.parse(localStorage.getItem('agent_notifications'));
        let finalNotifications = null;
        if (rawLocalNotifications && rawLocalNotifications.notifications) {

            // Convert RAW JSON to Collection of Notifications
            let localNotifications = NotificationCollection.parse(rawLocalNotifications.notifications);

            // =====================
            // 1. Get new notifications and add to finalNotifications
            // 2. Update notifications if active has changed
            // 3. Remove old notifications that are expired and readed
            // 4. Check if there are notifications that are not readed
            // 5. Remove old notifications that are expired and readed (commonly forced by the system)
            // =====================

            // 1. Get new notifications and add to finalNotifications
            
            // Iterate over new notifications
            notifications.notifications.forEach(notification => {
                
                // Check if notification is already in localNotifications
                let localNotification = localNotifications.notifications.find(notif => {
                    return notif.getId() === notification.getId();
                });

                // If not found, add it at the start
                if (!localNotification) {
                    localNotifications.notifications.unshift(notification);
                    newNotifications = true;
                }

            });

            // 2. Update notifications if active has changed
            localNotifications.notifications.forEach(notification => {
                let notificationInServer = notifications.notifications.find(notif => {
                    return notif.getId() === notification.getId();
                });
                if (notificationInServer && notificationInServer.getActive() !== notification.getActive()) {
                    notification.setActive(notificationInServer.getActive());
                }
                if (notificationInServer && notificationInServer.getSeen() !== notification.getSeen()) {
                    notification.setSeen(notificationInServer.getSeen());
                }
            });

            // 3. Remove old notifications that are expired and readed or not active
            localNotifications.notifications = localNotifications.notifications.filter(notification => {
                return !notification.isExpired() && !notification.isSeen() && notification.getActive();
            });

            // 4. Check if there are notifications that are not readed
            let notSeenNotifications = localNotifications.notifications.filter(notification => {
                return !notification.isSeen() && notifications.type != "GLOBAL";
            });
            if (notSeenNotifications && notSeenNotifications.length > 0) {
                newNotifications = true;
            }

            // 5. Get all notifications that are not in new notifications
            let notInNewNotifications = localNotifications.notifications.filter(notification => {
                return !notifications.notifications.find(notif => {
                    return notif.getId() === notification.getId();
                });
            });
            // Remove all notifications that are not in new notifications
            notInNewNotifications.forEach(notification => {
                localNotifications.notifications.splice(localNotifications.notifications.indexOf(notification), 1);
            });

            // 6. Update localNotifications with new notifications
            localNotifications.notifications.forEach(notification => {
                let notificationInServer = notifications.notifications.find(notif => {
                    return notif.getId() === notification.getId();
                });
                // Update localNotification with new data using update method of Notification
                notification.update(notificationInServer);
            });

            finalNotifications = localNotifications;
        } else {
            finalNotifications = notifications;
            newNotifications = true;

            // Filter some data to not show 
            finalNotifications.notifications = finalNotifications.notifications.filter(notification => {
                return !notification.isExpired() && !notification.isSeen() && notification.getActive();
            });
        }

        // Update localStorage with finalNotifications
        localStorage.setItem('agent_notifications', JSON.stringify(finalNotifications));

        let hasNotifications = false
        if (finalNotifications && finalNotifications.notifications && finalNotifications.notifications.length > 0) {
            hasNotifications = true;
        }

        if (finalNotifications && finalNotifications.notifications.length == 0) {
            hasNotifications = false;
            newNotifications = false;
        }

        this.setState({
            notifications: finalNotifications,
            newNotifications: newNotifications,
            hasNotifications: hasNotifications,
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

    async markNotificationsRead(notificationIds) {

        if (!notificationIds || notificationIds.length === 0) {
            return;
        }

        // If notificationsIds is a single notification, convert it to an array
        if (!Array.isArray(notificationIds)) {
            notificationIds = [notificationIds];
        }

        let notifications = this.state.notifications.getAll().filter(notification => {
            return notificationIds.includes(notification.getId());
        });

        // Mark notifications as read
        try {
            await this.api.markNotificationAsRead({
                token: this.state.token,
                notifications: notifications
            });
        } catch (e) {
            if (e instanceof APIError) {
                if (e.getStatus() === 401) {
                    this.logout();
                }
            }
            return;
        }

        // Update notifications
        await this.getNotifications();

        return;
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
            if (e instanceof APIError) {
                if (e.getStatus() === 401) {
                    this.logout();
                }
            }
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
            newNotifications: this.state.newNotifications,
            getNotifications: this.getNotifications.bind(this),
            isNotificationsOpen: this.isNotificationsOpen.bind(this),
            setNotificationsOpen: this.setNotificationsOpen.bind(this),
            markNotificationsRead: this.markNotificationsRead.bind(this),
            
            // =====================
            // Auth
            // =====================
            isLoggedIn: this.state.isLoggedIn,
            token: this.state.token,
            hasToLogIn: this.hasToLogIn.bind(this),
            logout: this.logout.bind(this),
            login: this.login.bind(this),
            validateSession: this.validateSession.bind(this),
        }}>
            <AppRouter config={this.config} />
        </AppContext.Provider>
    }

}

export default AppWrapper;