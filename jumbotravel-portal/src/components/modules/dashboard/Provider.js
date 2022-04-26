import React, { Component } from "react";
import AppContext from "../../context/app";
import { BarChart, Bar, XAxis, Tooltip, ResponsiveContainer, PieChart, Pie, Cell, Legend, ComposedChart, YAxis, CartesianGrid, Area, Line } from 'recharts';
import WelcomeBanner from "./Banner";
import classNames from "../../utils/classname";

const pieColors = ['rgb(47, 154, 255)', 'rgb(47, 209, 255)', 'rgb(239,68,68)'];

class ProviderDashboard extends Component {

    render() {
        return (
            <div className="relative w-full h-screen | p-4">
                <WelcomeBanner agent={this.context.agent} />
                {/* Controls */}
                <div className="w-full | flex flex-col sm:flex-row justify-between items-center | mb-4 | py-3">
                    <div className={classNames(
                        "flex | justify-start items-center | space-x-4",
                        this.context.agentDashboardPrimaryChart && this.context.agentDashboardSecondaryChart && this.context.agentDashboardCompositeChart ? "invisible" : ""
                    )}>
                        <p className="font-bold | animate-in zoom-in-50">Updating visuals, stay here</p>
                        <progress class="progress progress-info w-56 | animate-in zoom-in-50"></progress>
                    </div>
                    <div className="flex justify-between sm:justify-end items-center | w-full sm:w-1/4">
                        <p className="text-md font-bold mr-4">Dashboard Range</p>
                        <select className="select select-md select-info dark:bg-white w-1/2 | shadow-lg shadow-blue-100" onChange={(ev) => {
                            this.props.changeSelectedDays(ev.target.value)
                        }}>
                            <option selected value={30}>30 Days</option>
                            <option value={1}>1 Day</option>
                            <option value={7}>7 Days</option>
                        </select>
                    </div>
                </div>
                {/* Bookings Chart */}
                <div className="flex items-center justify-between flex-col sm:flex-row | w-full | space-y-4 sm:space-y-0 sm:space-x-4 | mb-8">
                    <div
                        className={classNames(
                            "w-full sm:w-1/2 | h-72 | flex flex-col items-start justify-between | bg-gradient-to-t from-gray-100 to-white | rounded-md | shadow | p-4 | space-y-4",
                            this.context.agentDashboardPrimaryChart && this.context.agentDashboardSecondaryChart && this.context.agentDashboardCompositeChart ? "" : "animate-pulse"
                        )}
                    >
                        <div>
                            <h4 className="text-2xl font-bold">Bookings count</h4>
                            <p className="text-xs text-gray-500">Booking count summarized by day</p>
                        </div>
                        {/* Bookings Chart */}
                        <ResponsiveContainer width="100%" height="100%">
                            <BarChart width={150} height={40}
                                // Change with days of super component
                                data={
                                    this.context.agentDashboardPrimaryChart ? this.context.agentDashboardPrimaryChart.getDays(this.props.selectedDays, "Bookings") : []
                                }
                            >
                                <XAxis dataKey="name" stroke="#000" />
                                <Tooltip wrapperStyle={{ width: 100, backgroundColor: '#ccc' }} />
                                <Bar dataKey="Bookings" fill="#54daff" barSize={30} />
                            </BarChart>
                        </ResponsiveContainer>
                    </div>
                    <div 
                        className={classNames(
                            "w-full sm:w-1/2 | h-72 | flex flex-col items-start justify-between | bg-gradient-to-t from-gray-100 to-white | rounded-md | shadow | p-4 | space-y-4",
                            this.context.agentDashboardPrimaryChart && this.context.agentDashboardSecondaryChart && this.context.agentDashboardCompositeChart ? "" : "animate-pulse"
                        )}
                    >

                        <div>
                            <h4 className="text-2xl font-bold">Bookings summary</h4>
                            <p className="text-xs text-gray-500">Bookings count summarized by status</p>
                        </div>

                        {/* Bookings Chart */}
                        <ResponsiveContainer width="100%" height="100%">
                            <PieChart width={400} height={400}>
                                <Legend align="right" verticalAlign="top" layout="vertical" iconType="circle" height={36} />
                                <Pie
                                    data={
                                        this.context.agentDashboardSecondaryChart ? this.context.agentDashboardSecondaryChart.getValues() : null
                                    }
                                    cx="50%"
                                    cy="50%"
                                    labelLine={false}
                                    outerRadius={80}
                                    fill="#8884d8"
                                    nameKey="name"
                                    dataKey="value"
                                    label
                                >
                                    {
                                        this.context.agentDashboardSecondaryChart ?
                                            this.context.agentDashboardSecondaryChart.getValues().map((entry, index) => (
                                                <Cell key={`cell-${index}`} fill={pieColors[index % pieColors.length]} />
                                            ))
                                            : null
                                    }
                                </Pie>
                            </PieChart>
                        </ResponsiveContainer>
                    </div>
                </div>
                {/* Summary Line Chart */}
                <div 
                    className={classNames(
                        "flex flex-col items-start justify-between | w-full h-80 | space-y-4 | bg-gradient-to-r from-gray-100 to-white | rounded-md | shadow | p-4",
                        this.context.agentDashboardPrimaryChart && this.context.agentDashboardSecondaryChart && this.context.agentDashboardCompositeChart ? "" : "animate-pulse"
                    )}
                >
                    <div>
                        <h4 className="text-2xl font-bold">Composite summary</h4>
                        <p className="text-xs text-gray-500">Get summary about flights, bookings and total</p>
                    </div>
                    <ResponsiveContainer width="100%" height="100%">
                        <ComposedChart width={730} height={250} data={
                            this.context.agentDashboardCompositeChart ?
                                this.context.agentDashboardCompositeChart.getDays(this.props.selectedDays) : null
                        }>
                            <XAxis dataKey="name" />
                            <YAxis />
                            <Tooltip />
                            <Legend />
                            <CartesianGrid stroke="#f5f5f5" />
                            <Area type="monotone" dataKey="Bookings" fill="rgb(84, 218, 255)" stroke="rgb(84, 218, 255)" />
                            <Bar dataKey="Flights" barSize={20} fill="rgb(47, 154, 255)" />
                            <Line type="monotone" dataKey="Total" stroke="rgb(239,68,68)" />
                        </ComposedChart>
                    </ResponsiveContainer>
                </div>
            </div>
        )
    }

}
ProviderDashboard.contextType = AppContext;

export default ProviderDashboard;