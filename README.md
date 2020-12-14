# Embeddable Service Mocking Engine (ESME)
ESME allows you to define destination services in `yaml` format.    
## Usage
Sample route config file
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