package json

import (
	"testing"
)

func TestJSONObject_Equal(t *testing.T) {
	var shouldBeEquivalent bool = JSONObject{
		"object": JSONObject{
			"number": 1.23,
			"string": "string",
		},
	}.Equal(JSONObject{
		"object": JSONObject{
			"number": 1.23,
			"string": "string",
		},
	})

	if !shouldBeEquivalent {
		t.Errorf("JSONObject.Equal should be able distinguish two JSONObject variables that have same values")
	}

	var shouldNotBeEquivalent bool = JSONObject{
		"object": JSONObject{
			"number": 0,
			"string": "hello world",
		},
	}.Equal(JSONObject{
		"object": JSONObject{
			"number": 1.23,
			"string": "string",
		},
	})

	if shouldNotBeEquivalent {
		t.Errorf("JSONObject.Equal should be able distinguish two JSONObject variables that have different values")
	}
}

func TestJSONObject_String(t *testing.T) {
	var (
		object = JSONObject{
			"string": "i-am-string",
			"number": 0,
			"array": []JSONObject{
				{"in-array1": "array1"},
				{"in-array2": "array2"},
			},
			"object": JSONObject{
				"in-object": "object",
			},
		}

		json = object.String()
	)

	if json != `{"array":[{"in-array1":"array1"},{"in-array2":"array2"}],"number":0,"object":{"in-object":"object"},"string":"i-am-string"}` {
		t.Errorf("should return correct JSON")
	}
}

func TestFromJSON(t *testing.T) {
	var (
		object JSONObject
		err    error
	)

	if object, err = FromJSON(`{"foo":"bar"}`); err != nil {
		t.Errorf("should be able to parse JSON")
	}

	if object["foo"] != "bar" {
		t.Errorf("should return corect JSONObject")
	}

	if object, err = FromJSON(``); err == nil {
		t.Errorf("should NOT be able to parse JSON")
	}

	if object != nil {
		t.Errorf("should be a nil JSONObject")
	}
}
