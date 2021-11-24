package dl

import "reflect"

// JsonObject is a general structure to represent JavaScript objects.
type JsonObject map[string]interface{}

// Equal checks if the JsonObject has the same values with another one.
func (j JsonObject) Equal(another JsonObject) bool {
	return reflect.DeepEqual(j, another)
}
