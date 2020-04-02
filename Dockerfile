# Using Multi-stage build process to minify final api container
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds

### STEP 1: build executable binary

# Based off the Golang Dockerfile example 
# https://hub.docker.com/_/golang

# Use the official Golang runtime as a parent image
FROM golang:latest AS builder

# Set the current working directory
WORKDIR /go/src/app

# Copy the current directory contents into the container
COPY . .

# Download and install necessary Golang dependencies
RUN go mod download

# Build api executable (to be used in next stage)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/api .

# download the cloudsql proxy binary
RUN wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O /go/bin/cloud_sql_proxy
RUN chmod +x /go/bin/cloud_sql_proxy

# copy the wrapper script and credentials
COPY run.sh /go/bin/run.sh
COPY credentials.json /go/bin/credentials.json


### STEP 2: Copy executable into a new smaller image
## Followed these comments to get a CloudSQL Proxy going to connect to postgres and let the api run
## https://medium.com/@petomalina/connecting-to-cloud-sql-from-cloud-run-dcff2e20152a

# Use the alpine linux runtime as a parent image
FROM alpine:latest  

# Set the current working directory
WORKDIR /root/

# add certificates
RUN apk --no-cache add ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin .

# Give the executable permission to run
RUN chmod +x ./api
RUN chmod +x ./run.sh

CMD ["./run.sh"]

# # Configure container to run as api executable
# ENTRYPOINT ["./api"]

# Expose port 8080 to the outside world
# EXPOSE 8080

