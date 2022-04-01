package entity

import "time"

type Token struct {
	Id        *string    `json:"id"`
	Subject   *string    `json:"subject"`
	Token     *string    `json:"token"`
	AgentId   *int       `json:"agent_id"`
	IssuedAt  *time.Time `json:"issued_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	Active    *bool      `json:"active"`
	CreatedAt *time.Time `json:"created_at"`
} // @name Token

func (v *Token) GetDestFields() []interface{} {
	return []interface{}{
		&v.Id,
		&v.Subject,
		&v.Token,
		&v.AgentId,
		&v.IssuedAt,
		&v.ExpiresAt,
		&v.Active,
		&v.CreatedAt,
	}
}

func (v *Token) New() {
	*v = Token{}
}

func (v *Token) Val() interface{} {
	return *v
}
