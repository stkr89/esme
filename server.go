package esme

import (
	"log"
	"net/http"
)

func Serve(port string, paths ...string) {
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
	m.HandleFunc("/shutdown", handleShutdown(port, &s))

	log.Println("starting ESME server on port " + port)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func setRoutes(configs []*config) {
	configMap := make(map[string]*route)

	for _, c := range configs {
		for _, r := range c.Routes {
			log.Printf("added route %s %s", r.Method, r.Url)
			configMap[getRouteMapKey(r.Method, r.Url)] = r
		}
	}

	routeConfigMap = configMap
}

func getRouteMapKey(method string, url string) string {
	return method + "::" + url
}
