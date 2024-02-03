package esme

import (
	"encoding/base64"
	"errors"
	"log/slog"
	"net/http"
)

const (
	AuthorizationBasicError       = "authorization basic header is invalid"
	AuthorizationBearerTokenError = "authorization bearer token is invalid"
	CustomHeaderError             = "custom header(s) is invalid"
)

func checkAuthorization(r *http.Request, route *route) error {
	for _, f := range getAuthCheckers() {
		err := f(r, route)
		if err != nil {
			slog.Error(err.Error())
			return err
		}
	}

	return nil
}

func getAuthCheckers() []func(r *http.Request, route *route) error {
	return []func(r *http.Request, route *route) error{
		checkBasicAuth,
		checkBearerTokenAuth,
		checkCustomHeaders,
	}
}

func checkCustomHeaders(r *http.Request, route *route) error {
	if route.Auth.Custom != nil {
		custom := route.Auth.Custom

		for k, v := range custom {
			headerVal := r.Header.Get(k)

			if headerVal != v {
				return errors.New(CustomHeaderError)
			}
		}
	}

	return nil
}

func checkBearerTokenAuth(r *http.Request, route *route) error {
	header := r.Header.Get("Authorization")

	if route.Auth.BearerToken != nil {
		bearer := route.Auth.BearerToken
		if header == "" || header != "Bearer "+bearer.Token {
			return errors.New(AuthorizationBearerTokenError)
		}
	}

	return nil
}

func checkBasicAuth(r *http.Request, route *route) error {
	header := r.Header.Get("Authorization")

	if route.Auth.Basic != nil {
		basic := route.Auth.Basic
		if header == "" || header != basicAuthHeaderValue(basic.Username, basic.Password) {
			return errors.New(AuthorizationBasicError)
		}
	}

	return nil
}

func basicAuthHeaderValue(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
