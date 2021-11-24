package dl

import (
	"testing"
)

func Test_EqualWithSameJsonObject_ShouldBeTrue(t *testing.T) {
	var shouldBeEquivalent bool = JsonObject{
		"object": JsonObject{
			"number": 1.23,
			"string": "string",
		},
	}.Equal(JsonObject{
		"object": JsonObject{
			"number": 1.23,
			"string": "string",
		},
	})

	if !shouldBeEquivalent {
		t.Errorf("JsonObject.Equal should be able distinguish two JsonObject variables that have same values")
	}

}

func Test_EqualWithDifferentJsonObject_ShouldBeFalse(t *testing.T) {
	var shouldNotBeEquivalent bool = JsonObject{
		"object": JsonObject{
			"number": 0,
			"string": "hello world",
		},
	}.Equal(JsonObject{
		"object": JsonObject{
			"number": 1.23,
			"string": "string",
		},
	})

	if shouldNotBeEquivalent {
		t.Errorf("JsonObject.Equal should be able distinguish two JsonObject variables that have different values")
	}
}

func Test_String_ShouldReturnCorrectJSON(t *testing.T) {
	var (
		object = JsonObject{
			"string": "i-am-string",
			"number": 0,
			"array": []JsonObject{
				{"in-array1": "array1"},
				{"in-array2": "array2"},
			},
			"object": JsonObject{
				"in-object": "object",
			},
		}

		json = object.String()
	)

	if json != `{"array":[{"in-array1":"array1"},{"in-array2":"array2"}],"number":0,"object":{"in-object":"object"},"string":"i-am-string"}` {
		t.Errorf("should return correct JSON")
	}
}
