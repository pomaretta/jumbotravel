// ===============
// Module: routes
// ===============

import LoginModule from './components/modules/login/Module';
import HomeModule from './components/modules/home/Module';
import NotFound from './components/modules/404';

const routes = [
    {
        path: '/',
        key: 'home',
        component: HomeModule,
        exact: true
    },
    {
        path: '/login',
        key: 'login',
        component: LoginModule,
        exact: true
    },
    {
        path: '*',
        key: '404',
        component: NotFound,
        exact: true
    }
];

export default routes;