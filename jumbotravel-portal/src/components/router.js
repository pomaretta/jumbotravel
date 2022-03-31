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

import AppContext from "./context/app";

class AppRouter extends React.Component {

    initApplication() {

        // Agent Data
        this.context.updateAgent();

        // Notifications
        this.context.getNotifications();

    }

    // ==================
    // EVENT LOOP FOR GENERAL EVENTS
    // ==================
    initEvents() {

        // ==================
        // 1 SECOND
        // ==================
        this.context.addInterval(setInterval(() => {
            // console.log('Interval');
        }, 1000));

        // ==================
        // 5 SECONDS
        // ==================
        this.context.addInterval(setInterval(() => {
            // console.log('Interval');
        }, 5000));

        // ==================
        // 15 SECONDS
        // ==================
        this.context.addInterval(setInterval(() => {
            this.context.getNotifications();
        }, 15000));

        // ==================
        // 30 SECONDS
        // ==================
        this.context.addInterval(setInterval(() => {
            // console.log('Interval'); 
        }, 30000));

        // ==================
        // 1 MINUTE
        // ==================
        this.context.addInterval(setInterval(() => {
            // console.log('Interval');
        }, 60000));

    }

    componentDidMount() {

        // Initial application
        this.initApplication();

        // Event loop
        this.initEvents();

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

AppRouter.contextType = AppContext;

export default AppRouter;
