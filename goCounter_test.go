package main

import (
	"testing"
)

var testMap *concurrentMap

func BeforeEach() {
	testMap = new()

	testMap.m["zero"] = 0
	testMap.m["one"] = 1
	testMap.m["many"] = 9
}

func TestConcurrentMapGet(t *testing.T) {
	BeforeEach()
	if testMap.get("zero") != 0 {
		t.Error("Get on map['zero'] did not return 0")
	}

	if testMap.get("one") != 1 {
		t.Error("Get on map['one'] did not return 1")
	}
}

func TestConcurrentMapSet(t *testing.T) {
	BeforeEach()
	testMap.set("test", 66)
	if testMap.get("test") != 66 {
		t.Error("Set on map['test'] of 66 did not return 66 on a get")
	}

}
