import { Component } from 'react';
import APIError from './error';

import JWTToken from '../components/utils/token';
import Agent from './domain/agent_data';

// Models
import NotificationCollection from '../api/collection/notification';

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
            throw new APIError(
                "error on validation",
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
            throw new APIError(
                "error on authorization",
                response.status,
                response.statusText
            )
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

        return true;
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
            throw new APIError(
                "error on getNotifications",
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
            throw new APIError(
                "error on markNotificationAsRead",
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
            throw new APIError(
                "error on getAgentData",
                response.status,
                response.statusText
            )
        }

        // Get response
        const data = await response.json();

        // Return data
        return new Agent(data);
    }

}

export default RestClient;