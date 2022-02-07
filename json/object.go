package json

import (
	ej "encoding/json"
	"reflect"
)

// Object is a general structure to represent JavaScript objects.
type Object map[string]interface{}

// Equal checks if the json.Object has the same values with another one.
func (o Object) Equal(another Object) bool {
	return reflect.DeepEqual(o, another)
}

// String returns JSON text.
func (o Object) String() (s string) {
	s = "{}"

	if marshaled, err := ej.Marshal(o); err == nil {
		s = (string)(marshaled)
	}

	return
}

// FromJSON creates json.Object from JSON.
func FromJSON(s string) (o Object, err error) {
	err = ej.Unmarshal([]byte(s), &o)
	return
}
