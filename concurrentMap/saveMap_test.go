package concurrentMap

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMapToJson(t *testing.T) {
	testMap := make(map[string]int)
	testMap["element1"] = 123

	testJSON, _ := mapToJson(testMap)
	if string(testJSON) != "{\"element1\":123}" {
		t.Error(string(testJSON))
	}
}

func TestFileExists(t *testing.T) {
	if fileExists("nonExistantFile") {
		t.Error("fileExists() returned true when no file didn't exist")
	}

	ioutil.WriteFile("testFile", []byte{}, 644)
	defer os.Remove("testFile")
	if !fileExists("testFile") {
		t.Error("fileExists() returned false when file existed")
	}
}

func TestSaveByteToFile(t *testing.T) {
	err := saveByteToFile([]byte{}, "testFile")
	defer os.Remove("testFile")
	if err != nil {
		t.Error(err)
	}

	if !fileExists("testFile") {
		t.Error("saveByteToFile('testFile') did not produce a file 'testFile'")
	}
}

func TestJsonToMap(t *testing.T) {
	invalidMapJSON := []byte(`{asdfasfasdfs}`)
	nilMap, err := jsonToMap(invalidMapJSON)
	if err == nil {
		t.Error("jsonToMap with invalid json input did not produce an error")
	}

	if nilMap != nil {
		t.Error("jsonToMap with invalid json input did not produce an error")
	}

	validMapJSON := []byte(`{"page1":1,"page2":2}`)
	validMap, err := jsonToMap(validMapJSON)
	if err != nil {
		t.Error("jsonToMap with valid json input produced an error:", err)
	}

	if (validMap["page1"] != 1) && (validMap["page2"] != 2) {
		t.Error("jsonToMap with valid json did not return correctly initialised map with values")
	}
}
