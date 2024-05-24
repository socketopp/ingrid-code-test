package server

import (
	"encoding/json"
	"home/osrm"
	"log"
	"net/http"
)

// RoutesHandler handles requests to the GET /routes endpoint.
func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	src := r.URL.Query().Get("src")
	dst := r.URL.Query()["dst"]

	if src == "" || len(dst) == 0 {
		http.Error(w, "Missing src or dst parameters", http.StatusBadRequest)
		return
	}
	resp, err := osrm.GetDurationAndDistances(src, dst)

	if err != nil {
		log.Printf("error getting durations and distances: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode the response to JSON and write it to the response writer
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("error encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
