import React from "react";
import {
    BrowserRouter as Router,
    Routes,
    Route
} from "react-router-dom"

import AppModule from "./modules/app";
import LoginModule from './modules/login/Module';
import FlightsModule from './modules/flights/Module';
import PlanesModule from './modules/planes/Module';
import NotFound from './modules/404';

class AppRouter extends React.Component {

    defineRoute({
        route,
        config
    }) {
        return <Route
            element={<route.component app={this.props.app} config={config} />}
            {...route}
        />
    }

    defineRoutes({
        routes,
        config
    }) {
        return routes.map(route => this.defineRoute({
            route: route,
            config: config
        }));
    }

    render() {
        return (
            <Router>
                <Routes>
                    <Route key='login' path='/login' element={<LoginModule app={this.props.app} config={this.props.config} />} />
                    
                    <Route key='home' path='/' element={<AppModule app={this.props.app} config={this.props.config} />} />
                    <Route key='flights' path='/flights' element={<FlightsModule app={this.props.app} config={this.props.config} />} />
                    <Route key='planes' path='/planes' element={<PlanesModule app={this.props.app} config={this.props.config} />} />
                    
                    <Route key='404' path='*' element={<NotFound app={this.props.app} config={this.props.config} />} />
                </Routes>
            </Router>
        );
    }

}

export default AppRouter;
