package esme

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var routeConfigMap map[string]*route

func handleAll(w http.ResponseWriter, req *http.Request) {
	if _, ok := routeConfigMap[getRouteMapKey(req.Method, req.URL.Path)]; !ok {
		http.Error(w, "Not found", http.StatusNotFound)
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

	if hasResponse(r) {
		sendResponse(w, r)
	}
}

func hasResponse(r *route) bool {
	return r.ResponseArr != nil || r.ResponseObj != nil || r.ResponseStr != ""
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
	var respStr []byte
	var err error

	if r.ResponseStr != "" {
		respStr = []byte(r.ResponseStr)
	} else if r.ResponseObj != nil {
		respStr, err = json.Marshal(r.ResponseObj)
	} else {
		respStr, err = json.Marshal(r.ResponseArr)
	}

	log.Println("---", string(respStr))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	_, _ = fmt.Fprintf(w, string(respStr))
}
