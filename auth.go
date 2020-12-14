package main

import (
	"encoding/base64"
	"log"
	"net/http"
)

const (
	AuthorizationBasicError       = "authorization basic header not found"
	AuthorizationBearerTokenError = "authorization bearer token not found"
	CustomHeaderError             = "custom header not found"
)

func checkAuthorization(w http.ResponseWriter, r *http.Request, route *routes) {
	for _, f := range getAuthCheckers() {
		errStr, statusCode := f(r, route)
		if errStr != "" {
			http.Error(w, errStr, statusCode)
		}
	}
}

func getAuthCheckers() []func(r *http.Request, route *routes) (string, int) {
	return []func(r *http.Request, route *routes) (string, int){
		checkBasicAuth,
		checkBearerTokenAuth,
		checkCustomHeaders,
	}
}

func checkCustomHeaders(r *http.Request, route *routes) (string, int) {
	custom := route.Auth.Custom

	if custom != nil {
		for k, v := range custom {
			headerVal := r.Header.Get(k)

			if headerVal != v {
				log.Println(AuthorizationBearerTokenError)
				return CustomHeaderError, http.StatusBadRequest
			}
		}
	}

	return "", 0
}

func checkBearerTokenAuth(r *http.Request, route *routes) (string, int) {
	bearer := route.Auth.BearerToken
	header := r.Header.Get("Authorization")

	if bearer != nil {
		if header == "" || header != "Bearer "+bearer.Token {
			log.Println(AuthorizationBearerTokenError)
			return AuthorizationBearerTokenError, http.StatusUnauthorized
		}
	}

	return "", 0
}

func checkBasicAuth(r *http.Request, route *routes) (string, int) {
	basic := route.Auth.Basic
	header := r.Header.Get("Authorization")

	if basic != nil {
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
