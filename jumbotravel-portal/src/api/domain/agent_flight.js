class Flight {

    constructor({
        route_id,
        departure_country,
        departure_city,
        arrival_country,
        arrival_city,
        departure_airport,
        arrival_airport,
        departure_commonname,
        arrival_commonname,
        airplane_id,
        carrier,
        flight_number,
        seats,
        flight_id,
        status,
        departure_time,
        arrival_time,
        flight_created,
        flight_updated,
        updated_at,
        created_at,
        has_booking,
    }) {
        this.route_id = route_id;
        this.departure_country = departure_country;
        this.departure_city = departure_city;
        this.arrival_country = arrival_country;
        this.arrival_city = arrival_city;
        this.departure_airport = departure_airport;
        this.arrival_airport = arrival_airport;
        this.departure_commonname = departure_commonname;
        this.arrival_commonname = arrival_commonname;
        this.airplane_id = airplane_id;
        this.carrier = carrier;
        this.flight_number = flight_number;
        this.seats = seats;
        this.flight_id = flight_id;
        this.status = status;
        this.departure_time = departure_time;
        this.arrival_time = arrival_time;
        this.flight_created = flight_created;
        this.flight_updated = flight_updated;
        this.updated_at = updated_at;
        this.created_at = created_at;
        this.has_booking = has_booking;
    }

    getRoute() {
        return {
            "departure": this.departure_airport,
            "arrival": this.arrival_airport,
        }
    }

}

export default Flight;