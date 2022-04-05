import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';

class FlightsContent extends Component {

    render() {
        return (
            <div className="w-full | py-4">
                <div className="flex flex-col | border-b-2 border-jt-primary | mx-4 pb-4">
                    <h2 className="text-3xl">Flights</h2>
                    <div>
                        <p className="text-sm  text-gray-400">
                            Flights ordered by routes.
                        </p>
                    </div>
                </div>
                <div className="flex flex-col | items-start justify-center">
                    {
                        // Map routes
                        this.context.agentFlights ?
                            (
                                this.context.agentFlights.getFlightsByRoutes().map((route, index) => {
                                    return (
                                        <div key={index} className="flex flex-col | items-start justify-start | w-full">
                                            <h3 className="text-xl | font-semibold | text-brand-blue p-4">
                                                {route.route.departure} - {route.route.arrival}
                                            </h3>
                                            <div className="flex flex-col | items-start justify-start | w-full px-2">
                                                {
                                                    route.flights.map((flight, index) => {
                                                        return (
                                                            <a
                                                                href={`/flights/${flight.flight_id}`}
                                                                key={index}
                                                                className="flex flex-row | justify-between | items-center | bg-white | w-full | p-2 px-4 | rounded-lg | shadow-sm hover:shadow-md"
                                                            >
                                                                <div>
                                                                    <p className="text-md | text-brand-blue | font-semibold">
                                                                        {flight.carrier} | {flight.flight_number}
                                                                    </p>
                                                                </div>
                                                                {/* Controls */}
                                                                <div className="flex flex-row | justify-center items-center | space-x-6">
                                                                    {/* Warning SVG */}
                                                                    <svg className="w-6 h-6 | fill-current | text-red-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z" /></svg>
                                                                    <p className="bg-jt-primary3 hover:bg-brand-blue-dark text-white font-bold py-2 px-4 rounded">
                                                                        {flight.status}
                                                                    </p>
                                                                </div>
                                                            </a>
                                                        )
                                                    })
                                                }
                                            </div>
                                        </div>
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
            <div className="relative | flex flex-col | w-full h-full | overflow-scroll | no-scrollbar">
                {/* Top */}
                <div className="w-full | flex flex-col items-center justify-center">
                    <FlightsContent />
                    {/* <div className="bg-yellow-50 | w-full | p-2">
                        <button
                            className="bg-brand-blue text-black font-bold py-2 px-4 rounded"
                            onClick={() => {
                                this.context.pushLocalNotification({
                                    title: 'Hello World',
                                    message: null,
                                    link: null,
                                    extra: null,
                                    type: 'INFO'
                                });
                            }}
                        >
                            Send notification
                        </button>
                    </div> */}
                </div>

                {/* Bottom */}
                <div className="w-full | flex flex-col items-center justify-center">
                    {/* <div className="bg-yellow-50 | w-full | p-2">
                        <h3>This is the bottom left</h3>
                    </div>
                    <div className="bg-yellow-50 | w-full | p-2">
                        <h3>This is the bottom right</h3>
                    </div> */}
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
        return (
            <div className="w-screen h-screen min-h-screen flex flex-row flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800">
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