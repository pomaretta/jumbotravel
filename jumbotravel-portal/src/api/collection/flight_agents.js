import FlightAgent from "../domain/flight_agent";

class FlightAgentsCollection {

    constructor({
        agents = [],
    }) {
        this.agents = agents;
    }

    getAll() {
        return this.flights;
    }

    static parse(data) {
        return new FlightAgentsCollection({
            agents: data.map(agent => new FlightAgent(agent))
        });
    }

}

export default FlightAgentsCollection;