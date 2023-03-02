package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.Handle("/", handlers())
	fmt.Println("Server is ready listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", serveIndex).Methods("GET")
	router.HandleFunc("/video/", hlsIndexHandler).Methods("GET")
	router.HandleFunc("/video/{segmentName:index[0-9]+.ts}", streamHandler).Methods("GET")
	return router
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	// Serve the plain html file
	http.ServeFile(w, r, "index.html")
}

func hlsIndexHandler(w http.ResponseWriter, r *http.Request) {
	// index.m3u8 location harcoded exclusively for didactic purposes,
	// don't do this in a production environment
	http.ServeFile(w, r, "video/index.m3u8")
	w.Header().Set("Content-Type", "application/x-mpegURL")
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// get the segment name
	vars := mux.Vars(r)
	segName := vars["segmentName"]
	// Generate the segment name according to path
	segmentFile := fmt.Sprintf("video/%s", segName)
	// Serve the appropiate segment
	http.ServeFile(w, r, segmentFile)
	w.Header().Set("Content-Type", "video/MP2T")
}
