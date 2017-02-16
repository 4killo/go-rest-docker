# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest


RUN mkdir -p /go/src/lb.com/customerservice
WORKDIR /go/src/lb.com/customerservice
# Copy the local package files to the container's workspace.
COPY . /go/src/lb.com/customerservice

RUN go-wrapper download
RUN go-wrapper install

# Set the PORT environment variable inside the container
ENV PORT 8080

# Build the customerservice command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install lb.com/customerservice

# Run the customerservice command by default when the container starts.
ENTRYPOINT /go/bin/customerservice

# Document that the service listens on port 8080.
EXPOSE 8080
