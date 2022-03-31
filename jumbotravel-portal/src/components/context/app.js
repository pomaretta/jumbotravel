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
    getNotifications: () => {},
    isNotificationsOpen: () => {},
    setNotificationsOpen: () => {},

    // =====================
    // Auth
    // =====================
    isLoggedIn: false,
    token: null,
    hasToLogIn: () => {},
    logout: () => {},
    login: () => {},

});

export default AppContext;