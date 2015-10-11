package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Charlesworth/goCounter/concurrentMap"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	pageViewMap *concurrentMap.Map
}

func main() {
	server := Server{pageViewMap: concurrentMap.New()}
	http.Handle("/", newRouter(&server))

	//start the server and listen for requests
	log.Println("goCounter Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func newRouter(server *Server) *httprouter.Router {
	router := httprouter.New()
	router.GET("/:pageID/count.js", server.getCountHandler)
	router.PUT("/:pageID", server.setCountHandler)
	router.GET("/", server.getStatsHandler)
	return router
}

func (server *Server) setCountHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageName := params.ByName("pageID")
	countString := r.FormValue("count")
	count, err := strconv.Atoi(countString)

	if err != nil {
		log.Println("Error: ", r.RemoteAddr, "put for page", pageName, "returned with error:", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to parse PUT form value 'Count'")
		return
	}

	log.Println(r.RemoteAddr, "puts count", count, "for page", pageName)
	server.pageViewMap.Set(pageName, count)
}

func (server *Server) getStatsHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println(r.RemoteAddr + " requests page count statistics")
	viewCountMap := server.pageViewMap.GetMap()

	for page, views := range viewCountMap {
		fmt.Fprintln(w, "Page: [", page, "] Views:", views)
	}
}

func (server *Server) getCountHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageName := params.ByName("pageID")
	log.Println(r.RemoteAddr + " requests " + pageName)

	pageViews := server.pageViewMap.Increment(pageName)

	fmt.Fprintf(w, "document.getElementById('viewCount').innerHTML = '%v Page Views';", pageViews)
}
