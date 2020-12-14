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
Let's break down this file to understand what each component means.
## Routes
`routes` contains all the routes which need to be mocked. ESME supports adding 
routes to multiple files which can represent different services running 
simultaneously.