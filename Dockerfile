# # We specify the base image we need for our GO app
FROM golang:1.25.0-alpine as builder

# #Create /workspace directory within the image to hold our application source code
WORKDIR /workspace
USER root
RUN go install github.com/go-delve/delve/cmd/dlv@latest
# # We copy everything in the root directory into our /workspace directory
COPY . .

# # download Go modules and dependencies
RUN go mod download

# # Build the app with optional configuration
RUN go build -o encrypt .

# # We copy go files into our /workspace directory


# # tells Docker that the container listens on specified network ports at runtime
EXPOSE 2345

# # command to be used to execute when the image is used to start a container
CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2", "--log"]