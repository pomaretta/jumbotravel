import { BookingStatus, BookingItem } from "../domain/agent_bookings";

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

class BookingItemCollection {

    constructor({
        items = [],
    }) {
        this.items = items;
    }

    static parse(data) {
        return new BookingItemCollection({
            items: data.map(item => new BookingItem(item))
        });
    }

}

export { BookingStatusCollection, BookingItemCollection };