package entity

import "time"

type Agent struct {
	AgentID   *int       `json:"agent_id"`
	DNI       *string    `json:"dni"`
	Name      *string    `json:"name"`
	Surname   *string    `json:"surname"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
	Type      *string    `json:"type"`
	AirportId *int       `json:"airport_id"`
	CreatedAt *time.Time `json:"created_at"`
	Active    *bool      `json:"active"`
} // @name Agent

func (v *Agent) GetDestFields() []interface{} {
	return []interface{}{
		&v.AgentID,
		&v.DNI,
		&v.Name,
		&v.Surname,
		&v.Email,
		&v.Password,
		&v.Type,
		&v.AirportId,
		&v.CreatedAt,
		&v.Active,
	}
}

func (v *Agent) New() {
	*v = Agent{}
}

func (v *Agent) Val() interface{} {
	return *v
}
