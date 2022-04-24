class Stat {

    constructor({
        name,
        value
    }) {
        this.name = name;
        this.value = value;
    }

}

class Composite {

    constructor({
        name,
        flights,
        bookings,
        total
    }) {
        this.name = name;
        this.flights = flights;
        this.bookings = bookings;
        this.total = total;
    }

}

export { Stat, Composite };