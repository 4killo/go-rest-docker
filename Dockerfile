# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest


RUN mkdir -p /go/src/github.com/4killo/go-rest-docker
WORKDIR /go/src/github.com/4killo/go-rest-docker
# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/4killo/go-rest-docker

RUN go-wrapper download
RUN go-wrapper install

# Set the PORT environment variable inside the container
ENV PORT 80

# Build the service command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/4killo/go-rest-docker

# Run the service command by default when the container starts.
ENTRYPOINT /go/bin/go-rest-docker

# Document that the service listens on port 80.
EXPOSE 80
