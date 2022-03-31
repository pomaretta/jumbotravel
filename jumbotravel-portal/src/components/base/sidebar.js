import { Component, Fragment } from "react";
import { Link } from "react-router-dom";

import { Disclosure, Menu, Transition } from '@headlessui/react'

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
            <div className="hidden sm:flex flex-col top-0 left-0 w-16 bg-white h-full border-r">
                <div className="flex items-center justify-center h-14 border-b">
                    {/* Image */}
                    <img className="h-8 w-auto" src="/resources/logo.svg" alt="Workflow" />
                </div>
                <div className="flex-grow">
                    <ul className="flex flex-col py-4 space-y-1 h-full">
                        <li>
                            <Link to="/flights" className={
                                classNames(
                                    this.props.current ? this.props.current == 1 ? "border-jt-primary" : "" : "",
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
                            <Link to="/planes" className={
                                classNames(
                                    this.props.current ? this.props.current == 2 ? "border-jt-primary" : "" : "",
                                    "relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-jt-primary pr-6"
                                )
                            }>
                                <span className="inline-flex justify-center items-center ml-4">
                                    <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                        <path fill="currentColor" d="M20.56 3.91C21.15 4.5 21.15 5.45 20.56 6.03L16.67 9.92L18.79 19.11L17.38 20.53L13.5 13.1L9.6 17L9.96 19.47L8.89 20.53L7.13 17.35L3.94 15.58L5 14.5L7.5 14.87L11.37 11L3.94 7.09L5.36 5.68L14.55 7.8L18.44 3.91C19 3.33 20 3.33 20.56 3.91Z" />
                                    </svg>
                                </span>
                            </Link>
                        </li>
                        <div className="mt-auto-important">
                            <li className="relative">
                                <a onClick={() => {
                                    this.setState({ isOpen: !this.state.isOpen })
                                }} className="border-red-500 relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-jt-primary pr-6">
                                    <span className="inline-flex justify-center items-center ml-4">
                                        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                                                d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9">
                                            </path>
                                        </svg>
                                    </span>
                                </a>
                            </li>
                            <li>
                                <a onClick={() => this.props.app.logout()} className="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-red-500 border-l-4 border-transparent hover:border-jt-primary pr-6">
                                    <span className="inline-flex justify-center items-center ml-4">
                                        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path></svg>
                                    </span>
                                </a>
                            </li>
                        </div>
                        {/* Notification */}
                        <div 
                            className="absolute w-1/4 h-1/4 | bottom-3 left-20 | bg-white | rounded-md | shadow | flex items-center flex-col justify-start | p-2 | overflow-scroll | no-scrollbar"
                            style={{
                                display: this.state.isOpen ? "flex" : "none"
                            }}
                        >
                            <div id="toast-success" className="flex | cursor-pointer hover:shadow-md transition-all duration-200 ease-in-out | items-center w-full p-2 mb-2 text-gray-500 bg-white rounded-lg outline outline-2 outline-green-200 dark:text-gray-400 dark:bg-gray-800" role="alert">
                                <div className="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200">
                                    <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd"></path></svg>
                                </div>
                                <div className="ml-3 text-sm font-normal">Item moved successfully.</div>
                                <button type="button" className="ml-auto -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700" data-dismiss-target="#toast-success" aria-label="Close">
                                    <span className="sr-only">Close</span>
                                    <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clipRule="evenodd"></path></svg>
                                </button>
                            </div>
                            <div className="absolute bottom-0 w-full h-12 p-2 | flex flex-row items-center justify-between">
                                <div>
                                </div>
                                <div>
                                    <button className="group bg-jt-primary px-2 w-auto font-bold items-center text-white hover:text-gray-800 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex h-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700 justify-center transition-all duration-150" data-dismiss-target="#toast-success" aria-label="Close">
                                        <span className="sr-only">Close</span>
                                        <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clipRule="evenodd"></path></svg>
                                        <div className="hidden group-hover:block">
                                            <span className="ml-2">Clear All</span>
                                        </div>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </ul>
                </div>
            </div>
        )
    }

}

export default Sidebar;