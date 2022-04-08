import FlightProduct from "../domain/flight_product";

class FlightProductsCollection {

    constructor({
        products = [],
    }) {
        this.products = products;
    }

    getAll() {
        return this.products;
    }

    static parse(data) {
        return new FlightProductsCollection({
            products: data.map(product => new FlightProduct(product))
        });
    }

}

export default FlightProductsCollection;