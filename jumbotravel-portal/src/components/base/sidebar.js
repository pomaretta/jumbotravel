import { Component } from "react";
import { Link } from "react-router-dom";

import Context from '../context/app';

function classNames(...classes) {
    return classes.filter(Boolean).join(' ')
}

class Sidebar extends Component {

    constructor(props) {
        super(props);
        this.state = {
            isOpen: false,
        }
    }

    render() {
        return (
            <div className="hidden sm:flex flex-col top-0 left-0 w-16 bg-white h-full border-r z-10">
                <div className="flex items-center justify-center h-14 border-b">
                    {/* Image */}
                    <img className="h-8 w-auto" src="/resources/logo.svg" alt="Workflow" />
                </div>
                <div className="flex-grow">
                    <ul className="flex flex-col py-4 space-y-1 h-full">
                        <li>
                            <Link to="/dashboard" className={
                                classNames(
                                    this.props.current ? this.props.current === 3 ? "border-jt-primary" : "" : "",
                                    "relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-jt-primary pr-6"
                                )
                            }>
                                <span className="inline-flex justify-center items-center ml-4">
                                    <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path fill="currentColor" d="M13,3V9H21V3M13,21H21V11H13M3,21H11V15H3M3,13H11V3H3V13Z" />
                                    </svg>
                                </span>
                            </Link>
                        </li>
                        <li className={classNames(
                            this.context.agent && this.context.agent.type === "ASSISTANT" ? "" : "hidden",
                        )}>
                            <Link to="/flights" className={
                                classNames(
                                    this.props.current ? this.props.current === 1 ? "border-jt-primary" : "" : "",
                                    "relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-jt-primary pr-6"
                                )
                            }>
                                <span className="inline-flex justify-center items-center ml-4">
                                    <svg className="w-5 h-5" fill="currentColor" stroke="" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                        <path d="M2.5,19H21.5V21H2.5V19M9.68,13.27L14.03,14.43L19.34,15.85C20.14,16.06 20.96,15.59 21.18,14.79C21.39,14 20.92,13.17 20.12,12.95L14.81,11.53L12.05,2.5L10.12,2V10.28L5.15,8.95L4.22,6.63L2.77,6.24V11.41L4.37,11.84L9.68,13.27Z" />
                                    </svg>
                                </span>
                            </Link>
                        </li>
                        <li>
                            <Link to="/bookings" className={
                                classNames(
                                    this.props.current ? this.props.current === 2 ? "border-jt-primary" : "" : "",
                                    "relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-jt-primary pr-6"
                                )
                            }>
                                <span className="inline-flex justify-center items-center ml-4">
                                    <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path fill="currentColor" d="M4,2C2.89,2 2,2.89 2,4V14H4V4H14V2H4M8,6C6.89,6 6,6.89 6,8V18H8V8H18V6H8M12,10C10.89,10 10,10.89 10,12V20C10,21.11 10.89,22 12,22H20C21.11,22 22,21.11 22,20V12C22,10.89 21.11,10 20,10H12Z" />
                                    </svg>
                                </span>
                            </Link>
                        </li>
                        {/* <li>
                            <Link to="/planes" className={
                                classNames(
                                    this.props.current ? this.props.current === 3 ? "border-jt-primary" : "" : "",
                                    "relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-jt-primary pr-6"
                                )
                            }>
                                <span className="inline-flex justify-center items-center ml-4">
                                    <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                        <path fill="currentColor" d="M20.56 3.91C21.15 4.5 21.15 5.45 20.56 6.03L16.67 9.92L18.79 19.11L17.38 20.53L13.5 13.1L9.6 17L9.96 19.47L8.89 20.53L7.13 17.35L3.94 15.58L5 14.5L7.5 14.87L11.37 11L3.94 7.09L5.36 5.68L14.55 7.8L18.44 3.91C19 3.33 20 3.33 20.56 3.91Z" />
                                    </svg>
                                </span>
                            </Link>
                        </li> */}
                        <div className="mt-auto-important">
                            <li
                                className={classNames(
                                    this.context.newNotifications ? "border-red-500" : "",
                                    "relative cursor-pointer flex flex-row items-center justify-end h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-jt-primary pr-6"
                                )}

                                onClick={() => {
                                    this.context.setNotificationsOpen()
                                }}
                            >
                                <lord-icon
                                    trigger={this.context.newNotifications ? "loop" : ""}
                                    src="/resources/notification-bell.json"
                                    style={{
                                        width: '24px',
                                        height: '24px',
                                    }}
                                />
                            </li>
                            <li>
                                <a onClick={() => {
                                    // this.props.app.logout()
                                    this.context.logout()
                                }} className="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-red-500 border-l-4 border-transparent hover:border-jt-primary pr-6">
                                    <span className="inline-flex justify-center items-center ml-4">
                                        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path></svg>
                                    </span>
                                </a>
                            </li>
                        </div>
                        {/* Notification */}
                        <div
                            className="absolute w-1/4 h-1/4 | bottom-3 left-20 | bg-white | rounded-md | shadow | flex items-center flex-col justify-center"
                            style={{
                                display: this.context.notificationsIsOpen ? "flex" : "none"
                            }}
                        >
                            <div className="w-full h-full flex items-center flex-col justify-between">
                                <div className="w-full h-full | flex items-center flex-col justify-start | p-2 | overflow-scroll | no-scrollbar">
                                    {
                                        this.context.hasNotifications ?
                                            this.context.notifications.getNotPopup().map((notification, index) => {
                                                return (
                                                    <div
                                                        key={index}
                                                        className="flex flex-row items-center justify-center w-full"
                                                    >
                                                        {
                                                            notification.getNotification()
                                                        }
                                                    </div>
                                                )
                                            })
                                            :
                                            <div className="w-full p-2 flex items-center justify-center">
                                                <lord-icon
                                                    src="/resources/inbox-gray-700.json"
                                                    trigger="loop"
                                                    style={{
                                                        width: '35px',
                                                        height: '35px',
                                                    }}
                                                />
                                                <p className="ml-5 text-sm text-gray-700">
                                                    Nothing to see here.
                                                </p>
                                            </div>
                                    }
                                </div>
                                <div className="w-full h-12 p-2 | flex flex-row items-center justify-between">
                                    <div>
                                    </div>
                                    <div>
                                        <button
                                            onClick={() => {
                                                if (this.context.notifications.notifications.length == 0) {
                                                    return false
                                                }
                                                this.context.markNotificationsRead(
                                                    this.context.notifications.notifications.map(notification => {
                                                        return notification.getId()
                                                    })
                                                );
                                                this.context.pushLocalNotification({
                                                    title: 'Successfully marked as read',
                                                    message: null,
                                                    link: null,
                                                    extra: null,
                                                    type: 'SUCCESS'
                                                });
                                            }}
                                            className="group bg-jt-primary px-2 w-auto items-center text-white rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-red-500 hover:text-white inline-flex h-8  justify-center transition-all duration-150" data-dismiss-target="#toast-success" aria-label="Close"
                                        >
                                            <span className="sr-only">Close</span>
                                            <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clipRule="evenodd"></path></svg>
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </ul>
                </div>
            </div>
        )
    }

}

Sidebar.contextType = Context;

export default Sidebar;