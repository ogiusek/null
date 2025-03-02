package null_test

import (
	"encoding/json"
	"testing"

	"github.com/ogiusek/null"
)

// go test .

type ExampleStruct struct {
	Val null.Nullable[int] `json:"val"`
}

func TestUnmarshalExpectNull(t *testing.T) {
	var examples []string = []string{
		"{}",
		"{\"val\": null}",
	}

	for _, encoded := range examples {
		var ex ExampleStruct
		if err := json.Unmarshal([]byte(encoded), &ex); err != nil {
			t.Errorf("error decoding \"%s\" %s", encoded, err.Error())
			continue
		}
		if ex.Val.Ok {
			t.Errorf("error decoding \"%s\" expeced !Ok", encoded)
			continue
		}
	}
}

func TestUnmarshalExpectNumber(t *testing.T) {
	var examples map[string]int = map[string]int{
		"{\"val\": 1}": 1,
		"{\"val\": 2}": 2,
		"{\"val\": 3}": 3,
		"{\"val\": 4}": 4,
	}

	for encoded, expected := range examples {
		var ex ExampleStruct
		if err := json.Unmarshal([]byte(encoded), &ex); err != nil {
			t.Errorf("error decoding \"%s\" %s", encoded, err.Error())
			continue
		}
		if !ex.Val.Ok {
			t.Errorf("error decoding \"%s\" expeced Ok", encoded)
			continue
		}
		if ex.Val.Val != expected {
			t.Errorf("error decoding \"%s\"decoded value is %d expected %d", encoded, ex.Val.Val, expected)
			continue
		}

	}
}

func TestUnmarshalExpectedError(t *testing.T) {
	var examples []string = []string{
		"{\"val\": \"\"}",
		"{\"val\": {}}",
		"{\"val\": []}",
	}

	for _, encoded := range examples {
		var ex ExampleStruct
		if err := json.Unmarshal([]byte(encoded), &ex); err == nil {
			t.Errorf("expected error decoding \"%s\"", encoded)
			continue
		}
	}
}
