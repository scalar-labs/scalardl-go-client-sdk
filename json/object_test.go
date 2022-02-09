package json

import (
	"testing"
)

func TestObject_Equal(t *testing.T) {
	var shouldBeEquivalent bool = Object{
		"object": Object{
			"number": 1.23,
			"string": "string",
		},
	}.Equal(Object{
		"object": Object{
			"number": 1.23,
			"string": "string",
		},
	})

	if !shouldBeEquivalent {
		t.Errorf("Object.Equal should be able distinguish two Object variables that have same values")
	}

	var shouldNotBeEquivalent bool = Object{
		"object": Object{
			"number": 0,
			"string": "hello world",
		},
	}.Equal(Object{
		"object": Object{
			"number": 1.23,
			"string": "string",
		},
	})

	if shouldNotBeEquivalent {
		t.Errorf("Object.Equal should be able distinguish two Object variables that have different values")
	}
}

func TestObject_String(t *testing.T) {
	var (
		object = Object{
			"string": "i-am-string",
			"number": 0,
			"array": []Object{
				{"in-array1": "array1"},
				{"in-array2": "array2"},
			},
			"object": Object{
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
		object Object
		err    error
	)

	if object, err = FromJSON(`{"foo":"bar"}`); err != nil {
		t.Errorf("should be able to parse JSON")
	}

	if object["foo"] != "bar" {
		t.Errorf("should return corect Object")
	}

	if object, err = FromJSON(``); err == nil {
		t.Errorf("should NOT be able to parse JSON")
	}

	if object != nil {
		t.Errorf("should be a nil Object")
	}
}
