FROM alpine:3.20

# Install system dependencies
RUN apk add --no-cache \
    bash \
    curl \
    iputils-ping \
    make \
    musl-dev

# Install Go 1.23
RUN curl -LO https://go.dev/dl/go1.23.1.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz && \
    rm go1.23.1.linux-amd64.tar.gz

# Configure Go
ENV GOPATH /go
ENV PATH /go/bin:/usr/local/go/bin:$PATH
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
