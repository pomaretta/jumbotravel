import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../base/sidebar';
// import NavBar from '../base/navbar';
import NavBar from '../base/navbar';

class App extends Component {

    componentDidUpdate() {
        if (this.props.app.hasToLogIn()) {
            this.props.app.logout();
        }
    }

    render() {
        return (
            <div className="w-screen h-screen min-h-screen flex flex-row flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800">
                <Helmet>
                    <title>Home - JumboTravel</title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={1} />
                <div className="w-full h-full">
                    <NavBar app={this.props.app} config={this.props.config} />
                </div>
            </div>
        );
    }

}

export default App;