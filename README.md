# govtech_onecv_assignment

Name: Tiang Hui Zheng

Assignment Received: 9 Oct 2023

Assignment Deadline: 14 Oct 2023

Link to deployment: http://heyzec.mooo.com:8000


## Project structure

```
cmd/
├── server
└── ...
internal/
├── routes
├── handlers
├── dataaccess
├── models
├── views
├── params
├── errors
└── helpers/
    ├── api
    ├── arrayutils
    ├── emailutils
    ├── json
    └── ...
tests/
└── ...
migrations
seeds
```

|Folder|Description|
|---|---|
|cmd|Entry points to the app, including server and migration tooling|
|internal/handlers|Controller (C in MVC, kinda)|
|internal/dataaccess|Layer between handlers and models|
|internal/models|GORM models (M in MVC)|
|internal/views|Define shape of API responses (V in MVC)|
|internal/params|Define shape of API requests|
|internal/errors|Custom errors definitions|
|internal/helpers|Many utility stuff, including wrappers, database abstractions, JSON handling|
|tests|Tests, this folder mirrors internal/|


## Getting started

1. Create the development database.
    ```
    psql -c "CREATE DATABASE govtech_assignment"
    ```

1. Fill up necessary environment variables
    ```
    cp .env.template .env.development
    vim .env.development
    ```

1. Migrate and seed the database.
    ```
    make migrate
    make seed
    ```

1. Start the server. This uses reflex if available to reload server on file changes.
    ```
    make run
    ```

## Running tests
1. Create the development database.
    ```
    psql -c "CREATE DATABASE govtech_assignment"
    ```

1. Migrate and seed the testing database.
    ```
    make migrate-test
    make seed-test
    ```

1. Run each test individually. Note that the order in which tests are executed may affect result.
This is definitely unideal, and it is caused by not having a mock storage implemented.
    ```
    go test ./tests/arrayutils/intersection_test.go
    ```