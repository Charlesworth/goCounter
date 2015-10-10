package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Charlesworth/goCounter/concurrentMap"
	"github.com/julienschmidt/httprouter"
)

var pageViewMap *concurrentMap.Map

func main() {
	pageViewMap = concurrentMap.New()

	//set the HTTP routing for the server
	router := httprouter.New()
	router.GET("/:pageID/count.js", getCountHandler)
	router.PUT("/:pageID", setCountHandler)
	router.GET("/", getStatsHandler)
	http.Handle("/", router)

	//start the server and listen for requests
	log.Println("goCounter Listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getStatsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println(r.RemoteAddr + " requests page count statistics")
	viewCountMap := pageViewMap.GetMap()

	for page, views := range viewCountMap {
		fmt.Fprintln(w, "Page: [", page, "] Views:", views)
	}
}

func getCountHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageName := params.ByName("pageID")
	log.Println(r.RemoteAddr + " requests " + pageName)

	pageViews := pageViewMap.Increment(pageName)

	fmt.Fprintf(w, "document.getElementById('viewCount').innerHTML = '%v Page Views';", pageViews)
}

func setCountHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageName := params.ByName("pageID")
	countString := r.FormValue("count")
	count, err := strconv.Atoi(countString)

	if err != nil {
		log.Println("Error: ", r.RemoteAddr, "put for page", pageName, "returned with error:", err)
		w.WriteHeader(400)
		fmt.Fprintf(w, "Unable to parse PUT form value 'Count'")
		return
	}

	log.Println(r.RemoteAddr, "puts count", count, "for page", pageName)
	pageViewMap.Set(pageName, count)
}
