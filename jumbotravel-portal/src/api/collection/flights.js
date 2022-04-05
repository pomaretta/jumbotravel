import Flight from "../domain/agent_flight";

class FlightsCollection {

    constructor({
        flights = [],
    }) {
        this.flights = flights;
    }

    getAll() {
        return this.flights;
    }

    getRoutes() {

        let routes = this.flights.map(flight => flight.getRoute());
        if (routes.length === 0) {
            return [];
        }

        let uniqueRoutes = routes.filter((route, index, self) => {
            return index === self.findIndex(r => r.departure === route.departure && r.arrival === route.arrival);
        });

        return uniqueRoutes;
    }

    getFlightsByRoute(route) {

        let flights = this.flights.filter(flight => {
            return flight.getRoute().departure === route.departure && flight.getRoute().arrival === route.arrival;
        });

        return flights;
    }

    getFlightsByRoutes() {
        let routes = this.getRoutes();
        let flightsByRoutes = routes.map(route => {
            return {
                route: route,
                flights: this.getFlightsByRoute(route)
            }
        });
        return flightsByRoutes;
    }

    static parse(data) {
        return new FlightsCollection({
            flights: data.map(flight => new Flight(flight))
        });
    }

}

export default FlightsCollection;