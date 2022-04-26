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

class BookingReport extends Component {

    render() {
        return (
            <div className="w-full | flex flex-col sm:flex-row justify-between items-center | mb-4 | py-3">
                <div className={ClassName(
                    "flex | justify-start items-center | space-x-4 w-full",
                )}>
                    <p className="text-md font-bold mr-4">Download Report Date</p>
                    <select className="select select-md select-info dark:bg-white w-48 | shadow-lg shadow-blue-100">
                        <option selected value={30}>30 Days</option>
                        <option value={1}>1 Day</option>
                        <option value={7}>7 Days</option>
                    </select>
                    <button className="btn btn-success btn-md | shadow-lg shadow-blue-100" onClick={() => { }}>
                        <span className="text-md">Download</span>
                    </button>
                </div>
                <div className="flex justify-between sm:justify-end items-center | w-full sm:w-2/4">

                </div>
            </div>
        );
    }

}
BookingReport.contextType = Context;

class BookingsContent extends Component {

    constructor(props) {
        super(props);

        this.state = {
            selectedDate: null,
            downloadLoading: false,
            downloadCompleted: false,
            downloadSuccess: false,
            downloadErrorMessage: null,
        }

        this.changeSelectedDate = this.changeSelectedDate.bind(this);
        this.downloadDaily = this.downloadDaily.bind(this);
    }

    changeSelectedDate(date) {
        this.setState({
            selectedDate: date
        });
    }

    async downloadDaily() {

        // Check if date input is filled
        if (!this.state.selectedDate) {
            this.context.pushLocalNotification({
                title: "Report date must be filled",
                type: "ERROR",
            })
            return;
        }

        if (this.state.downloadLoading) {
            return;
        }

        // Send request to server
        this.setState({
            downloadLoading: true,
            downloadCompleted: false,
            downloadSuccess: false,
            downloadErrorMessage: null,
        })

        let requestSuccess = false;
        let blob = null;

        await this.context.getReport(this.state.selectedDate)
            .then((resp) => {
                requestSuccess = true
                blob = resp;
            })
            .catch(error => {
                requestSuccess = false;
                this.setState({
                    downloadErrorMessage: error.message,
                })
            })

        setTimeout(() => {

            this.setState({
                downloadLoading: false,
                downloadCompleted: true,
                downloadSuccess: requestSuccess,
            });

            if (!requestSuccess) {
                // Local notification with error message
                this.context.pushLocalNotification({
                    title: 'Error downloading invoice',
                    message: this.state.downloadErrorMessage,
                    link: null,
                    extra: null,
                    type: "ERROR"
                });
            }

            setTimeout(() => {

                this.setState({
                    downloadCompleted: !this.state.downloadSuccess,
                    downloadSuccess: false,
                })

                if (window.navigator && window.navigator.msSaveOrOpenBlob) {
                    window.navigator.msSaveOrOpenBlob(blob);
                } else {
                    const data = window.URL.createObjectURL(blob);
                    var link = document.createElement('a');
                    link.href = data;
                    link.download = `invoice_${this.state.selectedDate}_${new Date().getTime()}.pdf`;
                    link.click();
                    setTimeout(function () {
                        // For Firefox it is necessary to delay revoking the ObjectURL
                        window.URL.revokeObjectURL(data);
                    }, 100);
                }
            }, 3000);
        }, 2500);


    }

    render() {

        let isProvider = false;
        if (this.context.agent && this.context.agent.type === "PROVIDER") {
            isProvider = true;
        }

        return (
            <div className="w-full | py-4">
                <div className="flex flex-col sm:flex-row | space-y-4 sm:space-y-0 | justify-between items-start sm:items-center | border-b-2 border-jt-primary | mx-4 pb-4">
                    <div className="flex flex-col">
                        <div>
                            <h2 className="text-6xl mb-4 sm:mb-0 sm:text-3xl">Bookings</h2>
                        </div>
                        <div>
                            <p className="text-xl sm:text-sm  text-gray-400">
                                All bookings placed by agent
                            </p>
                        </div>
                    </div>
                    {
                        this.context.agent && this.context.agent.type === "PROVIDER" ?
                            (
                                <div className="space-y-2 sm:space-y-0 sm:space-x-4 | flex flex-col sm:flex-row | justify-start items-start sm:justify-end sm:items-center | w-full">
                                    <p className="font-bold text-brand-blue text-2xl sm:text-sm">
                                        Download Daily Report
                                    </p>
                                    {/* Date field */}
                                    <input type="date" className="w-full | input input-md input-info dark:bg-white sm:w-48 | shadow-lg shadow-blue-100"
                                        onChange={(e) => { this.changeSelectedDate(e.target.value) }}
                                        defaultValue={this.state.selectedDate}
                                    />
                                    <button
                                        type="button"
                                        className="btn btn-success btn-md sm:w-auto w-full | text-white | shadow-lg shadow-blue-100"
                                        onClick={() => { this.downloadDaily() }}
                                    >
                                        Download
                                        {
                                            this.state.downloadLoading ?
                                                (
                                                    <svg role="status" className="inline ml-3 w-6 h-6 sm:w-4 sm:h-4 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                        <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB" />
                                                        <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor" />
                                                    </svg>
                                                ) :
                                                this.state.downloadCompleted ?
                                                    (
                                                        this.state.downloadSuccess ?
                                                            (
                                                                // Success Check SVG
                                                                <svg className="inline ml-3 w-4 h-4 text-white" viewBox="0 0 24 24">
                                                                    <path fill="currentColor" d="M20,12A8,8 0 0,1 12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4C12.76,4 13.5,4.11 14.2,4.31L15.77,2.74C14.61,2.26 13.34,2 12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12M7.91,10.08L6.5,11.5L11,16L21,6L19.59,4.58L11,13.17L7.91,10.08Z" />
                                                                </svg>

                                                            ) :
                                                            (
                                                                <svg className="inline ml-3 w-4 h-4 text-white" viewBox="0 0 24 24">
                                                                    <path fill="currentColor" d="M11,15H13V17H11V15M11,7H13V13H11V7M12,2C6.47,2 2,6.5 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4A8,8 0 0,1 20,12A8,8 0 0,1 12,20Z" />
                                                                </svg>
                                                            )
                                                    ) : null
                                        }
                                    </button>
                                </div>
                            ) : null
                    }
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