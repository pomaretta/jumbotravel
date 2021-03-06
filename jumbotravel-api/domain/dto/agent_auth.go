package dto

type AgentAuth struct {
	AgentId   *int    `json:"agent_id"`
	DNI       *string `json:"dni"`
	Password  *string `json:"password"`
	AgentType *string `json:"agent_type"`
} // @name AgentAuth

func (v *AgentAuth) GetDestFields() []interface{} {
	return []interface{}{
		&v.AgentId,
		&v.DNI,
		&v.Password,
		&v.AgentType,
	}
}

func (v *AgentAuth) New() {
	*v = AgentAuth{}
}

func (v *AgentAuth) Val() interface{} {
	return *v
}
