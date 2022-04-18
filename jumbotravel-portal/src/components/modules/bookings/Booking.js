import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';
import withRouter from '../../utils/router';
import ClassName from '../../utils/classname';

import { getSVG, getOutlineBackground } from '../../utils/types';
import APIError from "../../../api/error";

function Detail(props) {
    return (
        <div className="w-full sm:w-1/2">
            <p className="text-gray-500 text-xl sm:text-xs font-bold | mb-2 sm:mb-0">{props.title}</p>
            {
                props.isLink ?
                <a
                    href={props.href ? props.href : ""}
                    target={props.target ? props.target : "_blank"}
                    className="text-brand-blue text-xl sm:text-xs | underline"
                >{props.content}</a> :
                <p className="text-brand-blue text-xl sm:text-xs">{props.content}</p>
            }
        </div>
    )
}

function BookingOperation(props) {
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

function BookingItem(props) {

    let imageSource = "/images/0000-default.svg";
    switch (parseInt(props.item.productcode)) {
        case 1:
            imageSource = "/images/0001-coca-cola.png";
            break;
        case 2:
            imageSource = "/images/0002-agua.png";
            break;
        case 3:
            imageSource = "/images/0003-coca-cola-light.webp";
            break;
        case 4:
            imageSource = "/images/0004-wine.jpeg";
            break;
        case 5:
            imageSource = "/images/0005-lays.png";
            break;
        case 6:
            imageSource = "/images/0006-aceitunas.jpeg";
            break;
        case 7:
            imageSource = "/images/0007-fanta-naranja.png";
            break;
        case 8:
            imageSource = "/images/0008-fanta-limon.jpeg";
            break;
        case 9:
            imageSource = "/images/0009-frutos-secos.jpeg";
            break;
        case 10:
            imageSource = "/images/0010-galletas.jpeg";
            break;
        case 11:
            imageSource = "/images/0011-cerveza.jpeg";
            break;
        default:
            imageSource = "/images/0000-default.svg";
            break;
    }

    return (
        <div className="flex flex-col | w-full | items-start justify-start | p-1 | space-y-4">
            <div
                className="bg-gray-50 | shadow hover:shadow-md rounded-md | w-full sm:h-12 | flex flex-col sm:flex-row justify-between items-start sm:items-center | p-2 | space-y-2 sm:space-y-0"
            >
                <div
                    className="w-full | flex flex-col sm:flex-row | justify-between items-start sm:items-center | sm:w-1/3"
                >
                    <div className="flex | w-full sm:w-auto | justify-between sm:justify-center items-center | space-x-4">
                        <img
                            className="rounded-md"
                            style={{
                                width: "35px",
                                height: "35px",
                            }}
                            src={imageSource}
                        />
                        <p className="text-xl sm:text-xs | font-bold | text-brand-blue">
                            {
                                props.item.name
                            }
                        </p>
                    </div>
                    <div className="sm:hidden | w-full | flex | justify-between sm:justify-end items-center | space-x-2">
                        <p className="text-sm font-bold text-brand-blue">
                            Units:
                        </p>
                        <input
                            type="number"
                            className="border border-black | px-2 py-1 | rounded-md | bg-white | text-xs | w-10 | text-center"
                            defaultValue={props.item.items}
                        />
                        <p className="text-sm">
                            <span className="mr-2 font-bold text-brand-blue">Price:</span>
                            {
                                props.item.price
                            }€
                        </p>
                        <p className="text-sm">
                            <span className="mr-2 font-bold text-brand-blue">Total:</span>
                            {
                                (parseInt(props.item.items) * parseFloat(props.item.price))
                            }€
                        </p>
                    </div>
                </div>
                <div className="flex | w-full | justify-center sm:justify-end items-center | space-x-3">
                    <div className="hidden sm:flex justify-center sm:justify-end items-center space-x-3">
                        <p className="text-xs mr-2 font-bold text-brand-blue">
                            Units:
                        </p>
                        <p
                            className="border border-black | px-2 py-1 | rounded-md | bg-white | text-xs | w-10 | text-center"
                        >
                            {props.item.items}
                        </p>
                        <p className="text-xs">
                            <span className="mr-2 font-bold text-brand-blue">Price:</span>
                            {
                                props.item.price
                            }€
                        </p>
                        <p className="text-xs">
                            <span className="mr-2 font-bold text-brand-blue">Total:</span>
                            {
                                (parseInt(props.item.items) * parseFloat(props.item.price))
                            }€
                        </p>
                    </div>
                </div>
            </div>
        </div>
    );
}

class BookingDetails extends Component {

    render() {
        return (
            <div className="flex flex-col | items-start justify-start | w-full sm:w-3/5 sm:h-full">
                {/* Control */}
                <div className="flex flex-row | items-center justify-center | text-3xl | space-x-3 | p-4">
                    <a
                        href="/bookings"
                        className="text-gray-500"
                    >Bookings</a>
                    <span className="text-gray-500">\</span>
                    <p className="font-bold text-brand-blue">{String(this.props.router.params.id).substring(0, 8)}</p>
                </div>
                {/* Content */}
                <div className="flex flex-col | items-start justify-start | w-full h-full | p-4">
                    <div className="bg-white | w-full h-full | rounded-md | shadow p-4">
                        <div className="mb-4 | flex items-center justify-between">
                            <h4 className="text-2xl sm:text-xl font-bold text-brand-blue">Details</h4>
                        </div>
                        {
                            this.context.agentBookingDetails ?
                                (
                                    <div className="flex flex-col sm:flex-row flex-wrap | items-start justify-start | w-full space-y-2 sm:space-y-0">
                                        <Detail title="Booking Reference ID" content={this.context.agentBookingDetails.bookingreferenceid} />
                                        <Detail 
                                            title="Flight" 
                                            content={this.context.agentBookingDetails.flight_id} 
                                            isLink={true}
                                            href={`/flights/${this.context.agentBookingDetails.flight_id}`}
                                            target="_self"
                                        />
                                        <Detail
                                            title="Agent"
                                            content={`${this.context.agentBookingDetails.agent_name} ${this.context.agentBookingDetails.agent_surname}`}
                                            isLink={true}
                                            href={`/agents/${this.context.agentBookingDetails.agent_id}`}
                                            target="_self"
                                        />
                                        <Detail
                                            title="Provider"
                                            content={
                                                this.context.agentBookingDetails.provider_id ?
                                                    `${this.context.agentBookingDetails.provider_name} ${this.context.agentBookingDetails.provider_surname}`
                                                    : ""
                                            }
                                            isLink={this.context.agentBookingDetails.provider_id ? true : false}
                                            href={this.context.agentBookingDetails.provider_id ? `/agents/${this.context.agentBookingDetails.provider_id}` : ""}
                                            target="_self"
                                        />
                                        <Detail title="Items" content={this.context.agentBookingDetails.items} />
                                        <Detail title="Total" content={`${this.context.agentBookingDetails.total}€`} />
                                        <Detail title="Creation" content={this.context.agentBookingDetails.created_at} />
                                        <Detail title="Status" content={this.context.agentBookingDetails.status} />
                                    </div>
                                ) :
                                (
                                    <div className="flex flex-col sm:flex-row flex-wrap | items-start justify-start | w-full space-y-2 sm:space-y-0">
                                    </div>
                                )
                        }
                    </div>
                </div>
            </div>
        )
    }

}
BookingDetails.contextType = Context;

class BookingOperations extends Component {

    render() {
        return (
            <div className="relative | flex flex-col | items-start justify-start | w-full sm:h-full">

                {/* Control */}
                <div className="flex flex-row | items-center justify-center | text-3xl | space-x-3 | p-4">
                    <p className="font-regular text-brand-blue">Operations</p>
                </div>

                <div className="flex | w-full h-full | p-4">
                    <div className="relative | bg-white | shadow | rounded-md | flex flex-col | items-start justify-between | w-full h-full | overflow-hidden">
                        <div className="sm:absolute | w-full h-full | flex flex-col | items-start justify-start | sm:overflow-scroll | sm:no-scrollbar | sm:pb-16 | p-2">
                            {/* Flight Operations */}
                            {
                                this.context.agentBookingOperations && this.context.agentBookingOperations.notifications.length > 0 ?
                                    this.context.agentBookingOperations.notifications.map((item, index) => {
                                        return (
                                            <BookingOperation key={index} notification={item} />
                                        )
                                    })
                                    :
                                    <div className="relative flex flex-col | items-start justify-start | w-full h-full | p-4">
                                        <p className="text-gray-500 text-xl font-bold">No Booking Operations</p>
                                    </div>
                            }
                        </div>
                        <div className="sm:absolute bottom-0 | w-full | flex flex-wrap sm:flex-nowrap items-center justify-between p-3 | bg-white">
                            <div className="w-full | flex flex-col sm:flex-row | space-y-2 sm:space-y-0 sm:space-x-4">
                                <button
                                    className={ClassName(
                                        this.context.agentBookingDetails && this.context.agentBookingDetails.status !== "COMPLETED" ? "block" : "hidden",
                                        "w-full text-2xl | sm:w-auto bg-jt-primary text-white | sm:text-sm font-bold py-2 px-4 rounded-md | hover:shadow-sm"
                                    )}
                                >
                                    Request Review
                                </button>
                                <button
                                    className={ClassName(
                                        this.context.agentBookingDetails && this.context.agentBookingDetails.status === "COMPLETED" ? "block" : "hidden",
                                        "w-full text-2xl | sm:w-auto bg-jt-primary text-white | sm:text-sm font-bold py-2 px-4 rounded-md | hover:shadow-sm"
                                    )}
                                >
                                    Generate Invoice
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        )
    }

}
BookingOperations.contextType = Context;

class BookingItems extends Component {

    constructor(props) {
        super(props);

        this.state = {
            isEdit: false,
        }

        this.toggleEdit.bind(this);
    }

    toggleEdit() {
        this.setState({
            isEdit: !this.state.isEdit
        })
    }

    render() {
        return (
            <div className="relative | flex flex-col | items-start justify-start | w-full sm:h-full">

                {/* Control */}
                <div className="flex flex-row | items-center justify-center | text-3xl | space-x-3 | p-4">
                    <p className="font-regular text-brand-blue">Items</p>
                </div>

                <div className="flex | w-full h-full | p-4">
                    <div className="relative | bg-white | shadow | rounded-md | flex flex-col | items-start justify-between | w-full h-full | overflow-hidden">
                        <div className="sm:absolute | w-full h-full | flex flex-col | items-start justify-start | sm:overflow-scroll | sm:no-scrollbar | sm:pb-16 | p-2 | space-y-4 sm:space-y-0">
                            {/* Flight Operations */}
                            {
                                this.context.agentBookingItems && this.context.agentBookingItems.items.length > 0 ?
                                    this.context.agentBookingItems.items.map((item, index) => {
                                        return (
                                            <BookingItem key={index} item={item} />
                                        )
                                    })
                                    :
                                    <div className="relative flex flex-col | items-start justify-start | w-full h-full | p-4">
                                        <p className="text-gray-500 text-xl font-bold">No Booking Items</p>
                                    </div>
                            }
                        </div>
                        <div className="sm:absolute bottom-0 | w-full | flex flex-col sm:flex-row items-start sm:items-center justify-start sm:justify-between | space-y-2 sm:space-y-0 p-3 | bg-white">
                            <div className="flex flex-col w-full sm:w-auto | justify-center items-center | sm:flex-row | space-y-2 sm:space-y-0 sm:space-x-4">
                            </div>
                            <div
                                className={ClassName(
                                    "flex | flex-col w-full sm:w-auto | justify-center items-center sm:flex-row | space-y-2 sm:space-y-0 sm:space-x-2"
                                )}
                            >
                                <p
                                    className="text-xl w-full sm:w-auto sm:text-sm px-4 py-2 | bg-green-400 | rounded | shadow hover:shadow-md | font-bold | text-white"
                                >
                                    Total:
                                    <span className="ml-2">
                                        {
                                            this.context.agentBookingDetails ?
                                                `${this.context.agentBookingDetails.total}€` :
                                                "0€"
                                        }
                                    </span>
                                </p>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        )
    }

}
BookingItems.contextType = Context;

class Content extends Component {

    render() {
        return (
            <div className="relative | flex flex-col | w-full h-full | sm:overflow-scroll | sm:no-scrollbar | bg-gray-50">
                {/* Top */}
                <div className="w-full bg-gray-50 sm:h-1/2 | flex flex-col sm:flex-row | items-center justify-center">
                    <BookingDetails router={this.props.router} />
                    <BookingOperations router={this.props.router} />
                </div>

                {/* Bottom */}
                <div className="w-full bg-gray-50 sm:h-1/2 | flex flex-col sm:flex-row | items-center justify-center">
                    <BookingItems router={this.props.router} />
                </div>

            </div >
        )
    }

}

Content.contextType = Context;

class Module extends Component {

    componentDidMount() {
        // Get details
        this.context.getAgentBookingDetails(this.props.router.params.id);
        // Get operations
        this.context.getAgentBookingOperations(this.props.router.params.id);
        // Get items
        this.context.getAgentBookingItems(this.props.router.params.id);
    }

    componentDidUpdate() {
        if (this.context.hasToLogIn()) {
            this.context.logout();
        }
    }

    render() {

        if (this.context.agentBookingDetails && this.context.agentBookingDetails instanceof APIError && this.context.agentBookingDetails.getStatus() === 404) {
            // Redirect to /flights
            document.location.href = '/bookings';
        }

        return (
            <div className="w-screen h-screen flex flex-row flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800">
                <Helmet>
                    <title>
                        Booking - JumboTravel
                    </title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={2} />
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