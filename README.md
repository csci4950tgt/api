# API for Vigilante Web Heist

## How to run

Prerequisites:

* Golang is installed
* Docker is installed

#### Steps to run:

**With Golang**
1. Run `go build` to compile the executable from the file "main.go"
2. Run the executable named "api" (ex: if using mac run `./api`)

**With Docker**
1. Run `docker run -p 8080:8080 csci4950tgt/api`

#### Steps to rebuild Dockerfile and push new image to Docker repo

1. Run `docker build -t csci4950tgt/api .` to build the image
2. Run `docker push csci4950tgt/api` to push new image to Docker repo

#### Steps to test:

1. Run `go test`
