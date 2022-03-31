class Agent {

    constructor({
        agent_id,
        dni,
        name,
        surname,
        email,
        type,
        airport_id,
        country,
        city,
        airport,
        common_name
    }) {
        this.agent_id = agent_id;
        this.dni = dni;
        this.name = name;
        this.surname = surname;
        this.email = email;
        this.type = type;
        this.airport_id = airport_id;
        this.country = country;
        this.city = city;
        this.airport = airport;
        this.common_name = common_name;
    }

    getAgentId() {
        return this.agent_id;
    }

    getDni() {
        return this.dni;
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

    getType() {
        return this.type;
    }

    getAirportId() {
        return this.airport_id;
    }

    getCountry() {
        return this.country;
    }

    getCity() {
        return this.city;
    }

    getAirport() {
        return this.airport;
    }

    getCommonName() {
        return this.common_name;
    }

    getAirportFullName() {
        return `${this.common_name} (${this.country}, ${this.city})`;
    }

}

export default Agent;