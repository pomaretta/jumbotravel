// ===============
// Module: routes
// ===============

// import LoginModule from './components/modules/login/Module';

// ===============
// TESTING
// ===============

import Dashboard from "./components/modules/dashboard";
import Preferences from './components/modules/preferences';
import Login from './components/modules/login';

const routes = [
    // {
    //     path: '/',
    //     key: 'home',
    //     component: LoginModule,
    //     exact: true
    // },
    // {
    //     path: '/login',
    //     key: 'login',
    //     component: LoginModule,
    //     exact: true
    // }
    {
        path: '/dashboard',
        key: 'dashboard',
        component: Dashboard,
        exact: true
    },
    {
        path: '/preferences',
        key: 'preferences',
        component: Preferences,
        exact: true
    },
    {
        path: '/login',
        key: 'login',
        component: Login,
        exact: true
    }
];

export default routes;