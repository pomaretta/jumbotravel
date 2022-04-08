class FlightAgent {

    constructor({
        agent_id,
        name,
        surname,
        email
    }) {
        this.agent_id = agent_id;
        this.name = name;
        this.surname = surname;
        this.email = email;
    }

    getAgentId() {
        return this.agent_id;
    }

    getName() {
        return this.name;
    }

    getSurname() {
        return this.surname;
    }

    getFullName() {
        return `${this.name} ${this.surname}`;
    }

    getEmail() {
        return this.email;
    }

}

export default FlightAgent;