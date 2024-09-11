FROM alpine:3.20.2

# Install system dependencies
RUN apk add --no-cache \
    bash \
    curl \
    iputils-ping \
    make \
    musl-dev \
    go

# Configure Go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# Install app dependencies
WORKDIR /src
COPY src/go.mod src/go.sum ./
RUN go mod download

# Build the binary
COPY src/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

# Run the binary
EXPOSE 8080
ENTRYPOINT ["/app"]
