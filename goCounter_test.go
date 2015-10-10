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

type setCountHandlerTestStruct struct {
	pageName           string
	url                string
	expectedReturnCode int
	expectedMapValue   int
}

func TestSetCountHandler(t *testing.T) {
	pageViewMap = concurrentMap.New()

	router := httprouter.New()
	router.PUT("/:pageID", setCountHandler)

	tableTestStruct := [3]setCountHandlerTestStruct{
		//test case 0: set "test1" to 1 count
		{"page1", "/page1?count=1", 200, 1},
		//test case 1: set "test1" to 0 count
		{"page1", "/page1?count=0", 200, 0},
		//test case 2: set "test1" to 10000 count
		{"page2", "/page2?count=10000", 200, 10000},
	}

	for i := range tableTestStruct {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("PUT", tableTestStruct[i].url, nil)

		router.ServeHTTP(w, req)
		if w.Code != tableTestStruct[i].expectedReturnCode {
			t.Error("setCountHandler test case", i, "returned", w.Code, "instead of", tableTestStruct[i].expectedReturnCode)
		}

		mapValue := pageViewMap.Get(tableTestStruct[i].pageName)
		if mapValue != tableTestStruct[i].expectedMapValue {
			t.Error("setCountHandler test case", i, " expected", tableTestStruct[i].expectedMapValue,
				"in pageViewMap instead of", mapValue)
		}
	}
}

func TestSetCountHandlerInvalidQueryString(t *testing.T) {
	pageViewMap = concurrentMap.New()

	router := httprouter.New()
	router.PUT("/:pageID", setCountHandler)

	tableTestStruct := [2]setCountHandlerTestStruct{
		//test case 0: set "test1" with text instead of a number in the count query string
		{"page1", "/page1?count=ten", 400, 0},
		//test case 1: set "test1" with empty query string
		{"page1", "/page1?count=", 400, 0},
	}

	for i := range tableTestStruct {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest("PUT", tableTestStruct[i].url, nil)

		router.ServeHTTP(w, req)
		if w.Code != tableTestStruct[i].expectedReturnCode {
			t.Error("setCountHandler test case", i, "returned", w.Code, "instead of", tableTestStruct[i].expectedReturnCode)
		}

		mapValue := pageViewMap.Get(tableTestStruct[i].pageName)
		if mapValue != tableTestStruct[i].expectedMapValue {
			t.Error("setCountHandler test case", i, " expected", tableTestStruct[i].expectedMapValue,
				"in pageViewMap instead of", mapValue)
		}
	}
}

func TestGetStatsHandler(t *testing.T) {
	pageViewMap = concurrentMap.New()
	pageViewMap.Set("test", 8)

	router := httprouter.New()
	router.GET("/", getStatsHandler)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)

	router.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Error("getStatsHandler did not return 200 on a valid request")
	}
	expectedStringOutput := "Page: [ test ] Views: 8\n"
	if w.Body.String() != expectedStringOutput {
		t.Error("getStatsHandler returned:\n", w.Body.String(), "and expected:\n", expectedStringOutput)
	}
}
