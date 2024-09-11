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

# Deploy full network:

`docker compose up -d`

To test ping between containers:

`docker exec lab_0-node-1 ping -c 4 lab_0-node-2`
