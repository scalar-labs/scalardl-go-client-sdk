package dl

import (
	"encoding/json"
	"reflect"
)

// JSONObject is a general structure to represent JavaScript objects.
type JSONObject map[string]interface{}

// Equal checks if the JSONObject has the same values with another one.
func (j JSONObject) Equal(another JSONObject) bool {
	return reflect.DeepEqual(j, another)
}

// String returns JSON text.
func (j JSONObject) String() (s string) {
	s = "{}"

	if marshaled, err := json.Marshal(j); err == nil {
		s = (string)(marshaled)
	}

	return
}

// FromJSON creates JSONObject from JSON.
func FromJSON(s string) (o JSONObject, err error) {
	err = json.Unmarshal([]byte(s), &o)
	return
}
