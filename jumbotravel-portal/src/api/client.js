import { Component } from 'react';

import JWTToken from '../components/utils/token';

function requestWithEnvironment({schema, hostname, path}) {
    return `${schema}://${hostname}${path}`;
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

    validate({
        token
    }) {
        return;
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
            }
        )

        if (response.status !== 200) {
            throw new Error(`${response.status} ${response.statusText}`);
        }

        // Get response
        const data = await response.json();
    
        // Set state with the token
        this.app.setToken(new JWTToken({
            jti: data.jti,
            exp: data.exp,
            iat: data.iat,
            token: data.token
        }));

        return true;
    }

}

export default RestClient;