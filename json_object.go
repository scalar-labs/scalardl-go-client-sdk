package dl

import (
	"encoding/json"
	"reflect"
)

// JsonObject is a general structure to represent JavaScript objects.
type JsonObject map[string]interface{}

// Equal checks if the JsonObject has the same values with another one.
func (j JsonObject) Equal(another JsonObject) bool {
	return reflect.DeepEqual(j, another)
}

// String returns JSON text.
func (j JsonObject) String() (s string) {
	s = "{}"

	if marshaled, err := json.Marshal(j); err == nil {
		s = (string)(marshaled)
	}

	return
}
