import React, { Component } from "react";
import { Helmet } from "react-helmet";

import Sidebar from '../../base/sidebar';
import NavBar from '../../base/navbar';
import Notifications from "../../base/notifications";
import Context from '../../context/app';

import AssistantDashboard from './Assistant';
import ProviderDashboard from './Provider';

class Content extends Component {
    render() {
        return (
            <div className="relative | flex flex-col | w-full h-full | sm:overflow-scroll | sm:no-scrollbar | bg-gray-50 | mb-2">
                {
                    this.context.agent && this.context.agent.type === 'ASSISTANT' ?
                    <AssistantDashboard update={this.props.update} selectedDays={this.props.selectedDays} changeSelectedDays={this.props.changeSelectedDays} />
                    :
                    <ProviderDashboard update={this.props.update} selectedDays={this.props.selectedDays} changeSelectedDays={this.props.changeSelectedDays} />
                }
            </div>
        )
    }
}
Content.contextType = Context;

class Module extends Component {

    constructor(props) {
        super(props);

        this.state = {
            selectedDays: 30,
        }

        this.updateDashboard = this.updateDashboard.bind(this);
        this.changeSelectedDays = this.changeSelectedDays.bind(this);
    }

    componentDidMount() {
        this.context.getAgentDashboardPrimaryChart(this.state.selectedDays);
        this.context.getAgentDashboardSecondaryChart(this.state.selectedDays);
        this.context.getAgentDashboardCompositeChart(this.state.selectedDays);
    }

    componentDidUpdate() {
        if (this.context.hasToLogIn()) {
            this.context.logout();
        }
    }

    async changeSelectedDays(days) {
        await this.setState({
            selectedDays: days
        });
        this.context.removeDashboardCharts();
        await this.updateDashboard();
    }

    async updateDashboard() {
        setTimeout(() => {
            this.context.getAgentDashboardPrimaryChart(this.state.selectedDays);
            this.context.getAgentDashboardSecondaryChart(this.state.selectedDays);
            this.context.getAgentDashboardCompositeChart(this.state.selectedDays);
        }, 3500);
    }

    render() {
        return (
            <div className="w-screen h-screen flex flex-row flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800">
                <Helmet>
                    <title>Dashboard - JumboTravel</title>
                </Helmet>
                <Sidebar app={this.props.app} config={this.props.config} current={3} />
                <Notifications app={this.props.app} config={this.props.config} />
                <div className="relative w-full h-full | flex flex-col justify-start items-start">
                    {/* NavBar */}
                    <NavBar app={this.props.app} config={this.props.config} />
                    {/* Content */}
                    <Content update={this.updateDashboard} selectedDays={this.state.selectedDays} changeSelectedDays={this.changeSelectedDays} /> 
                </div>
            </div>
        );
    }

}
Module.contextType = Context;

export default Module;