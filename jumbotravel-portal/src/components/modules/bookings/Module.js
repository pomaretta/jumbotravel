import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';
import ClassName from '../../utils/classname';

function Booking(props) {

    let statusBackground = "text-red-500";
    switch (String(props.booking.status).toLowerCase()) {
        case "pending":
            statusBackground = "text-yellow-500";
            break;
        case "completed":
            statusBackground = "text-green-500";
            break;
        case "cancelled":
            statusBackground = "text-red-500";
            break;
        default:
            statusBackground = "text-red-500";
            break;
    }

    return (
        <tr className="bg-white border-b">
            {/* Reference ID */}
            <th scope="row" className="px-6 py-4 text-gray-900 whitespace-nowrap | underline | font-bold | text-brand-blue">
                <a
                    href={`/bookings/${props.booking.bookingreferenceid}`}
                >
                    {props.booking.bookingreferenceid}
                </a>
            </th>
            {/* Flight ID */}
            <td className={ClassName(
                !props.isProvider ? "underline" : "",
                "px-6 py-4 | font-bold | text-brand-blue"
            )}>
                <a
                    href={!props.isProvider ? `/flights/${props.booking.flight_id}` : null}
                >
                    {props.booking.flight_id}
                </a>
            </td>
            {/* Agent Name */}
            <td className={ClassName(
                !props.isProvider ? "underline" : "",
                "px-6 py-4 | font-bold | text-brand-blue"
            )}>
                <a
                    href={!props.isProvider ? `/agents/${props.booking.agent_id}` : null}
                >
                    {`${props.booking.agent_name} ${props.booking.agent_surname}`}
                </a>
            </td>
            {/* Provider Name */}
            <td className={ClassName(
                !props.isProvider ? "underline" : "",
                "px-6 py-4 | font-bold | text-brand-blue"
            )}>
                <a
                    href={!props.isProvider ? props.booking.provider_id ? `/agents/${props.booking.provider_id}` : null : null}
                >
                    {
                        props.booking.provider_id ?
                            `${props.booking.provider_name} ${props.booking.provider_surname}` :
                            ""
                    }
                </a>
            </td>
            {/* Items Count */}
            <td className="px-6 py-4">
                {props.booking.items}
            </td>
            {/* Total */}
            <td className="px-6 py-4">
                {props.booking.total}€
            </td>
            {/* Created At */}
            <td className="px-6 py-4">
                {props.booking.created_at}
            </td>
            {/* Status */}
            <td
                className={`px-6 py-4 | font-bold | ${statusBackground}`}
            >
                {props.booking.status}
            </td>
        </tr>
    );
}

function MobileBooking(props) {

    let statusBackground = "text-red-500";
    switch (String(props.booking.status).toLowerCase()) {
        case "pending":
            statusBackground = "text-yellow-500";
            break;
        case "completed":
            statusBackground = "text-green-500";
            break;
        case "cancelled":
            statusBackground = "text-red-500";
            break;
        default:
            statusBackground = "text-red-500";
            break;
    }

    return (
        <div className="bg-white | p-4 | rounded-md | shadow-sm | flex flex-col justify-start items-start | space-y-4">
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Booking Reference ID</p>
                <a
                    href={`/bookings/${props.booking.bookingreferenceid}`}
                    className="text-brand-blue | font-bold | underline"
                >
                    {props.booking.bookingreferenceid}
                </a>
            </div>
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Flight</p>
                <a
                    href={!props.isProvider ? `/flights/${props.booking.flight_id}` : null}
                    className={ClassName(
                        !props.isProvider ? "underline" : "",
                        "font-bold | text-brand-blue"
                    )}
                >
                    {props.booking.flight_id}
                </a>
            </div>
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Agent</p>
                <a
                    href={!props.isProvider ? `/agents/${props.booking.agent_id}` : null}
                    className={ClassName(
                        !props.isProvider ? "underline" : "",
                        "font-bold | text-brand-blue"
                    )}
                >
                    {`${props.booking.agent_name} ${props.booking.agent_surname}`}
                </a>
            </div>
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Provider</p>
                <a
                    href={!props.isProvider ? props.booking.provider_id ? `/agents/${props.booking.provider_id}` : null : null}
                    className={ClassName(
                        !props.isProvider ? "underline" : "",
                        "font-bold | text-brand-blue"
                    )}
                >
                    {
                        props.booking.provider_id ?
                            `${props.booking.provider_name} ${props.booking.provider_surname}` :
                            ""
                    }
                </a>
            </div>
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Items</p>
                <p>
                    {props.booking.items}
                </p>
            </div>
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Total</p>
                <p>
                    {props.booking.total}€
                </p>
            </div>
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Created At</p>
                <p>
                    {props.booking.created_at}
                </p>
            </div>
            <div className="text-2xl space-y-2">
                <p className="text-brand-blue | font-bold">Status</p>
                <p
                    className={`font-bold | ${statusBackground}`}
                >
                    {props.booking.status}
                </p>
            </div>
        </div>
    );
}

class BookingsContent extends Component {

    render() {

        let isProvider = false;
        if (this.context.agent && this.context.agent.type === "PROVIDER") {
            isProvider = true;
        }

        return (
            <div className="w-full | py-4">
                <div className="flex flex-col | border-b-2 border-jt-primary | mx-4 pb-4">
                    <div>
                        <h2 className="text-6xl mb-4 sm:mb-0 sm:text-3xl">Bookings</h2>
                        {/* Loading animation */}
                    </div>
                    <div>
                        <p className="text-xl sm:text-sm  text-gray-400">
                            All bookings placed by agent
                        </p>
                    </div>
                </div>
                <div className="flex flex-col | items-start justify-center | w-full | px-4 | mt-5">
                    {/* Desktop */}
                    <div className="hidden sm:block relative overflow-x-auto shadow-md sm:rounded-lg | w-full">
                        <table className="w-full text-sm text-left text-gray-500">
                            <thead className="text-xs text-white uppercase bg-jt-primary">
                                <tr>
                                    <th scope="col" className="px-6 py-3">
                                        Booking Reference
                                    </th>
                                    <th scope="col" className="px-6 py-3">
                                        Flight
                                    </th>
                                    <th scope="col" className="px-6 py-3">
                                        Agent
                                    </th>
                                    <th scope="col" className="px-6 py-3">
                                        Provider
                                    </th>
                                    <th scope="col" className="px-6 py-3">
                                        Items
                                    </th>
                                    <th scope="col" className="px-6 py-3">
                                        Total
                                    </th>
                                    <th scope="col" className="px-6 py-3">
                                        Created At
                                    </th>
                                    <th scope="col" className="px-6 py-3">
                                        Status
                                    </th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.context.agentBookingsStatus ?
                                        (
                                            this.context.agentBookingsStatus.bookings.map((booking, index) => {
                                                return <Booking key={index} booking={booking} isProvider={isProvider} />
                                            })
                                        ) : (
                                            <tr>
                                                <td colSpan="8" className="px-6 py-4">
                                                    <p className="text-center">
                                                        No bookings found
                                                    </p>
                                                </td>
                                            </tr>
                                        )
                                }
                            </tbody>
                        </table>
                    </div>
                    {/* Mobile */}
                    <div className="sm:hidden w-full | flex flex-col | justify-start items-start | space-y-4">
                        {
                            this.context.agentBookingsStatus ?
                                (
                                    this.context.agentBookingsStatus.bookings.map((booking, index) => {
                                        return <MobileBooking key={index} booking={booking} isProvider={isProvider} />
                                    })
                                ) : (
                                    <tr>
                                        <td colSpan="8" className="px-6 py-4">
                                            <p className="text-center">
                                                No bookings found
                                            </p>
                                        </td>
                                    </tr>
                                )
                        }
                    </div>
                </div>
            </div>
        )
    }

}
BookingsContent.contextType = Context;

class Content extends Component {

    render() {
        return (
            <div className="relative | flex flex-col | w-full h-full | sm:overflow-scroll | sm:no-scrollbar | bg-gray-50">
                {/* Top */}
                <div className="w-full | flex flex-col items-center justify-center | bg-gray-50 | pb-6">
                    <BookingsContent />
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

        // // First load
        if (!this.context.agentBookingsStatus) {
            this.context.getAgentBookingsStatus();
        }

        // // Set interval
        if (!this.state.interval) {
            this.setState({
                interval: setInterval(() => {
                    console.log("Refreshing... Booking Status");
                    this.context.getAgentBookingsStatus();
                }, 25000)
            });
        }

    }

    componentWillUnmount() {
        clearInterval(this.state.interval);
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
                    <title>Bookings - JumboTravel</title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={2} />
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