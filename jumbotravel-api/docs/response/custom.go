package response

type JWTToken struct {
	Exp   string `json:"exp"`
	Iat   string `json:"iat"`
	Jti   string `json:"jti"`
	Token string `json:"token"`
} // @name JWTToken
