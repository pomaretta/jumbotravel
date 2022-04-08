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

function FlightProduct(props) {

    let imageSource = "/images/0000-default.svg";
    switch (props.product.product_code) {
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

    // Use state for range slider
    const [range, setRange] = React.useState(0);


    return (
        <div className="flex flex-col | w-full | items-start justify-start | p-1 | space-y-4">
            <div
                className="bg-gray-50 | shadow hover:shadow-md rounded-md | w-full sm:h-12 | flex flex-col sm:flex-row justify-between items-start sm:items-center | p-2 | space-y-2 sm:space-y-0"
            >
                <div
                    className="w-full | flex | justify-between items-center | sm:w-1/3"
                >
                    <div className="flex | justify-center items-center | space-x-4">
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
                                props.product.name
                            }
                        </p>
                    </div>
                    <div className="sm:hidden | flex | justify-end items-center | space-x-2">
                        <p className="text-sm font-bold text-brand-blue">
                            Stock:
                        </p>
                        <input
                            type="number"
                            className="border border-black | px-2 py-1 | rounded-md | bg-white | text-xs | w-10 | text-center"
                            defaultValue={props.product.stock}
                        />
                        <p className="text-sm">
                            <span className="mr-2 font-bold text-brand-blue">Max:</span>
                            {
                                props.product.max
                            }
                        </p>
                    </div>
                </div>
                <div className="flex | w-full | justify-center sm:justify-end items-center | space-x-3">
                    <div
                        className={ClassName(
                            props.product.stock === props.product.max || !props.editing ? "hidden" : "flex",
                            "w-full sm:w-auto justify-end items-center | space-x-3"
                        )}
                    >
                        <input
                            type="range"
                            min={0}
                            max={eval(`${props.product.max} - ${props.product.stock}`)}
                            value={range}
                            onChange={(e) => {
                                setRange(e.target.value);
                            }}
                            className="w-full sm:w-auto"
                        />
                        <p className="text-md sm:text-sm font-bold | text-gray-700">
                            +
                            {range}
                        </p>
                    </div>
                    <div className="hidden sm:flex justify-center sm:justify-end items-center space-x-3">
                        <p className="text-xs mr-2 font-bold text-brand-blue">
                            Stock:
                        </p>
                        <input
                            type="number"
                            className="border border-black | px-2 py-1 | rounded-md | bg-white | text-xs | w-10 | text-center"
                            defaultValue={props.product.stock}
                        />
                        <p className="text-xs">
                            <span className="mr-2 font-bold text-brand-blue">Max:</span>
                            {
                                props.product.max
                            }
                        </p>
                    </div>
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
                            <h4 className="text-2xl sm:text-xl font-bold text-brand-blue">Details</h4>
                            <p
                                className={ClassName(
                                    this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'FLYING' ?
                                        'bg-jt-primary3' :
                                        this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'DEPARTURE' ?
                                            'bg-red-400' :
                                            this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'BUSY' ?
                                                'bg-orange-400' :
                                                'bg-jt-primary',
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
                                                this.context.agentFlightDetails.status === 'FLYING' ?
                                                    'bg-jt-primary3' :
                                                    this.context.agentFlightDetails.status === 'DEPARTURE' ?
                                                        'bg-red-400' :
                                                        this.context.agentFlightDetails.status === 'BUSY' ?
                                                            'bg-orange-400'
                                                            : 'bg-jt-primary',
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
                    <p className="font-regular text-brand-blue">Operations</p>
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
                        <div className="sm:absolute bottom-0 | w-full | flex flex-wrap sm:flex-nowrap items-center justify-between p-3 | bg-white">
                            <div className="flex flex-col sm:flex-row | space-y-2 sm:space-y-0 sm:space-x-4">
                                <button

                                    className={ClassName(
                                        this.context.agentFlightDetails && (this.context.agentFlightDetails.status === 'BUSY' || this.context.agentFlightDetails.active === 0) ? 'hidden' : 'sm:block',
                                        "text-lg sm:text-xs px-4 py-2 | bg-jt-primary | rounded | shadow hover:shadow-md | font-bold | text-white"
                                    )}
                                >
                                    {
                                        this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'DEPARTURE' ?
                                            'Take Off' :
                                            this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'FLYING' ?
                                                'Land' :
                                                this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'ARRIVAL' ?
                                                    'Complete' :
                                                    'Unknown Operation'
                                    }
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

class FlightAgents extends Component {

    render() {
        return (
            <div className="flex flex-col | items-start justify-start | w-full sm:w-3/5 sm:h-full">
                {/* Control */}
                <div className="flex flex-row | items-center justify-center | text-3xl | space-x-3 | p-4">
                    <p className="font-regular text-brand-blue">Agents</p>
                </div>
                {/* Content */}
                <div className="flex flex-col | items-start justify-start | w-full h-full | p-4">
                    <div className="relative sm:bg-white | w-full h-full | rounded-md | sm:shadow | flex | overflow-scroll | no-scrollbar">
                        <div className="sm:absolute top-0 flex flex-col | w-full | space-y-4 sm:space-y-3 | p-1 sm:p-3">
                            {
                                this.context.agentFlightAgents && this.context.agentFlightAgents.agents.length > 0 ?
                                    this.context.agentFlightAgents.agents.map((item, index) => {
                                        return (
                                            <a
                                                key={index}
                                                className="w-full | bg-gray-50 | h-12 | shadow hover:shadow-md | rounded-md | flex flex-row | justify-between items-center | p-2"
                                                href={
                                                    `/agent/${item.getAgentId()}`
                                                }
                                            >
                                                <p className="text-md sm:text-sm">
                                                    {
                                                        item.getFullName()
                                                    }
                                                </p>
                                                <p
                                                    href={`mailto:john@example.com`}
                                                    className="underline text-xs px-4 py-2 bg-jt-primary | rounded | text-white"
                                                >
                                                    {
                                                        item.getEmail()
                                                    }
                                                </p>
                                            </a>
                                        )
                                    })
                                    :
                                    <div className="relative flex flex-col | items-start justify-start | w-full h-full | p-4">
                                        <p className="text-gray-500 text-xl font-bold">No Agents</p>
                                    </div>
                            }
                        </div>

                    </div>
                </div>
            </div>
        )
    }

}
FlightAgents.contextType = Context;

class FlightStock extends Component {

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
                    <p className="font-regular text-brand-blue">Stock</p>
                </div>

                <div className="flex | w-full h-full | p-4">
                    <div className="relative | bg-white | shadow | rounded-md | flex flex-col | items-start justify-between | w-full h-full | overflow-hidden">
                        <div className="sm:absolute | w-full h-full | flex flex-col | items-start justify-start | sm:overflow-scroll | sm:no-scrollbar | sm:pb-16 | p-2 | space-y-4 sm:space-y-0">
                            {/* Flight Operations */}
                            {
                                this.context.agentFlightProducts && this.context.agentFlightProducts.products.length > 0 ?
                                    this.context.agentFlightProducts.products.map((item, index) => {
                                        return (
                                            <FlightProduct key={index} product={item} editing={this.state.isEdit} toggleEdit={this.toggleEdit} />
                                        )
                                    })
                                    :
                                    <div className="relative flex flex-col | items-start justify-start | w-full h-full | p-4">
                                        <p className="text-gray-500 text-xl font-bold">No Flight Products</p>
                                    </div>
                            }
                        </div>
                        <div className="sm:absolute bottom-0 | w-full | flex flex-col sm:flex-row items-start sm:items-center justify-start sm:justify-between | space-y-2 sm:space-y-0 p-3 | bg-white">
                            <div className="flex flex-col w-full sm:w-auto | justify-center items-center | sm:flex-row | space-y-2 sm:space-y-0 sm:space-x-4">
                                <button
                                    className={ClassName(
                                        this.state.isEdit ? "hidden" : "block",
                                        "text-2xl w-full sm:text-xs px-4 py-2 | bg-jt-primary | rounded | shadow hover:shadow-md | font-bold | text-white"
                                    )}
                                    onClick={() => {
                                        this.toggleEdit();
                                    }}
                                >
                                    Fill
                                </button>
                            </div>
                            <div 
                                className={ClassName(
                                    this.state.isEdit ? "flex" : "hidden",
                                    "flex-col w-full sm:w-auto | justify-center items-center sm:flex-row | space-y-2 sm:space-y-0 sm:space-x-2"
                                )}
                            >
                                <button
                                    className="text-2xl w-full sm:w-auto sm:text-xs px-4 py-2 | bg-green-400 | rounded | shadow hover:shadow-md | font-bold | text-white"
                                    onClick={() => {
                                        this.context.pushLocalNotification({
                                            title: "Successfull order",
                                            type: "SUCCESS"
                                        });
                                        this.toggleEdit();
                                    }}
                                >
                                    Place Order
                                </button>
                                <button
                                    className="text-2xl w-full sm:w-auto sm:text-xs px-4 py-2 | bg-red-400 | rounded | shadow hover:shadow-md | font-bold | text-white"
                                    onClick={() => {
                                        this.toggleEdit();
                                    }}
                                >
                                    Clear Order
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        )
    }

}
FlightStock.contextType = Context;

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
                <div className="w-full bg-gray-50 sm:h-1/2 | flex flex-col sm:flex-row | items-center justify-center">
                    <FlightAgents router={this.props.router} />
                    <FlightStock router={this.props.router} />
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
        // Get agents
        this.context.getAgentFlightAgents(this.props.router.params.id);
        // Get products
        this.context.getAgentFlightProducts(this.props.router.params.id);
    }

    componentDidUpdate() {
        if (this.context.hasToLogIn()) {
            this.context.logout();
        }
    }

    render() {

        if (this.context.agentFlightDetails && this.context.agentFlightDetails instanceof APIError && this.context.agentFlightDetails.getStatus() === 404) {
            // Redirect to /flights
            document.location.href = '/flights';
        }

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