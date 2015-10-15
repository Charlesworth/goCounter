package concurrentMap

import (
	"testing"
)

func TestMapToJson(t *testing.T) {
	testMap := make(map[string]int)
	testMap["element1"] = 123

	testJson, _ := mapToJson(testMap)
	if string(testJson) != "{\"element1\":123}" {
		t.Error(string(testJson))
	}
}
