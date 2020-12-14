package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var routeConfigMap map[string]*route

func serve(paths ...string) {
	routeConfig, err := getRouteConfig(paths)
	if err != nil {
		log.Println(err.Error())
		return
	}

	setRoutes(routeConfig)

	http.HandleFunc("/", handleAll)
	_ = http.ListenAndServe(":8080", nil)
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
