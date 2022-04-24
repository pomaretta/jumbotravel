import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Modal from '../../base/modal';
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
                                (parseInt(props.item.items) * parseFloat(props.item.price)).toFixed(2)
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
                                (parseInt(props.item.items) * parseFloat(props.item.price)).toFixed(2)
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

    constructor(props) {
        super(props);

        this.state = {
            requestLoading: false,
            requestCompleted: false,
            requestSuccess: false,
            requestErrorMessage: null,
            createInvoiceLoading: false,
            createInvoiceCompleted: false,
            createInvoiceSuccess: false,
            createInvoiceErrorMessage: null,
            fillLoading: false,
            fillCompleted: false,
            fillSuccess: false,
            fillErrorMessage: null,
        }

        this.requestReview = this.requestReview.bind(this);
        this.createInvoice = this.createInvoice.bind(this);
        this.fillBooking = this.fillBooking.bind(this);
    }

    async requestReview() {

        if (this.state.requestLoading) {
            return;
        }

        this.setState({
            requestLoading: true,
            requestCompleted: false,
            requestSuccess: false,
            requestErrorMessage: null,
        })

        let requestSuccess = false;
        
        await this.context.putBookingRequest(this.props.router.params.id)
            .then(() => {requestSuccess = true})
            .catch(error => {
                requestSuccess = false;
                this.setState({
                    requestErrorMessage: error.statusMessage,
                })
            })
        
        setTimeout(() => {

            this.setState({
                requestLoading: false,
                requestCompleted: true,
                requestSuccess: requestSuccess,
            });

            // TODO: If success play sound, else send notification
            if (requestSuccess) {
                let audio = new Audio('/resources/success.mp3');
                audio.play();
            } else {
                // Local notification with error message
                this.context.pushLocalNotification({
                    title: 'Error requesting review',
                    message: this.state.statusErrorMessage,
                    link: null,
                    extra: null,
                    type: "ERROR"
                });
            }

            setTimeout(() => {
                this.setState({
                    requestLoading: false,
                    requestCompleted: !this.state.requestSuccess,
                    requestSuccess: false,
                })

                if (requestSuccess) {
                    this.props.update();
                }

            }, 3000);

        }, 2500);

    }

    async createInvoice() {

        if (this.state.createInvoiceLoading) {
            return;
        }

        this.setState({
            createInvoiceLoading: true,
            createInvoiceCompleted: false,
            createInvoiceSuccess: false,
            createInvoiceErrorMessage: null,
        })

        let requestSuccess = false;
        let blob = null;

        await this.context.putInvoice(this.props.router.params.id)
            .then((resp) => {
                requestSuccess = true
                blob = resp;
            })
            .catch(error => {
                requestSuccess = false;
                this.setState({
                    createInvoiceErrorMessage: error.statusMessage,
                })
            })
        
        setTimeout(() => {

            this.setState({
                createInvoiceLoading: false,
                createInvoiceCompleted: true,
                createInvoiceSuccess: requestSuccess,
            });

            if (!requestSuccess) {
                // Local notification with error message
                this.context.pushLocalNotification({
                    title: 'Error requesting invoice',
                    message: this.state.createInvoiceErrorMessage,
                    link: null,
                    extra: null,
                    type: "ERROR"
                });
            }

            setTimeout(() => {
                
                this.setState({
                    createInvoiceLoading: false,
                    createInvoiceCompleted: !this.state.createInvoiceSuccess,
                    createInvoiceSuccess: false,
                })

                if (window.navigator && window.navigator.msSaveOrOpenBlob) {
                    window.navigator.msSaveOrOpenBlob(blob);
                } else {
                    const data = window.URL.createObjectURL(blob);
                    var link = document.createElement('a');
                    link.href = data;
                    link.download = `invoice_${this.props.router.params.id}_${new Date().getTime()}.pdf`;
                    link.click();
                    setTimeout(function(){
                        // For Firefox it is necessary to delay revoking the ObjectURL
                        window.URL.revokeObjectURL(data);
                    }, 100);
                }
            }, 3000);
        }, 2500);

    }

    async fillBooking() {

        if (this.state.fillLoading) {
            return;
        }

        this.setState({
            fillLoading: true,
            fillCompleted: false,
            fillSuccess: false,
            fillErrorMessage: null,
        })

        let requestSuccess = false;

        await this.context.fillBooking(this.props.router.params.id)
            .then(() => {
                requestSuccess = true
            })
            .catch(error => {
                requestSuccess = false;
                this.setState({
                    fillErrorMessage: error.message,
                })
            })
        
        setTimeout(() => {

            this.setState({
                fillLoading: false,
                fillCompleted: true,
                fillSuccess: requestSuccess,
            });

            if (requestSuccess) {
                let audio = new Audio('/resources/success.mp3');
                audio.play();
            } else {
                // Local notification with error message
                this.context.pushLocalNotification({
                    title: 'Error completing booking',
                    message: this.state.fillErrorMessage,
                    link: null,
                    extra: null,
                    type: "ERROR"
                });
            }

            setTimeout(() => {

                this.setState({
                    fillLoading: false,
                    fillCompleted: !this.state.fillSuccess,
                    fillSuccess: false,
                    fillErrorMessage: null,
                })

                if (requestSuccess) {
                    this.props.update();
                }

            }, 3000);

        }, 2500);

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
                            <div className={ClassName(
                                "w-full | flex flex-col sm:flex-row | space-y-2 sm:space-y-0",
                                this.context.agentBookingDetails && this.context.agentBookingDetails.status !== "COMPLETED" && this.context.agent && this.context.agent.type != "PROVIDER" ? "sm:space-x-4" : ""
                            )}>
                                <button
                                    className={ClassName(
                                        this.context.agentBookingDetails && this.context.agentBookingDetails.status !== "COMPLETED" && this.context.agent && this.context.agent.type != "PROVIDER" ? "block" : "hidden",
                                        "btn btn-md sm:btn-sm btn-success | w-full | sm:w-auto text-white",
                                        this.state.requestCompleted ? this.state.requestSuccess ? 'btn-success' : 'btn-error' : 'bg-jt-primary',
                                    )}
                                    onClick={this.requestReview}
                                >
                                    Request Review
                                    {
                                        this.state.requestLoading ?
                                            (
                                                <svg role="status" className="inline ml-3 w-6 h-6 sm:w-4 sm:h-4 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                    <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB" />
                                                    <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor" />
                                                </svg>
                                            ) :
                                            this.state.requestCompleted ?
                                                (
                                                    this.state.requestSuccess ?
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
                                        this.context.agentBookingDetails && this.context.agentBookingDetails.status == "COMPLETED" && this.context.agent && this.context.agent.type != "PROVIDER" ? "block" : "hidden",
                                        "btn btn-md sm:btn-sm btn-success | w-full | sm:w-auto text-white",
                                        this.state.createInvoiceCompleted ? this.state.createInvoiceSuccess ? 'btn-success' : 'btn-error' : 'bg-jt-primary',
                                    )}
                                    onClick={this.createInvoice}
                                >
                                    {
                                        this.context.agentBookingDetails && this.context.agentBookingDetails.has_invoice ? "View Invoice" : "Create Invoice"
                                    }
                                    {
                                        this.state.createInvoiceLoading ?
                                            (
                                                <svg role="status" className="inline ml-3 w-6 h-6 sm:w-4 sm:h-4 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                    <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB" />
                                                    <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor" />
                                                </svg>
                                            ) :
                                            this.state.createInvoiceCompleted ?
                                                (
                                                    this.state.createInvoiceSuccess ?
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
                                        this.context.agentBookingDetails && this.context.agentBookingDetails.status != "COMPLETED" && this.context.agent && this.context.agent.type == "PROVIDER" ? "block" : "hidden",
                                        "btn btn-md sm:btn-sm btn-success | w-full | sm:w-auto text-white",
                                        this.state.fillCompleted ? this.state.fillSuccess ? 'btn-success' : 'btn-error' : 'bg-jt-primary',
                                    )}
                                    onClick={this.fillBooking}
                                >
                                    Complete Order
                                    {
                                        this.state.fillLoading ?
                                            (
                                                <svg role="status" className="inline ml-3 w-6 h-6 sm:w-4 sm:h-4 text-white animate-spin" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                                                    <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="#E5E7EB" />
                                                    <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentColor" />
                                                </svg>
                                            ) :
                                            this.state.fillCompleted ?
                                                (
                                                    this.state.fillSuccess ?
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
                                                `${parseFloat(this.context.agentBookingDetails.total).toFixed(2)}€` :
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
                    <BookingOperations router={this.props.router} update={this.props.update} />
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

    constructor(props) {
        super(props);

        this.state = {
            interval: null
        };

        this.updateDashboard = this.updateDashboard.bind(this);
    }

    componentDidMount() {
        // Get details
        this.context.getAgentBookingDetails(this.props.router.params.id);
        // Get operations
        this.context.getAgentBookingOperations(this.props.router.params.id);
        // Get items
        this.context.getAgentBookingItems(this.props.router.params.id);


        // Set interval
        this.setState({
            interval: setInterval(() => {
                // Get details
                this.context.getAgentBookingDetails(this.props.router.params.id);
                // Get operations
                this.context.getAgentBookingOperations(this.props.router.params.id);
                // Get items
                this.context.getAgentBookingItems(this.props.router.params.id);
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
        this.context.getAgentBookingDetails(this.props.router.params.id);
        // Get operations
        this.context.getAgentBookingOperations(this.props.router.params.id);
        // Get items
        this.context.getAgentBookingItems(this.props.router.params.id);
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
                        Booking {String(this.props.router.params.id).substring(0, 8)} - JumboTravel
                    </title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={2} />
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