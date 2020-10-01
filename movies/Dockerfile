FROM golang:onbuild
#FROM golang:1.11.2 as dev

# Install dependencies


# copy the local package files to the container workspace
ADD . /go/src/github.com/cleverRonald/prueba_jhon_2

# Setting up working directory
WORKDIR /go/src/github.com/cleverRonald/prueba_jhon_2

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o movies .


# alpine image
#FROM alpine:3.9.2 as prod
# Setting up working directory
#WORKDIR /root/
# copy movies binary
#COPY --from=dev /go/src/github.com/cleverRonald/prueba_jhon_2 .
# Service listens on port 8080.
EXPOSE 8080
# Run the movies microservice when the container starts.
ENTRYPOINT ["./movies"]