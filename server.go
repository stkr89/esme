package esme

import (
	"errors"
	"log/slog"
	"net/http"
)

/*
Serve method takes a port and route config file(s) as arguments.
It is responsible for parsing the route config files(s), generate routes, set authentication and
start the server at the specified port.
*/
func Serve(port string, paths ...string) error {
	routeConfig, err := getRouteConfig(paths)
	if err != nil {
		return err
	}

	setRoutes(routeConfig)
	launchServer(port)

	return nil
}

func launchServer(port string) {
	m := http.NewServeMux()
	s := http.Server{Addr: ":" + port, Handler: m}

	m.HandleFunc("/", handleAll)
	m.HandleFunc("/shutdown", handleShutdown(port, &s))

	slog.Info("starting ESME server", "port", port)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server error", "error", err)
	}
}

func setRoutes(configs []*config) {
	configMap := make(map[string]*route)

	for _, c := range configs {
		for _, group := range c.RouteGroups {
			for _, endpoint := range group.Endpoints {
				endpoint.Auth = group.Auth
				configMap[getRouteMapKey(endpoint.Method, endpoint.Url)] = endpoint

				slog.Info("added route %s %s", endpoint.Method, endpoint.Url)
			}
		}
	}

	routeConfigMap = configMap
}

func getRouteMapKey(method string, url string) string {
	return method + "::" + url
}
