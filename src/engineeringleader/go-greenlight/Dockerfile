# Specifies a parent image
FROM golang:1.19
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
 
 
# Builds your app with optional configuration
RUN go build -o ./out/go-sample-app ./cmd/api
# Tells Docker which network port your container listens on
CMD ["./out/go-sample-app"]
