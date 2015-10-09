package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Charlesworth/goCounter/concurrentMap"
	"github.com/julienschmidt/httprouter"
)

func TestMain(t *testing.T) {

}

type getCountHandlerTestStruct struct {
	pageName           string
	url                string
	expectedReturnCode int
	expectedMapValue   int
}

func TestGetCountHandler(t *testing.T) {
	pageViewMap = concurrentMap.New()

	router := httprouter.New()
	router.GET("/:pageID/count.js", getCountHandler)

	tableTestStruct := [3]getCountHandlerTestStruct{
		//test case 0: add "test1" to the counter
		{"page1", "/page1/count.js", 200, 1},
		//test case 1: add a second "test1" to the counter
		{"page1", "/page1/count.js", 200, 2},
		//test case 2: add "test2" to the counter
		{"page2", "/page2/count.js", 200, 1},
	}

	for i := range tableTestStruct {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", tableTestStruct[i].url, nil)

		router.ServeHTTP(w, req)
		if w.Code != tableTestStruct[i].expectedReturnCode {
			t.Error("getCountHandler test case", i, "returned", w.Code, "instead of", tableTestStruct[i].expectedReturnCode)
		}

		mapValue := pageViewMap.Get(tableTestStruct[i].pageName)
		if mapValue != tableTestStruct[i].expectedMapValue {
			t.Error("getCountHandler test case", i, " expected", tableTestStruct[i].expectedMapValue,
				"in pageViewMap instead of", mapValue)
		}
	}
}
