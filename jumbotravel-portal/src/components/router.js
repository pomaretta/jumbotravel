import React from "react";
import {
    BrowserRouter as Router,
    Routes,
    Route,
} from "react-router-dom"

// Dashboard
import DashboardModule from "./modules/dashboard/Module";

// Flights
import FlightsModule from './modules/flights/Module';
import FlightIndividual from './modules/flights/Flight';

// Bookings
import BookingsModule from './modules/bookings/Module';
import BookingIndividual from './modules/bookings/Booking';

import PlanesModule from './modules/planes/Module';
import NotFound from './modules/404';

import AppContext from "./context/app";

class AppRouter extends React.Component {

    initApplication() {

        // Validate session
        this.context.validateSession({
            token: this.context.token
        });

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
            // Validate session
            // this.context.validateSession({
            //     token: this.context.token
            // });
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

                    {/* General Routes */}
                    {/* <Route key='login' path='/login' element={<LoginModule app={this.props.app} config={this.props.config} />} /> */}
                    <Route key='home' path='/' element={<DashboardModule app={this.props.app} config={this.props.config} />} />
                    <Route key='dashboard' path='/dashboard' element={<DashboardModule app={this.props.app} config={this.props.config} />} />

                    {/* Assistant */}
                    <Route key='flights' path='/flights' element={<FlightsModule app={this.props.app} config={this.props.config} />} />
                    <Route
                        key='flights_details'
                        path="/flights/:id"
                        element={<FlightIndividual />}
                    />

                    {/* Bookings */}
                    <Route key='bookings' path='/bookings' element={<BookingsModule app={this.props.app} config={this.props.config} />} />
                    <Route
                        key='bookings_details'
                        path="/bookings/:id"
                        element={<BookingIndividual />}
                    />

                    {/* 404 Not Found */}
                    <Route key='404' path='*' element={<NotFound app={this.props.app} config={this.props.config} />} />
                </Routes>
            </Router>
        );
    }

}

AppRouter.contextType = AppContext;

export default AppRouter;
