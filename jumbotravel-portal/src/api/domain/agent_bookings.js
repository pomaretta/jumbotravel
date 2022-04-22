class BookingStatus {

    constructor({
        bookingreferenceid,
        status,
        flight_id,
        agent_id,
        agent_name,
        agent_surname,
        provider_id,
        provider_name,
        provider_surname,
        items,
        total,
        created_at,
    }) {
        this.bookingreferenceid = bookingreferenceid;
        this.status = status;
        this.flight_id = flight_id;
        this.agent_id = agent_id;
        this.agent_name = agent_name;
        this.agent_surname = agent_surname;
        this.provider_id = provider_id;
        this.provider_name = provider_name;
        this.provider_surname = provider_surname;
        this.items = items;
        this.total = total;
        this.created_at = created_at;
    }

}

class BookingDetails {

    constructor({
        bookingreferenceid,
		status,
		flight_id,
		agent_id,
		agent_name,
		agent_surname,
		provider_id,
		provider_name,
		provider_surname,
		items,
		total,
		created_at,
        has_invoice,
    }) {
        this.bookingreferenceid = bookingreferenceid;
        this.status = status;
        this.flight_id = flight_id;
        this.agent_id = agent_id;
        this.agent_name = agent_name;
        this.agent_surname = agent_surname;
        this.provider_id = provider_id;
        this.provider_name = provider_name;
        this.provider_surname = provider_surname;
        this.items = items;
        this.total = total;
        this.created_at = created_at;
        this.has_invoice = has_invoice;
    }

}

class BookingItem {

    constructor({
        bookingreferenceid,
        productcode,
        status,
        items,
        price,
        name,
        brand,
        saleprice,
        created_at,
        updated_at,
    }) {
        this.bookingreferenceid = bookingreferenceid;
        this.productcode = productcode;
        this.status = status;
        this.items = items;
        this.price = price;
        this.name = name;
        this.brand = brand;
        this.saleprice = saleprice;
        this.created_at = created_at;
        this.updated_at = updated_at;
    }

}

export { BookingStatus, BookingDetails, BookingItem };