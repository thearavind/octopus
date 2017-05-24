package models

type Model interface {
	toMap() map[string]interface{}
}
