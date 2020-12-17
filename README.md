# Embeddable Service Mocking Engine (ESME)
ESME allows you to define the destination mock services in `yaml` format. These 
services can then simply be consumed by test cases to get the expected behaviour. 
## Usage
Here is a sample `route-config.yml` file that can be processed by ESME
```yaml
routes:
  - url: /users/1
    method: GET
    status_code: 200
    response:
      - firstName: jane
        lastName: doe
        id: 1
      - firstName: john
        lastName: doe
        id: 2
    auth:
      basic:
        username: username
        password: password
  - url: /user
    method: POST
    status_code: 201
    body:
      firstName: foo
      lastName: bar
    response:
      - firstName: foo
        lastName: bar
    auth:
      bearer_token:
        token: token
  - url: /user/1
    method: DELETE
    status_code: 200
    auth:
      custom:
        my_header: value
```
Start a mock server using above `route-config.yml` file
```go
package main

import (
	"github.com/stkr89/esme-runner/esme"
)

func main() {
    esme.Serve("8080", "./route-config.yml")
}
```

> You can also provide multiple route configs as arguments to `Serve` method.  

Let's break down this file to understand what each component means.
## Routes
`routes` contains the list of routes which need to be mocked. ESME supports adding 
routes to multiple files which can represent different services running 
simultaneously.
### URL
`url` defines the route that need to be mocked.
### Method
`method` defines the http method associated with an url.
### Status Code
`status_code` defines the http status code that needs to be returned from the 
endpoint.
### Response
```yaml
response:
  - firstName: jane
    lastName: doe
    id: 1
  - firstName: john
    lastName: doe
    id: 2
```
`response` defines a list of objects that the endpoint returns on success.
### Auth
`auth` defines the authentication scheme required for an endpoint. Each `url` can
have its own authentication scheme. ESME supports following authentication schemes:
#### Basic
```yaml
auth:
  basic:
    username: username
    password: password
```
`basic` authentication checks for a header field in the form of 
`Authorization: Basic <credentials>`, where `<credentials>` is the Base64 
encoding of username and password joined by a single colon `:`.
#### Bearer Token
```yaml
auth:
  bearer_token:
    token: token
```
`bearer_token` authentication checks for a header field in the form of 
`Authorization: Bearer <token>`.
#### Custom
```yaml
auth:
  custom:
    my_header_1: value1
    my_header_2: value2
```
`custom` authentication checks for headers `my_header_1` and `my_header_2` 
with values `value1` and `value2` respectively.