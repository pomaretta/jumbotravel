import { Component } from 'react';
import APIError from './error';

import JWTToken from '../components/utils/token';
import Agent from './domain/agent_data';
import Flight from './domain/agent_flight';
import { BookingDetails } from './domain/agent_bookings';

// Models
import NotificationCollection from '../api/collection/notification';
import FlightsCollection from './collection/flights';
import FlightAgentsCollection from './collection/flight_agents';
import FlightProductsCollection from './collection/flight_products';
import { BookingStatusCollection, BookingItemCollection } from './collection/bookings';
import { StatCollection, CompositeCollection } from './collection/stats'; 

function requestWithEnvironment({ schema, hostname, path }) {
    return `${schema}://${hostname}${path}`;
}

function getAgentPath({ schema, hostname, token, path }) {
    return `${schema}://${hostname}/agent/${token.getAgentId()}${path}`;
}

function requestWithParameters({ url, params = {} }) {
    return `${url}?${Object.keys(params).map(key => `${key}=${params[key]}`).join('&')}`;
}

class RestClient {

    constructor({
        environment,
        config,
        app
    }) {
        // Set variables
        this.environment = environment;
        this.config = config;
        this.app = app;

        // Request DATA
        this.hostname = this.config.hostname;
        this.schema = this.config.schema;
    }

    async validate({
        token
    }) {
        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: '/validate'
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        }
        )
        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }
        return true;
    }

    async authorize({
        identifier,
        password
    }) {

        // Make request
        const response = await fetch(
            requestWithEnvironment({
                schema: this.schema,
                hostname: this.hostname,
                path: '/auth/login'
            }), {
            method: 'POST',
            body: JSON.stringify({
                dni: identifier,
                password: password
            }),
        });

        if (response.status !== 200) {
            // Try to parse error
            let errorData = await response.json();
            if (!errorData) {
                return false, "error on authorize";
            }
            return false, errorData["error"];
        }

        // Get response
        const data = await response.json();

        // Set state with the token
        this.app.setToken(new JWTToken({
            jti: data.jti,
            exp: data.exp,
            iat: data.iat,
            token: data.token,
            agentId: data.id
        }));

        return true, null;
    }

    // ================
    // NOTIFICATIONS
    // ================

    async getNotifications({ token }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: '/notifications'
                }),
                params: {
                    active: "1",
                    seen: "2",
                    expired: "2",
                }
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        let isNullResponse = data["result"] === null;
        if (isNullResponse) {
            return new NotificationCollection([]);
        }

        // Return data
        return NotificationCollection.parse(data["result"]);
    }

    async markNotificationAsRead({ token, notifications }) {

        // Make request
        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: '/notifications'
            }), {
            method: 'POST',
            body: notifications.map(notification => notification.signature).join(','),
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            },
        })

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

    }

    // ================
    // AGENT DATA
    // ================

    async getAgentData({ token }) {

        // Make request
        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: '/data'
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            },
        }
        )

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Return data
        return new Agent(data);
    }

    async getAgentFlights({ token }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: '/flights'
                })
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        let isNullResponse = data["result"] === null;
        if (isNullResponse) {
            return new FlightsCollection();
        }

        // Return data
        return FlightsCollection.parse(data["result"]);
    }

    async getAgentFlightDetails({ token, flightId }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/flights/${flightId}/details`
                })
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return null;
        }

        // Get flight data (must be unique)
        const flightData = data["result"][0];

        // Return data
        return new Flight(flightData);
    }

    async getAgentFlightOperations({ token, flightId }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/flights/${flightId}/operations`
                })
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return new NotificationCollection([]);
        }

        return NotificationCollection.parse(data["result"]);
    }

    async getAgentFlightAgents({ token, flightId }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/flights/${flightId}/agents`
                })
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return new FlightAgentsCollection([]);
        }

        return FlightAgentsCollection.parse(data["result"]);
    }

    async getAgentFlightProducts({ token, flightId }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/flights/${flightId}/products`
                })
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return new FlightProductsCollection([]);
        }

        return FlightProductsCollection.parse(data["result"]);
    }

    async updateFlightStatus({ token, flightId }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/flights/${flightId}/status`
                })
            }), {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        return true;
    }

    // ================
    // BOOKING
    // ================

    async getAgentBookingStats({ token, days, type, target }) {

        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/bookings/count`
                }),
                params: {
                    "days": days,
                    "type": type,
                    "target": target
                }
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return new StatCollection([]);
        }

        return StatCollection.parse(data["result"]);
    }

    async getAgentBookingComposite({ token, days }) {

        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/bookings/composite`
                }),
                params: {
                    "days": days,
                }
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return new CompositeCollection([]);
        }

        return CompositeCollection.parse(data["result"]);
    }

    async getAgentBookingStatus({ token }) {

        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: `/bookings/status`
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return new BookingStatusCollection([]);
        }

        return BookingStatusCollection.parse(data["result"]);
    }

    async getBookingDetails({ token, bookingReferenceId }) {

        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: `/bookings/${bookingReferenceId}/details`
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Check if data is null
        if (data["result"] === null) {
            return null;
        }

        return new BookingDetails(data["result"]);
    }

    async getBookingOperations({ token, bookingReferenceId }) {

        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: `/bookings/${bookingReferenceId}/operations`
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        let isNullResponse = data["result"] === null;
        if (isNullResponse) {
            return new NotificationCollection([]);
        }

        // Return data
        return NotificationCollection.parse(data["result"]);
    }

    async getBookingItems({ token, bookingReferenceId }) {
        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: `/bookings/${bookingReferenceId}/items`
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        let isNullResponse = data["result"] === null;
        if (isNullResponse) {
            return new BookingItemCollection([]);
        }

        // Return data
        return BookingItemCollection.parse(data["result"]);
    }

    async putBookingOrder({ token, flightId, items }) {

        // Make request
        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: "/bookings"
                }),
                params: {
                    flightId: flightId,
                }
            }), {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            },
            body: JSON.stringify({
                items: items,
            })
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        return true;
    }

    async putBookingRequest({ token, bookingReferenceId }) {

        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: `/bookings/${bookingReferenceId}/request`
            }), {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            throw new APIError(
                "error on post booking request",
                response.status,
                response.statusText
            )
        }

        return true;
    }

    async putInvoice({ token, bookingReferenceId }) {

        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: `/bookings/${bookingReferenceId}/invoice`
            }), {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        return response.blob();
    }

    async getReport({ token, reportDate }) {

        const response = await fetch(
            requestWithParameters({
                url: getAgentPath({
                    schema: this.schema,
                    hostname: this.hostname,
                    token: token,
                    path: `/report`
                }),
                params: {
                    "day": reportDate,
                }
            }), {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        return response.blob();
    }

    async fillBooking({ token, bookingReferenceId }) {

        const response = await fetch(
            getAgentPath({
                schema: this.schema,
                hostname: this.hostname,
                token: token,
                path: `/bookings/${bookingReferenceId}/complete`
            }), {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token.getToken()}`
            }
        });

        console.log(response.statusText);

        if (response.status !== 200) {
            let error = await response.json();
            throw new APIError(
                error && error["error"] ? error["error"] : response.statusText,
                response.status,
                response.statusText
            )
        }

        return true;
    }

}

export default RestClient;