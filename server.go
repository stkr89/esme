package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var routeConfigMap map[string]*route

func serve(port string, paths ...string) {
	routeConfig, err := getRouteConfig(paths)
	if err != nil {
		log.Println(err.Error())
	}

	setRoutes(routeConfig)
	launchServer(port)
}

func launchServer(port string) {
	m := http.NewServeMux()
	s := http.Server{Addr: ":" + port, Handler: m}

	m.HandleFunc("/", handleAll)
	m.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
		go func() {
			log.Println("shutting down ESME server on port " + port)
			if err := s.Shutdown(context.Background()); err != nil {
				log.Println(err)
			}
		}()
	})

	log.Println("starting ESME server on port " + port)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func handleAll(w http.ResponseWriter, req *http.Request) {
	if _, ok := routeConfigMap[getRouteMapKey(req.Method, req.URL.Path)]; !ok {
		http.Error(w, "Not found", http.StatusNotFound)
	}

	r := routeConfigMap[getRouteMapKey(req.Method, req.URL.Path)]

	errStr, statusCode := checkAuthorization(req, r)
	if errStr != "" {
		http.Error(w, errStr, statusCode)
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
	_, _ = fmt.Fprintf(w, string(respStr))
}

func setRoutes(configs []*config) {
	configMap := make(map[string]*route)

	for _, c := range configs {
		for _, r := range c.Routes {
			configMap[getRouteMapKey(r.Method, r.Url)] = r
		}
	}

	routeConfigMap = configMap
}

func getRouteMapKey(method string, url string) string {
	return method + "::" + url
}
