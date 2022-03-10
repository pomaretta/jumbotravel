package domain

type Data interface {
	GetDestFields() []interface{}
	New()
}
