package dl

import (
	"encoding/json"
	"reflect"

	"github.com/google/uuid"
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
	if err = json.Unmarshal([]byte(s), &o); err != nil {
		o = JSONObject{}
	}

	return
}

// WithNonce checks if the object contains `nonce` properties.
// If not, then add one.
func (j *JSONObject) WithNonce() {
	if _, ok := (*j)["nonce"]; !ok {
		(*j)["nonce"] = uuid.NewString()
	}

	if (*j)["nonce"] == "" {
		(*j)["nonce"] = uuid.NewString()
	}
}
