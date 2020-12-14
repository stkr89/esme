# Embeddable Service Mocking Engine (ESME)

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
      basic:
        username: username
        password: password
  - url: /user
    method: PUT
    status_code: 201
    body:
      firstName: Foo
      lastName: Bar
    response:
      - firstName: Foo
        lastName: Bar
    auth:
      basic:
        username: username
        password: password
  - url: /user/1
    method: DELETE
    status_code: 200
    auth:
      basic:
        username: username
        password: password


```