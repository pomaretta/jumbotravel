package domain

type Entity interface {
	GetDestFields() []interface{}
	New()
	Val() interface{}
}
