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
                    <title>Planes - JumboTravel</title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={2} />
                <Notifications app={this.props.app} config={this.props.config} />
                <div className="w-full h-full">
                    <NavBar app={this.props.app} config={this.props.config} />
                </div>
            </div>
        );
    }

}

Module.contextType = Context;

export default Module;