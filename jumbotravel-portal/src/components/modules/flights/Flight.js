import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Modal from "../../base/modal";
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
                        {
                            props.notification.extra && props.notification.extra["booking"] ?
                                <a
                                    href={`/bookings/${props.notification.extra["booking"]}`}
                                    className="text-brand-blue underline | ml-2"
                                >
                                    {
                                        props.notification.extra["booking"]
                                    }
                                </a> :
                                null
                        }
                    </p>
                </div>

                <div className="mt-3 sm:mt-0 w-1/3">
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
    // const [range, setRange] = React.useState(0);


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
                            value={
                                props.value ? props.value : 0
                            }
                            onChange={(e) => {
                                props.changeValue(props.product.product_code, e.target.value);
                            }}
                            className="w-full sm:w-auto | range range-xs"
                        />
                        <p className="text-md sm:text-sm font-bold | text-gray-700">
                            +
                            {
                                props.value ? props.value : 0
                            }
                        </p>
                    </div>
                    <div className="hidden sm:flex justify-center sm:justify-end items-center space-x-3">
                        <p className="text-xs mr-2 font-bold text-brand-blue">
                            Stock:
                        </p>
                        <p
                            className="border border-black | px-2 py-1 | rounded-md | bg-white | text-xs | w-10 | text-center"
                        >
                            {props.product.stock}
                        </p>
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

    constructor(props) {
        super(props);

        this.state = {
            statusLoading: false,
            statusCompleted: false,
            statusSuccess: false,
            statusErrorMessage: null,
        }
    }

    async changeStatusRequest() {

        if (this.state.statusLoading) {
            return;
        }

        this.setState({
            statusLoading: true,
            statusCompleted: false,
            statusSuccess: false,
            statusErrorMessage: null,
        })

        let requestSuccess = false;

        // TODO: If the status is ARRIVAL and there's no booking associated, then we should ask the user if they want to complete the flight.
        let canBeCreated = true;
        if (this.context.agentFlightDetails.status === 'ARRIVAL' && !this.context.agentFlightDetails.has_booking) {
            await this.context.createModal({
                data: {
                    type: 'actiondelete',
                    title: 'This flight does not have booking created. Are you sure you want to complete this flight?',
                }
            })
                .then(event => {
                    if (event.proceed && event.proceed === false) {
                        canBeCreated = false;
                    }
                })
                .catch(() => {
                    canBeCreated = false;
                })
        }
        if (!canBeCreated) {
            this.setState({
                statusLoading: false,
                statusCompleted: false,
                statusSuccess: false,
                statusErrorMessage: null,
            })
            return;
        }

        // TODO: Make request using context
        await this.context.updateFlightStatus(this.props.router.params.id)
            .then(() => {
                requestSuccess = true;
            })
            .catch(error => {
                requestSuccess = false;
                this.setState({
                    statusErrorMessage: error.message,
                })
            });

        // Visual feedback
        setTimeout(() => {

            this.setState({
                statusLoading: false,
                statusCompleted: true,
                statusSuccess: requestSuccess,
            })

            // Play sound effect
            if (requestSuccess) {
                let audio = new Audio('/resources/success.mp3');
                audio.play();
            } else {
                // Local notification with error message
                this.context.pushLocalNotification({
                    title: 'Error updating flight status',
                    message: this.state.statusErrorMessage,
                    link: null,
                    extra: null,
                    type: "ERROR"
                });
            }

            // Reset state
            setTimeout(() => {
                this.setState({
                    statusLoading: false,
                    statusCompleted: !this.state.statusSuccess,
                    statusSuccess: false,
                })

                if (requestSuccess) {
                    this.props.update();
                }

            }, 3500)

        }, 2500)

    }

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
                            <div className="flex flex-col sm:flex-row | w-full | space-y-2 sm:space-y-0 sm:space-x-4">
                                <button
                                    className={ClassName(
                                        this.context.agentFlightDetails && (this.context.agentFlightDetails.status === 'COMPLETED' || this.context.agentFlightDetails.active === 0) ? 'hidden' : 'sm:block',
                                        "btn btn-md sm:btn-sm btn-success | w-full sm:w-auto border-0 text-lg sm:text-xs px-4 py-2 | rounded | shadow hover:shadow-md | font-bold | text-white | flex | items-center | justify-center",
                                        this.state.statusCompleted ? this.state.statusSuccess ? 'btn-success' : 'btn-error' : 'bg-jt-primary',
                                    )}
                                    onClick={() => this.changeStatusRequest()}
                                >
                                    {
                                        this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'BUSY' ?
                                            'Init Flight' :
                                            this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'DEPARTURE' ?
                                                'Take Off' :
                                                this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'FLYING' ?
                                                    'Land' :
                                                    this.context.agentFlightDetails && this.context.agentFlightDetails.status === 'ARRIVAL' ?
                                                        'Complete' :
                                                        'Unknown Operation'
                                    }
                                    {
                                        this.state.statusLoading ?
                                            (
                                                <svg role="status" className="inline ml-3 w-4 h-4 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                    <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB" />
                                                    <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor" />
                                                </svg>
                                            ) :
                                            this.state.statusCompleted ?
                                                (
                                                    this.state.statusSuccess ?
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
                                                )
                                                :
                                                ''
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
                                                    `/agents/${item.getAgentId()}`
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
            bookingLoading: false,
            bookingCompleted: false,
            bookingSuccess: false,
            bookingErrorMessage: null,
            products: {},
        }

        this.toggleEdit.bind(this);
        this.changeProductValue = this.changeProductValue.bind(this);
        this.clearOrder = this.clearOrder.bind(this);
    }

    toggleEdit() {

        if (this.context.agentFlightDetails && this.context.agentFlightDetails.has_booking) {
            return;
        }

        this.setState({
            isEdit: !this.state.isEdit
        })
    }

    async createBooking() {

        if (this.state.bookingLoading || (this.context.agentFlightDetails && this.context.agentFlightDetails.has_booking)) {
            return;
        }

        this.setState({
            bookingLoading: true,
            bookingCompleted: false,
            bookingSuccess: false,
            bookingErrorMessage: null,
        })

        let requestSuccess = false;

        // TODO: Make request using the API
        let products = this.state.products;
        let flightId = this.context.agentFlightDetails.flight_id;

        // If there are no products, return and show error
        if (Object.keys(products).length === 0) {
            this.context.pushLocalNotification({
                title: 'Error creating booking',
                message: 'No products selected',
                type: 'ERROR',
            })
            this.setState({
                bookingLoading: false,
                bookingCompleted: false,
                bookingSuccess: false,
                bookingErrorMessage: null,
            })
            return;
        }

        let items = Object.keys(products).map((productCode) => {
            return {
                product_code: parseInt(productCode),
                quantity: products[productCode],
            }
        })

        // TODO: Modal to confirm booking
        let canBeCreated = false;
        await this.context.createModal({
            data: {
                type: 'actionplaceorder',
                items: items,
                products: this.context.agentFlightProducts ? this.context.agentFlightProducts.products : null
            }
        })
            .then(event => {
                if (event.proceed && event.proceed === true) {
                    canBeCreated = true;
                }
            })
            .catch(() => {
                canBeCreated = false;
            })

        if (!canBeCreated) {
            this.setState({
                bookingLoading: false,
                bookingCompleted: false,
                bookingSuccess: false,
                bookingErrorMessage: null,
            });
            return;
        }

        await this.context.putBookingOrder(flightId, items)
            .then(() => {
                requestSuccess = true;
            })
            .catch(error => {
                requestSuccess = false;
                this.setState({
                    bookingErrorMessage: error.statusMessage,
                })
            })

        setTimeout(() => {

            this.setState({
                bookingLoading: false,
                bookingCompleted: true,
                bookingSuccess: requestSuccess,
            });

            // TODO: If success play sound, else send notification
            if (requestSuccess) {
                let audio = new Audio('/resources/success.mp3');
                audio.play();
            } else {
                // Local notification with error message
                this.context.pushLocalNotification({
                    title: 'Error placing order',
                    message: this.state.statusErrorMessage,
                    link: null,
                    extra: null,
                    type: "ERROR"
                });
            }

            setTimeout(() => {

                this.setState({
                    bookingLoading: false,
                    bookingCompleted: !this.state.bookingSuccess,
                    bookingSuccess: false,
                })

                if (requestSuccess) {
                    this.props.update();
                }

            }, 3000);

        }, 2500);

    }

    clearOrder = () => {

        if (this.state.bookingLoading) {
            return;
        }

        this.setState({
            products: {},
        });
    }

    changeProductValue = (productCode, value) => {
        let products = this.state.products;
        value = parseInt(value);
        products[productCode] = value;
        // If the new value is 0, remove the product
        if (value === 0) {
            delete products[productCode];
        }
        this.setState({
            products: products
        });
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
                                this.context.agentFlightProducts && this.context.agentFlightProducts.products.length > 0 && (this.context.agentFlightDetails && this.context.agentFlightDetails.status !== "BUSY") ?
                                    this.context.agentFlightProducts.products.map((item, index) => {
                                        return (
                                            <FlightProduct key={index} product={item} editing={this.state.isEdit} toggleEdit={this.toggleEdit} changeValue={this.changeProductValue} value={this.state.products[item.product_code]} />
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
                                        this.state.isEdit || !(this.context.agentFlightDetails && this.context.agentFlightDetails.status !== "BUSY" && this.context.agentFlightDetails.status !== "COMPLETED" && !this.context.agentFlightDetails.has_booking) ? "hidden" : "block",
                                        "text-2xl uppercase w-full sm:text-xs px-4 py-2 | bg-jt-primary | rounded | shadow hover:shadow-md | font-bold | text-white"
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
                                    className={ClassName(
                                        "btn btn-md sm:btn-sm btn-success | text-2xl w-full sm:w-auto sm:text-xs px-4 py-2 | rounded | shadow hover:shadow-md | font-bold | text-white",
                                        this.state.bookingCompleted ? this.state.bookingSuccess ? 'btn-success' : 'btn-error' : 'bg-jt-primary',
                                    )}
                                    onClick={() => this.createBooking()}
                                >
                                    Place Order
                                    {
                                        this.state.bookingLoading ?
                                            (
                                                <svg role="status" className="inline ml-3 w-6 h-6 sm:w-4 sm:h-4 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                    <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB" />
                                                    <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor" />
                                                </svg>
                                            ) :
                                            this.state.bookingCompleted ?
                                                (
                                                    this.state.bookingSuccess ?
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
                                                )
                                                :
                                                ''
                                    }
                                </button>
                                <button
                                    className={ClassName(
                                        this.state.bookingLoading ? 'bg-gray-300' : 'btn btn-md sm:btn-sm border-0 btn-error',
                                        this.state.bookingLoading ? 'cursor-not-allowed' : '',
                                        "uppercase | text-2xl w-full sm:w-auto sm:text-xs px-4 py-2 | bg-red-400 | rounded | shadow hover:shadow-md | font-bold | text-white",
                                    )}
                                    onClick={() => this.clearOrder()}
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
                    <FlightOperations router={this.props.router} update={this.props.update} />
                </div>

                {/* Bottom */}
                <div className="w-full bg-gray-50 sm:h-1/2 | flex flex-col sm:flex-row | items-center justify-center">
                    <FlightAgents router={this.props.router} />
                    <FlightStock router={this.props.router} update={this.props.update} />
                </div>

            </div >
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

        this.updateDashboard = this.updateDashboard.bind(this);
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

        // Set interval
        this.setState({
            interval: setInterval(() => {
                // Get details
                this.context.getAgentFlightDetails(this.props.router.params.id);
                // Get operations
                this.context.getAgentFlightOperations(this.props.router.params.id);
            }, 10000)
        });
    }

    componentWillUnmount() {
        clearInterval(this.state.interval);
    }

    componentDidUpdate() {
        if (this.context.hasToLogIn()) {
            this.context.logout();
        }
    }

    updateDashboard = async () => {
        // Get details
        this.context.getAgentFlightDetails(this.props.router.params.id);
        // Get operations
        this.context.getAgentFlightOperations(this.props.router.params.id);
        // Get products
        this.context.getAgentFlightProducts(this.props.router.params.id);
    }

    render() {

        // If the user is not assistant, redirect to dashboard
        if (this.context.agent && this.context.agent.type !== "ASSISTANT") {
            document.location.href = "/";
        }

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
                <Modal />
                <div className="relative w-full h-full | flex flex-col justify-start items-start">
                    {/* NavBar */}
                    <NavBar app={this.props.app} config={this.props.config} />
                    {/* Content */}
                    <Content router={this.props.router} update={this.updateDashboard} />
                </div>
            </div>
        );
    }

}

Module.contextType = Context;

export default withRouter(Module);