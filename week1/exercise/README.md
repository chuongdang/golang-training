# Exercise 1 - User Service

### 1. Requirements
 Based on the lesson and the code in [./webservice](./webservice), develop a fully featured service, which have:

 - Get list of user
    - Endpoint: `/users`
    - Method: `GET`
    - Response:
        - Response Code: `200`
        - Response Body: User list as JSON string

 - Get user by id
    - Endpoint: `/users/{id}`
    - Method: `GET`
    - Response:
        - Response Code: `200`
        - Response Body: User as JSON string

 - Add user
    - Endpoint: `/users`
    - Method: `POST`
    - Request Body: User data as JSON string
    - Response:
        - Response Code: `200`
        - Response Body: User as JSON string (with ID)

 - Update user
    - Endpoint: `/users`
    - Method: `PUT`
    - Request Body: User data as JSON string (with ID)
    - Response:
        - Response Code: `200`
        - Response Body: User as JSON string (with ID)

 - Delete user
    - Endpoint: `/users/{id}`
    - Method: `DELETE`
    - Response:
        - Response Code: `200`

### 2. Notices

- Not every request should return response code `200`, try to handle error when something wrong.
- For routing with parameters, for example `users/{id}`, kindly check the document in [https://github.com/gin-gonic/gin#parameters-in-path](https://github.com/gin-gonic/gin#parameters-in-path)
- The document for `sqlx` is located in [https://jmoiron.github.io/sqlx/](https://jmoiron.github.io/sqlx/)
