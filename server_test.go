package esme

import (
	"encoding/json"
	"github.com/stkr89/skeleton"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_getAuthBasic_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_getAuthBasic_success(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{
			Basic: &skeleton.AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)

	respBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var data []map[string]interface{}
	err = json.Unmarshal(respBytes, &data)
	assert.NoError(t, err)
	assert.Len(t, data, 2)
}

func Test_postAuthBasic_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/user",
		Method:  http.MethodPost,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_postAuthBasic_noBody_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/user",
		Method:  http.MethodPost,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{
			Basic: &skeleton.AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_postAuthBasic_success(t *testing.T) {
	r := skeleton.Request{
		Url:    "http://localhost:8080/auth/basic/user",
		Method: http.MethodPost,
		Body: map[string]string{
			"firstName": "foo",
			"lastName":  "bar",
		},
		Timeout: 10,
		Auth: &skeleton.Auth{
			Basic: &skeleton.AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	respBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var data []map[string]interface{}
	err = json.Unmarshal(respBytes, &data)
	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, data[0]["firstName"], "foo")
	assert.Equal(t, data[0]["lastName"], "bar")
	assert.Equal(t, data[0]["id"], "3")
}

func Test_putAuthBasic_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/user",
		Method:  http.MethodPut,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_putAuthBasic_noBody_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/user",
		Method:  http.MethodPut,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{
			Basic: &skeleton.AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_putAuthBasic_invalidCredentials_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/user",
		Method:  http.MethodPut,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{
			Basic: &skeleton.AuthBasic{
				Username: "username",
				Password: "passR",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_putAuthBasic_success(t *testing.T) {
	r := skeleton.Request{
		Url:    "http://localhost:8080/auth/basic/user",
		Method: http.MethodPut,
		Body: map[string]string{
			"firstName": "Foo",
			"lastName":  "Bar",
		},
		Timeout: 10,
		Auth: &skeleton.Auth{
			Basic: &skeleton.AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	respBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var data []map[string]interface{}
	err = json.Unmarshal(respBytes, &data)
	assert.NoError(t, err)
	assert.Len(t, data, 1)
	assert.Equal(t, data[0]["firstName"], "Foo")
	assert.Equal(t, data[0]["lastName"], "Bar")
	assert.Equal(t, data[0]["id"], "3")
}

func Test_deleteAuthBasic_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/user/1",
		Method:  http.MethodDelete,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_deleteAuthBasic_success(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/basic/user/1",
		Method:  http.MethodDelete,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{
			Basic: &skeleton.AuthBasic{
				Username: "username",
				Password: "password",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func Test_getAuthBearerToken_missingToken_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/bearer_token/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_getAuthBearerToken_invalidToken_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/bearer_token/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth:    &skeleton.Auth{BearerToken: &skeleton.AuthBearerToken{Token: "t"}},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func Test_getAuthBearerToken_success(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/bearer_token/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{
			BearerToken: &skeleton.AuthBearerToken{
				Token: "token",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func Test_getAuthCustomHeader_missingHeader_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/custom/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_getAuthCustomHeader_invalidHeader_failure(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/custom/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{Custom: map[string]string{
			"my_custom_header": "val",
		}},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func Test_getAuthCustomHeader_success(t *testing.T) {
	r := skeleton.Request{
		Url:     "http://localhost:8080/auth/custom/users",
		Method:  http.MethodGet,
		Body:    nil,
		Timeout: 10,
		Auth: &skeleton.Auth{
			Custom: map[string]string{
				"my_custom_header": "header val",
			},
		},
	}

	resp, err := skeleton.Send(&r)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
