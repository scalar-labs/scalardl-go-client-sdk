package dl

import "testing"

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
