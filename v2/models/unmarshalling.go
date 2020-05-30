package models

type Unmarshaller interface {
	Unmarshall(s string) error
}
