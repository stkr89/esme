package esme

import (
	"encoding/base64"
	"log"
	"net/http"
)

const (
	AuthorizationBasicError       = "authorization basic header is invalid"
	AuthorizationBearerTokenError = "authorization bearer token is invalid"
	CustomHeaderError             = "custom header(s) is invalid"
)

func checkAuthorization(r *http.Request, route *route) (string, int) {
	for _, f := range getAuthCheckers() {
		if route.Auth == nil {
			continue
		}

		errStr, statusCode := f(r, route)
		if errStr != "" {
			return errStr, statusCode
		}
	}

	return "", 0
}

func getAuthCheckers() []func(r *http.Request, route *route) (string, int) {
	return []func(r *http.Request, route *route) (string, int){
		checkBasicAuth,
		checkBearerTokenAuth,
		checkCustomHeaders,
	}
}

func checkCustomHeaders(r *http.Request, route *route) (string, int) {
	if route.Auth != nil && route.Auth.Custom != nil {
		custom := route.Auth.Custom
		for k, v := range custom {
			headerVal := r.Header.Get(k)

			if headerVal != v {
				log.Println(CustomHeaderError)
				return CustomHeaderError, http.StatusBadRequest
			}
		}
	}

	return "", 0
}

func checkBearerTokenAuth(r *http.Request, route *route) (string, int) {
	header := r.Header.Get("Authorization")

	if route.Auth != nil && route.Auth.BearerToken != nil {
		bearer := route.Auth.BearerToken
		if header == "" || header != "Bearer "+bearer.Token {
			log.Println(AuthorizationBearerTokenError)
			return AuthorizationBearerTokenError, http.StatusUnauthorized
		}
	}

	return "", 0
}

func checkBasicAuth(r *http.Request, route *route) (string, int) {
	header := r.Header.Get("Authorization")

	if route.Auth != nil && route.Auth.Basic != nil {
		basic := route.Auth.Basic
		if header == "" || header != basicAuthHeaderValue(basic.Username, basic.Password) {
			log.Println(AuthorizationBasicError)
			return AuthorizationBasicError, http.StatusUnauthorized
		}
	}

	return "", 0
}

func basicAuthHeaderValue(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
