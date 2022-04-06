import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';
import FlightRoute from './route';

class FlightsContent extends Component {

    render() {
        return (
            <div className="w-full | py-4">
                <div className="flex flex-col | border-b-2 border-jt-primary | mx-4 pb-4">
                    <div>
                        <h2 className="text-6xl mb-4 sm:mb-0 sm:text-3xl">Flights</h2>
                        {/* Loading animation */}
                    </div>
                    <div>
                        <p className="text-xl sm:text-sm  text-gray-400">
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