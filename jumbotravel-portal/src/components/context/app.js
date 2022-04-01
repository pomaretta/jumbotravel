import React from 'react';

const AppContext = React.createContext({

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
    hasNotifications: false,
    newNotifications: false,
    getNotifications: () => {},
    isNotificationsOpen: () => {},
    setNotificationsOpen: () => {},
    markNotificationsRead: () => {},

    // =====================
    // Auth
    // =====================
    isLoggedIn: false,
    token: null,
    hasToLogIn: () => {},
    logout: () => {},
    login: () => {},
    validateSession: () => {},

});

export default AppContext;