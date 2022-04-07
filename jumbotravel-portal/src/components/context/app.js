import React from 'react';

const AppContext = React.createContext({

    // APP
    app: null,
    config: null,

    // FEATURES
    intervals: {},
    clearIntervals: () => {},
    addInterval: () => {},
    removeInterval: () => {},

    // =====================
    // AGENT DATA
    // =====================
    agent: null,
    getAgent: () => {},
    updateAgent: () => {},

    // =====================
    // Notifications
    // =====================
    notificationsIsOpen: false,
    notifications: [],
    localNotifications: null,
    hasNotifications: false,
    newNotifications: false,
    getNotifications: () => {},
    isNotificationsOpen: () => {},
    setNotificationsOpen: () => {},
    markNotificationsRead: () => {},
    pushLocalNotification: () => {},

    // =====================
    // Auth
    // =====================
    isLoggedIn: false,
    token: null,
    hasToLogIn: () => {},
    logout: () => {},
    login: () => {},
    validateSession: () => {},

    // =====================
    // Functionality
    // =====================
    agentFlights: null,
    agentFlightDetails: null,
    agentFlightOperations: null,
    getAgentFlights: () => {},
    getAgentFlightDetails: () => {},
    removeAgentFlightDetails: () => {},
    getAgentFlightOperations: () => {},
    removeAgentFlightOperation: () => {},

});

export default AppContext;