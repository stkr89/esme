package esme

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var routeConfigMap map[string]*route

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func handleAll(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)

	if (*req).Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "")
		return
	}

	if _, ok := routeConfigMap[getRouteMapKey(req.Method, req.URL.Path)]; !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	r := routeConfigMap[getRouteMapKey(req.Method, req.URL.Path)]

	if r.Auth != nil {
		errStr, statusCode := checkAuthorization(req, r)
		if errStr != "" {
			http.Error(w, errStr, statusCode)
			return
		}
	}

	err := checkBody(r.Body, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Response != nil {
		sendResponse(w, r)
	}
}

func handleShutdown(port string, s *http.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
		go func() {
			log.Println("shutting down ESME server on port " + port)
			if err := s.Shutdown(context.Background()); err != nil {
				log.Println(err)
			}
		}()
	}
}

func sendResponse(w http.ResponseWriter, r *route) {
	respStr, err := json.Marshal(r.Response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	_, _ = fmt.Fprintf(w, string(respStr))
}
