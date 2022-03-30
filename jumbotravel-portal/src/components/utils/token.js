class JWTToken {

    constructor({
        jti,
        exp,
        iat,
        token,
        agentId
    }) {
        this.jti = jti;
        this.exp = exp;
        this.iat = iat;
        this.token = token;
        this.agentId = agentId;
    }

    isValid() {
        let exp = Date.parse(this.exp);
        return Date.now() < exp;
    }

    getToken() {
        return this.token;
    }

    getJTI() {
        return this.jti;
    }

    getIAT() {
        return Date.parse(this.iat);
    }

    getEXP() {
        return Date.parse(this.exp);
    }

    getAgentId() {
        return this.agentId;
    }

    stringify() {
        return JSON.stringify({
            jti: this.jti,
            exp: this.exp,
            iat: this.iat,
            token: this.token,
            id: this.agentId
        })
    }

}

export default JWTToken;