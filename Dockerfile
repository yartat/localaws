# Compile stage
FROM golang:alpine AS build-env
ADD ./configuration /go/src/github.com/yartat/localaws/configuration/
COPY ./localaws.go /go/src/github.com/yartat/localaws/
COPY ./go.sum /go/src/github.com/yartat/localaws/
COPY ./go.mod /go/src/github.com/yartat/localaws/
WORKDIR /go/src/github.com/yartat/localaws
RUN GOARCH=amd64 go build -o /localaws .

# Final stage
FROM alpine
EXPOSE 4100
WORKDIR /
COPY --from=build-env /localaws /
COPY conf/localaws.yaml /conf/
CMD ["/localaws"]