package esme

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var routeConfigMap map[string]*route

func handleAll(w http.ResponseWriter, req *http.Request) {
	if _, ok := routeConfigMap[getRouteMapKey(req.Method, req.URL.Path)]; !ok {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	r := routeConfigMap[getRouteMapKey(req.Method, req.URL.Path)]

	errStr, statusCode := checkAuthorization(req, r)
	if errStr != "" {
		http.Error(w, errStr, statusCode)
		return
	}

	err := checkBody(r.Body, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if r.Response != nil {
		sendResponse(w, r)
	}
}

func sendResponse(w http.ResponseWriter, r *route) {
	respStr, err := json.Marshal(r.Response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	_, _ = fmt.Fprintf(w, string(respStr))
}
