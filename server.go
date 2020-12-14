package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func serve(paths ...string) {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	setRoutes(getRouteConfig(paths), r)

	_ = http.ListenAndServe(":8080", r)
}

func setRoutes(configs []config, r *mux.Router) {
	for _, c := range configs {
		for _, route := range c.Routes {
			r.HandleFunc(route.Url, func(w http.ResponseWriter, r *http.Request) {
				checkAuthorization(w, r, route)
			}).Methods(route.Method)
			log.Printf("added route: %s %s\n", route.Method, route.Url)
		}
	}
}
