import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';

class Module extends Component {

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
                    <div className="flex flex-col flex-1 h-full p-2">
                        Hello World!
                    </div>
                </div>
            </div>
        );
    }

}

Module.contextType = Context;

export default Module;