# Specifies a parent image
FROM golang:1.19
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
 
 
# Builds your app with optional configuration
RUN go build -o docker-gs-ping
# Tells Docker which network port your container listens on

ENTRYPOINT ./docker-gs-ping