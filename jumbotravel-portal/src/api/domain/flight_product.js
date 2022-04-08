class FlightProduct {

    constructor({
        product_id,
        product_code,
        name,
        description,
        brand,
        type,
        stock,
        max,
        saleprice
    }) {
        this.product_id = product_id;
        this.product_code = product_code;
        this.name = name;
        this.description = description;
        this.brand = brand;
        this.type = type;
        this.stock = stock;
        this.max = max;
        this.saleprice = saleprice;
    }

    getProductId() {
        return this.product_id;
    }

    getProductCode() {
        return this.product_code;
    }

    getName() {
        return this.name;
    }

    getDescription() {
        return this.description;
    }

    getBrand() {
        return this.brand;
    }

    getType() {
        return this.type;
    }

    getStock() {
        return this.stock;
    }

    getMax() {
        return this.max;
    }

    getSalePrice() {
        return this.saleprice;
    }

}

export default FlightProduct;