package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Charlesworth/goCounter/concurrentMap"
	"github.com/julienschmidt/httprouter"
)

var pageViewMap *concurrentMap.Map

func main() {
	pageViewMap = concurrentMap.New()

	//set the HTTP routing for the server
	router := httprouter.New()
	router.GET("/:pageID/count.js", getCountHandler)
	//router.PUT("/:pageID", setCountHandler)
	//router.GET("/stats", statsHandler)
	http.Handle("/", router)

	//start the server and listen for requests
	log.Println("goCounter Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getCountHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageName := params.ByName("pageID")
	log.Println(r.RemoteAddr + " requests " + pageName)

	pageViews := pageViewMap.Increment(pageName)

	fmt.Fprintf(w, "document.getElementById('viewCount').innerHTML = '%v Page Views';", pageViews)
}
