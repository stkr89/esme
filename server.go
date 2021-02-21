package esme

import (
	"log"
	"net/http"
)

/*
Serve method takes a port and route config file(s) as arguments.
It is responsible for parsing the route config files(s), generate routes, set authentication and
start the server at the specified port.
*/
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
		for _, group := range c.RouteGroups {
			for _, endpoint := range group.Endpoints {
				endpoint.Auth = group.Auth
				configMap[getRouteMapKey(endpoint.Method, endpoint.Url)] = endpoint

				log.Printf("added route %s %s", endpoint.Method, endpoint.Url)
			}
		}
	}

	routeConfigMap = configMap
}

func getRouteMapKey(method string, url string) string {
	return method + "::" + url
}
