import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';
import FlightRoute from './route';

function CurrentFlight(props) {

    return (
        <div className="w-full | mt-5">
            <h3
                className="text-2xl sm:text-xl | font-semibold | text-brand-blue px-4 | w-3/6 sm:w-1/6"
            >
                Nearest Flight
            </h3>
            <div className="p-4 | w-full h-full">
                <a
                    href={`/flights/${props.flight.flight_id}`}
                    className="w-full | h-full | bg-white | rounded-xl | shadow-sm | p-4 | hover:shadow-md | cursor-pointer | flex"
                >
                    <div className="flex flex-col sm:flex-row | w-full h-full | items-center justify-between | sm:space-y-0 space-y-4">
                        {/* Logo | Airplane Details */}
                        <div className="flex flex-col sm:flex-row | w-full sm:w-auto | justify-center items-center | space-y-4 sm:space-y-0 sm:space-x-4">
                            {/* Logo */}
                            <img
                                src="/resources/spainair.png"
                                className="sm:h-10 h-16"
                            />
                            <div className="flex | justify-between w-full sm:w-auto sm:justify-center | items-center | sm:text-sm text-lg | sm:space-x-2">
                                <p className="font-bold text-brand-blue">Flight ID</p>
                                <p>{props.flight.flight_id}</p>
                            </div>
                            <div className="flex | justify-between w-full sm:w-auto sm:justify-center | items-center | sm:text-sm text-lg | sm:space-x-2">
                                <p className="font-bold text-brand-blue">Airplane</p>
                                <a
                                    className="underline font-bold text-brand-blue"
                                >{
                                    `${props.flight.carrier}-${props.flight.flight_number}`
                                }</a>
                            </div>
                            <div className="flex | justify-between w-full sm:w-auto sm:justify-center | items-center | sm:text-sm text-lg | sm:space-x-2">
                                <p className="font-bold text-brand-blue">Departure</p>
                                <p>{props.flight.departure_time}</p>
                            </div>
                            <div className="flex | justify-between w-full sm:w-auto sm:justify-center | items-center | sm:text-sm text-lg | sm:space-x-2">
                                <p className="font-bold text-brand-blue">Arrival</p>
                                <p>{props.flight.arrival_time}</p>
                            </div>
                            <div className="flex | justify-between w-full sm:w-auto sm:justify-center | items-center | sm:text-sm text-lg | sm:space-x-2">
                                <p className="font-bold text-brand-blue">Route</p>
                                <a
                                    href={`/routes/${props.flight.route_id}`}
                                    className="underline font-bold text-brand-blue"
                                >{props.flight.departure_airport}-{props.flight.arrival_airport}</a>
                            </div>
                        </div>
                        <div className="flex | w-full sm:w-auto | items-center justify-center | sm:space-x-4">
                            <p
                                className="px-6 py-2 | w-full sm:w-auto | text-center | text-white | font-bold | bg-green-500 | shadow-sm | rounded-md | sm:text-xs text-lg"
                            >
                                {props.flight.status}
                            </p>
                        </div>
                    </div>
                </a>
            </div>
        </div>
    );
}

class FlightsContent extends Component {

    render() {
        let flight = null;
        if (this.context.agentFlights) {
            flight = this.context.agentFlights.getCurrent();
        }

        return (
            <div className="w-full | py-4">
                <div className="flex flex-col | border-b-2 border-jt-primary | mx-4 pb-4">
                    <div>
                        <h2 className="text-6xl mb-4 sm:mb-0 sm:text-3xl">Flights</h2>
                        {/* Loading animation */}
                    </div>
                    <div>
                        <p className="text-xl sm:text-sm  text-gray-400">
                            Current flight and ordered by routes.
                        </p>
                    </div>
                </div>
                {/* Current Flight */}
                {
                    flight ?
                        <CurrentFlight flight={flight} />
                        : ''
                }
                <div className="flex flex-col | items-start justify-center">
                    {
                        // Map routes
                        this.context.agentFlights ?
                            (
                                this.context.agentFlights.getFlightsByRoutes().map((route, index) => {
                                    return (
                                        <FlightRoute key={index} index={index} route={route} />
                                    );
                                })
                            ) :
                            (
                                <div className="w-full | flex flex-col items-center justify-center">
                                    <p>Loading...</p>
                                </div>
                            )
                    }
                </div>
            </div>
        )
    }

}

FlightsContent.contextType = Context;

class Content extends Component {

    render() {
        return (
            <div className="relative | flex flex-col | w-full h-full | sm:overflow-scroll | sm:no-scrollbar | bg-gray-50">
                {/* Top */}
                <div className="w-full | flex flex-col items-center justify-center | bg-gray-50 | pb-6">
                    <FlightsContent />
                </div>

                {/* Bottom */}
                <div className="w-full | flex flex-col items-center justify-center">
                </div>

            </div>
        )
    }

}

Content.contextType = Context;

class Module extends Component {

    constructor(props) {
        super(props);

        this.state = {
            interval: null
        };

    }

    componentDidMount() {

        // First load
        if (!this.context.agentFlights) {
            this.context.getAgentFlights();
        }

        // Set interval
        if (!this.state.interval) {
            this.setState({
                interval: setInterval(() => {
                    this.context.getAgentFlights();
                }, 25000)
            });
        }

    }

    componentWillUnmount() {
        this.context.removeInterval(this.state.interval);
    }

    componentDidUpdate() {
        if (this.context.hasToLogIn()) {
            this.context.logout();
        }
    }

    render() {

        // If the user is not assistant, redirect to dashboard
        if (this.context.agent && this.context.agent.type !== "ASSISTANT") {
            document.location.href = "/";
        }

        return (
            <div className="w-screen h-screen flex flex-row flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800">
                <Helmet>
                    <title>Flights - JumboTravel</title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={1} />
                <Notifications app={this.props.app} config={this.props.config} />
                <div className="relative w-full h-full | flex flex-col justify-start items-start">
                    {/* NavBar */}
                    <NavBar app={this.props.app} config={this.props.config} />
                    {/* Content */}
                    <Content />
                </div>
            </div>
        );
    }

}

Module.contextType = Context;

export default Module;