# Run a local node

Make sure you are in project src folder

## First init

`$ go mod download`

`$ go mod tidy`

Start node

`$ go run main.go`

Make test request
`$ curl http://localhost:8080/ping`

## FAQ

| Error | Solution |
| ----- | -------- |
| go: could not create module cache: mkdir /go: read-only file system | Make sure your $GOPATH variable is set to <User_dir>/go |

# Run and watch for changes

[just](https://github.com/casey/just) in conjunction to [entr](https://github.com/eradman/entr) can be used to run commands to watch for code changes. Currently only tested on MacOS.

To run all tests and watch for changes:

`just src/test`

To run all tests and watch for changes:

`just src/watch`

# Deploy full network:

`docker compose up -d`

To test ping between containers:

`docker exec lab_0-node-1 ping -c 4 lab_0-node-2`
