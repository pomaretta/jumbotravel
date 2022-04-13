import { BookingStatus } from "../domain/agent_bookings";

class BookingStatusCollection {

    constructor({
        bookings = [],
    }) {
        this.bookings = bookings;
    }

    static parse(data) {
        return new BookingStatusCollection({
            bookings: data.map(booking => new BookingStatus(booking))
        });
    }

}

export { BookingStatusCollection };