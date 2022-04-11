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

    getCurrent() {
        // Get the now time in local time
        let now = new Date();
        // Get the current time in UTC
        let utcNow = new Date(now.getTime() + now.getTimezoneOffset() * 60000);

        // Iterate over all flights and find the one that is currently active
        for (let flight of this.flights) {
            let departureTime = flight.departure_time;
            let arrivalTime = flight.arrival_time;

            // Check if the flight is currently active
            if (departureTime <= utcNow && utcNow <= arrivalTime) {
                // Return the flight
                console.log(`Flight ${flight.flight_id} is currently active`);
                return flight;
            }
        }

        // If no flight is currently active, return the first flight if it exists
        if (this.flights.length > 0) {
            return this.flights[0];
        } else {
            return null;
        }
    }

    static parse(data) {
        return new FlightsCollection({
            flights: data.map(flight => new Flight(flight))
        });
    }

}

export default FlightsCollection;