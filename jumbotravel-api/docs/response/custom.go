package response

type JWTToken struct {
	Exp   string `json:"exp"`
	Iat   string `json:"iat"`
	Jti   string `json:"jti"`
	Token string `json:"token"`
} // @name JWTToken

type AgentData struct {
	AgentId    *int    `json:"agent_id"`
	Dni        *string `json:"dni"`
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Email      *string `json:"email"`
	Type       *string `json:"type"`
	AirportId  *int    `json:"airport_id"`
	Country    *string `json:"country"`
	City       *string `json:"city"`
	Airport    *string `json:"airport"`
	CommonName *string `json:"common_name"`
} // @name AgentData
