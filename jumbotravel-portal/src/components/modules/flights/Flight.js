import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';
import withRouter from '../../utils/router';
import ClassName from '../../utils/classname';

import { getSVG, getBackground, getOutlineBackground } from '../../utils/types';

function Detail(props) {
    return (
        <div className="w-full sm:w-1/2">
            <p className="text-gray-500 text-xl sm:text-xs font-bold | mb-2 sm:mb-0">{props.title}</p>
            <p className="text-brand-blue text-xl sm:text-xs">{props.content}</p>
        </div>
    )
}

function FlightOperation(props) {
    return (
        <div className="flex flex-col | w-full | items-start justify-start | p-1 | space-y-4">
            <div
                className={ClassName(
                    getOutlineBackground(props.notification.type),
                    "bg-gray-50 | shadow hover:shadow-md rounded-md | w-full sm:h-12 | flex flex-col sm:flex-row justify-between items-center | p-2 | cursor-pointer | outline"
                )}
                onClick={() => {
                    this.context.pushLocalNotification({
                        title: "Flight Operation",
                        type: props.notification.type
                    });
                }}
            >
                <div className="flex | items-center justify-between sm:justify-start w-full | space-x-4">
                    <div className="flex items-center justify-start">
                        {
                            getSVG(props.notification.type)
                        }
                        <a 
                            href={
                                props.notification.extra && props.notification.extra.agentid ?
                                `/agent/${props.notification.extra.agentid}` :
                                null
                            }
                            className="ml-5 text-md sm:text-xs text-brand-blue underline"
                        >
                            {
                                props.notification.extra ? props.notification.extra["agent"] : "No agent assigned"
                            }
                        </a>
                    </div>
                    <p className="sm:hidden text-gray-500 text-xs font-light | flex | items-end justify-center | space-x-1">
                        {/* <span>2022-02-02</span>
                        <span>00:00:00</span> */}
                        {
                            props.notification.created_at
                        }
                    </p>
                    <p className="hidden sm:block text-xs text-gray-500">
                        {
                            props.notification.message
                        }
                    </p>
                </div>

                <div className="mt-3 sm:mt-0 w-full">
                    <p className="sm:hidden text-xl text-gray-500">
                        {
                            props.notification.message
                        }
                    </p>
                    <p className="hidden sm:flex text-gray-500 text-xs font-light | items-center justify-end | w-full">
                        {
                            props.notification.created_at
                        }
                    </p>
                </div>
            </div>
        </div>
    );
}

class FlightDetails extends Component {

    render() {
        return (
            <div className="flex flex-col | items-start justify-start | w-full sm:w-3/5 sm:h-full">
                {/* Control */}
                <div className="flex flex-row | items-center justify-center | text-3xl | space-x-3 | p-4">
                    <a
                        href="/flights"
                        className="text-gray-500"
                    >Flights</a>
                    <span className="text-gray-500">\</span>
                    <p className="font-bold text-brand-blue">Flight {this.props.router.params.id}</p>
                </div>
                {/* Content */}
                <div className="flex flex-col | items-start justify-start | w-full h-full | p-4">
                    <div className="bg-white | w-full h-full | rounded-md | shadow p-4">
                        <div className="mb-4 | flex items-center justify-between">
                            <h4 className="text-2xl sm:text-xl font-bold text-brand-blue">Flight Details</h4>
                            <p
                                className={ClassName(
                                    this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'FLYING' ? 'bg-jt-primary3' : 'bg-jt-primary',
                                    "hidden sm:block sm:text-xs font-bold px-4 py-2 | rounded | text-white"
                                )}
                            >
                                {
                                    this.context.agentFlightDetails ?
                                        this.context.agentFlightDetails.status :
                                        ''
                                }
                            </p>
                        </div>
                        {
                            this.context.agentFlightDetails ?
                                <div className="flex flex-col sm:flex-row flex-wrap | items-start justify-start | w-full space-y-2 sm:space-y-0">
                                    <Detail title="Flight ID" content={this.context.agentFlightDetails.flight_id} />
                                    <Detail title="Carrier" content={this.context.agentFlightDetails.carrier} />
                                    <Detail title="Airplane Number" content={this.context.agentFlightDetails.flight_number} />
                                    <Detail title="Seats" content={this.context.agentFlightDetails.seats} />
                                    <Detail title="Departure" content={this.context.agentFlightDetails.departure_commonname} />
                                    <Detail title="Arrival" content={this.context.agentFlightDetails.arrival_commonname} />
                                    <Detail title="Departure Time" content={this.context.agentFlightDetails.departure_time} />
                                    <Detail title="Arrival Time" content={this.context.agentFlightDetails.arrival_time} />
                                    <Detail title="Departure Country" content={this.context.agentFlightDetails.departure_country} />
                                    <Detail title="Arrival Country" content={this.context.agentFlightDetails.arrival_country} />
                                    <div className="sm:hidden w-full sm:w-1/2">
                                        <p className="text-gray-500 text-xl sm:text-xs font-bold mb-2 sm:mb-0">Status</p>
                                        <p
                                            className={ClassName(
                                                this.context.agentFlightDetails.status === 'FLYING' ? 'bg-jt-primary3' : 'bg-jt-primary',
                                                "sm:hidden sm:text-xs font-bold px-4 py-2 | text-center | rounded | text-white"
                                            )}
                                        >
                                            {this.context.agentFlightDetails.status}
                                        </p>
                                    </div>
                                </div>
                                : ''
                        }
                    </div>
                </div>
            </div>
        )
    }

}
FlightDetails.contextType = Context;

class FlightOperations extends Component {

    render() {
        return (
            <div className="relative | flex flex-col | items-start justify-start | w-full sm:h-full">

                {/* Control */}
                <div className="flex flex-row | items-center justify-center | text-3xl | space-x-3 | p-4">
                    <p className="font-regular text-brand-blue">Flight Operations</p>
                </div>

                <div className="flex | w-full h-full | p-4">
                    <div className="relative | bg-white | shadow | rounded-md | flex flex-col | items-start justify-between | w-full h-full | overflow-hidden">
                        <div className="sm:absolute | w-full h-full | flex flex-col | items-start justify-start | sm:overflow-scroll | sm:no-scrollbar | sm:pb-16 | p-2">
                            {/* Flight Operations */}
                            {
                                this.context.agentFlightOperations && this.context.agentFlightOperations.notifications.length > 0 ?
                                    this.context.agentFlightOperations.notifications.map((item, index) => {
                                        return (
                                            <FlightOperation key={index} notification={item} />
                                        )
                                    })
                                    : 
                                    <div className="relative flex flex-col | items-start justify-start | w-full h-full | p-4">
                                        <p className="text-gray-500 text-xl font-bold">No Flight Operations</p>
                                    </div>
                            }
                        </div>
                        <div className="sm:absolute bottom-0 | w-full | flex flex-wrap sm:flex-nowrap items-center justify-between p-2 | bg-white">
                            <div>
                                <button className="text-lg sm:text-xs px-4 py-2 | bg-jt-primary | rounded | shadow hover:shadow-md | font-bold | text-white">
                                    Toggle State
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        )
    }

}
FlightOperations.contextType = Context;

class Content extends Component {

    render() {
        return (
            <div className="relative | flex flex-col | w-full h-full | sm:overflow-scroll | sm:no-scrollbar | bg-gray-50">
                {/* Top */}
                <div className="w-full bg-gray-50 sm:h-1/2 | flex flex-col sm:flex-row | items-center justify-center">
                    <FlightDetails router={this.props.router} />
                    <FlightOperations router={this.props.router} />
                </div>

                {/* Bottom */}
                <div className="w-full sm:h-1/2 | flex flex-col sm:flex-row | items-center justify-center">
                    <div className="bg-yellow-50 | w-full sm:w-3/5 sm:h-full">
                        <h3>This is the bottom left</h3>
                    </div>
                    <div className="bg-yellow-50 | w-full sm:h-full">
                        <h3>This is the bottom right</h3>
                    </div>
                </div>

            </div >
        )
    }

}

Content.contextType = Context;

class Module extends Component {

    constructor(props) {
        super(props);

        // this.state = {
        //     interval: null
        // };
    }

    componentDidMount() {

        // Get details
        this.context.getAgentFlightDetails(this.props.router.params.id);

        // Get operations
        this.context.getAgentFlightOperations(this.props.router.params.id);

    }

    componentWillUnmount() {
        // this.context.removeAgentFlight();
        // this.context.removeAgentFlightOperation();
    }

    componentDidUpdate() {
        if (this.context.hasToLogIn()) {
            this.context.logout();
        }
    }

    render() {
        return (
            <div className="w-screen h-screen flex flex-row flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800">
                <Helmet>
                    <title>
                        Flight {this.props.router.params.id} - JumboTravel
                    </title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={1} />
                <Notifications app={this.props.app} config={this.props.config} />
                <div className="relative w-full h-full | flex flex-col justify-start items-start">
                    {/* NavBar */}
                    <NavBar app={this.props.app} config={this.props.config} />
                    {/* Content */}
                    <Content router={this.props.router} />
                </div>
            </div>
        );
    }

}

Module.contextType = Context;

export default withRouter(Module);