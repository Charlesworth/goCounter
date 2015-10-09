package concurrentMap

import (
	"testing"
)

var testMap *Map

func BeforeEach() {
	testMap = New()

	testMap.m["zero"] = 0
	testMap.m["one"] = 1
}

func TestConcurrentMapGet(t *testing.T) {
	BeforeEach()
	if testMap.Get("zero") != 0 {
		t.Error("Get on map['zero'] did not return 0")
	}

	if testMap.Get("one") != 1 {
		t.Error("Get on map['one'] did not return 1")
	}
}

func TestConcurrentMapSet(t *testing.T) {
	BeforeEach()
	testMap.Set("test", 66)
	if testMap.Get("test") != 66 {
		t.Error("Set on previosly un-initialized map['test'] of 66 did not return 66 on a get")
	}

	testMap.Set("test", 22)
	if testMap.Get("test") != 22 {
		t.Error("Set on already initialized map['test'] of 22 did not return 22 on a get")
	}
}

func TestConcurrentMapIncrement(t *testing.T) {
	BeforeEach()
	if testMap.Increment("zero") != 1 {
		t.Error("Increment on map['zero']  did not return 1")
	}

	if testMap.Increment("one") != 2 {
		t.Error("Increment on map['one'] did not return 2")
	}
}

func TestGetMap(t *testing.T) {
	BeforeEach()
	exportedMap := testMap.GetMap()
	if exportedMap["zero"] != 0 {
		t.Error("getMap did not return the initialized map")
	}
}
