package models

type Model interface {
	Equals(a interface{}) bool
}
