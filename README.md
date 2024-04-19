# Multi-platform demo

## Building or running the app

This project utilises custom tags to differentiate platform builds. Current supported tags are as follows:

* `evokeos`
* `f1`

### EvokeOS

To build the app for EvokeOS, run the following from the root directory:

```shell
go build -tags evokeos main.go
```

To run the app for EvokeOS, run the following from the root directory:

```shell
go run -tags evokeos main.go
```

### F1

To build the app for F1, run the following from the root directory:

```shell
go build -tags f1 main.go
```

To run the app for F1, run the following from the root directory:

```shell
go run -tags f1 main.go
```

## Unit testing

Since we use custom tags here, this means we would also need to customise the test commands

### EvokeOS

To run all the unit tests for EvokeOS, run the following the root directory:

```shell
go test -tags evokeos ./...
```

### F1

To run all the unit tests for F1, run the following the root directory:

```shell
go test -tags f1 ./...
```

## Run configurations

Run configurations for Goland users are readily available in the project. For the other IDEs you prefer,
feel free to add it in the VCS.