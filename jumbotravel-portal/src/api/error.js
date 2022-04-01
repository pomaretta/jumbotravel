class APIError extends Error {

    constructor(message, status, statusText) {
        super(message);
        this.status = status;
        this.statusText = statusText;
    }

    getStatus() {
        return this.status;
    }

    getStatusText() {
        return this.statusText;
    }

    toString() {
        return `${this.status} ${this.statusText}`;
    }
    
}

export default APIError;