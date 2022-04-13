import React from "react";
import AppContext from "../../context/app";
import ClassNames from '../../utils/classname';

class Route extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            isOpen: this.props.index != null && this.props.index === 0 ? true : false
        }

    }

    render() {
        return (
            <div className="flex flex-col | items-start justify-start | w-full">
                <div className="flex flex-row | items-center justify-start | w-full"
                    onClick={() => {
                        this.setState({
                            isOpen: !this.state.isOpen
                        });
                    }}
                >
                    <h3
                        className="text-2xl sm:text-xl | font-semibold | text-brand-blue p-4 | w-3/6 sm:w-1/6"
                    >
                        {this.props.route.route.departure} - {this.props.route.route.arrival}
                    </h3>
                    <div className={ClassNames(
                        this.state.isOpen ? "" : "-rotate-90 | -transform-rotate-90",
                        "ml-2 | text-brand-blue | font-semibold | cursor-pointer | px-1 py-1 | rounded-md | hover:bg-white | hover:text-gray-500"
                    )}>
                        <svg className="w-6 h-6 | fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
                            <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
                        </svg>
                    </div>
                </div>
                <div
                    className={ClassNames(
                        this.state.isOpen ? "flex" : "hidden",
                        "flex-col | items-start justify-start | w-full px-4 | space-y-3"
                    )}
                >
                    {
                        this.props.route.flights.map((flight, index) => {
                            return (
                                <a
                                    href={`/flights/${flight.flight_id}`}
                                    key={index}
                                    className="flex flex-row | justify-between | items-center | bg-white | w-full | p-2 px-4 | rounded-lg | shadow-sm hover:shadow-md"
                                >
                                    {/* Desktop */}
                                    <div className="hidden sm:flex w-full">
                                        <div className="flex sm:flex-row | w-full | items-start sm:items-center justify-start | space-y-1 sm:space-y-0 sm:space-x-3">
                                            <p className="text-sm | text-brand-blue | font-semibold">
                                                {flight.carrier} | {flight.flight_number}
                                            </p>
                                            <p className="text-xs | text-gray-700">
                                                Departure: {flight.departure_time}
                                            </p>
                                            <p className="text-xs | text-gray-700">
                                                Arrival: {flight.arrival_time}
                                            </p>
                                        </div>
                                        {/* Controls */}
                                        <div className="flex flex-row | justify-center items-center | space-x-2 sm:space-x-6">
                                            {/* Warning SVG */}
                                            {/* <svg className="w-6 h-6 | fill-current | text-red-500" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zm12.73-1.41A8 8 0 1 0 4.34 4.34a8 8 0 0 0 11.32 11.32zM9 11V9h2v6H9v-4zm0-6h2v2H9V5z" /></svg> */}
                                            <p
                                                className={ClassNames(
                                                    flight.status === "FLYING" ? "bg-jt-primary3" :
                                                    flight.status === "DEPARTURE" ? "bg-red-400" :
                                                    flight.status === "BUSY" ? "bg-orange-400" :
                                                    "bg-jt-primary",
                                                    "hover:bg-brand-blue-dark text-white text-xs font-bold py-2 px-4 rounded | flex items-center justify-center | w-28 sm:w-32",
                                                )}
                                            >
                                                {flight.status}
                                            </p>
                                        </div>
                                    </div>
                                    {/* Mobile */}
                                    <div className="flex sm:hidden w-full flex-col">
                                        <div className="flex flex-row | w-full | items-center justify-between | mb-5">
                                            <p className="text-md | text-brand-blue | font-semibold">
                                                {flight.carrier} | {flight.flight_number}
                                            </p>
                                            <p
                                                className={ClassNames(
                                                    flight.status === "FLYING" ? "bg-jt-primary3" : "bg-jt-primary",
                                                    "hover:bg-brand-blue-dark text-white text-xs font-bold py-2 px-4 rounded | flex items-center justify-center | w-28 sm:w-32",
                                                )}
                                            >
                                                {flight.status}
                                            </p>

                                        </div>
                                        {/* Controls */}
                                        <div className="flex flex-col | w-full | justify-start items-start | space-y-2">
                                            {/* Warning SVG */}
                                            <p className="text-md | text-gray-700 flex flex-row | items-center justify-between | w-full">
                                                <span className="font-bold">Departure</span>
                                                <span>{flight.departure_time}</span>
                                            </p>
                                            <p className="text-md | text-gray-700 flex flex-row | items-center justify-between | w-full">
                                                <span className="font-bold">Arrival</span>
                                                <span>{flight.arrival_time}</span>
                                            </p>
                                            <p className="text-md | text-gray-700 flex flex-row | items-center justify-between | w-full">
                                                <span className="font-bold">Departure Airport</span>
                                                <span>{flight.departure_commonname}</span>
                                            </p>
                                            <p className="text-md | text-gray-700 flex flex-row | items-center justify-between | w-full">
                                                <span className="font-bold">Arrival Airport</span>
                                                <span>{flight.arrival_commonname}</span>
                                            </p>
                                        </div>
                                    </div>
                                </a>
                            )
                        })
                    }
                </div>
            </div>
        )
    }

}

Route.contextType = AppContext;

export default Route;