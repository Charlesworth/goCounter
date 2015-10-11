package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Charlesworth/goCounter/concurrentMap"
)

type handlerTest struct {
	TestName       string
	Method         string
	Path           string
	QueryString    string
	ExpectedStatus int
	ExpectedBody   string
}

var handlerTests = []handlerTest{
	{TestName: "test setCountHandler with valid query string on new page",
		Method:         "PUT",
		Path:           "/newPage",
		QueryString:    "?count=1",
		ExpectedStatus: http.StatusOK},
	{TestName: "test setCountHandler with valid query string on existing page",
		Method:         "PUT",
		Path:           "/pageInitialisedWithCount1",
		QueryString:    "?count=1",
		ExpectedStatus: http.StatusOK},
	{TestName: "test setCountHandler with invalid query string",
		Method:         "PUT",
		Path:           "/newPage",
		QueryString:    "?count=one",
		ExpectedStatus: 400,
		ExpectedBody:   "Unable to parse PUT form value 'Count'"},
	{TestName: "test setCountHandler with empty query string",
		Method:         "PUT",
		Path:           "/newPage",
		QueryString:    "?count=",
		ExpectedStatus: 400,
		ExpectedBody:   "Unable to parse PUT form value 'Count'"},
	{TestName: "test setCountHandler with no query string",
		Method:         "PUT",
		Path:           "/newPage",
		ExpectedStatus: 400,
		ExpectedBody:   "Unable to parse PUT form value 'Count'"},
	{TestName: "test getCountHandler on new page",
		Method:         "GET",
		Path:           "/newPage/count.js",
		ExpectedStatus: 200,
		ExpectedBody:   "document.getElementById('viewCount').innerHTML = '1 Page Views';"},
	{TestName: "test getCountHandler on existing page",
		Method:         "GET",
		Path:           "/pageInitialisedWithCount1/count.js",
		ExpectedStatus: 200,
		ExpectedBody:   "document.getElementById('viewCount').innerHTML = '2 Page Views';"},
	{TestName: "test getStatsHandler",
		Method:         "GET",
		Path:           "/",
		ExpectedStatus: 200,
		ExpectedBody:   "Page: [ pageInitialisedWithCount1 ] Views: 1\n"},
}

func TestAllHandlers(t *testing.T) {
	testServer := Server{}

	testRouter := newRouter(&testServer)

	for i := range handlerTests {
		testMap := concurrentMap.New()
		testMap.Increment("pageInitialisedWithCount1")
		testServer.pageViewMap = testMap

		//make the test request and response and then serve the handler.
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(handlerTests[i].Method, handlerTests[i].Path+handlerTests[i].QueryString, nil)
		testRouter.ServeHTTP(w, req)

		if w.Code != handlerTests[i].ExpectedStatus {
			t.Error(handlerTests[i].TestName, "returned status code", w.Code, "instead of", handlerTests[i].ExpectedStatus)
		}

		if w.Body.String() != handlerTests[i].ExpectedBody {
			t.Error(handlerTests[i].TestName, "returned body:\n{", w.Body.String(), "}\ninstead of:\n{", handlerTests[i].ExpectedBody, "}")
		}
	}
}
