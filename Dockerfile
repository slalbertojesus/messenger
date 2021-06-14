
FROM golang:1.16

WORKDIR /app/
ENV GO111MODULE=on
COPY /server/src/go.mod .
COPY /server/src/go.sum .
COPY /server/src/grpc-server/ .
COPY /server/src/server.go .
RUN go mod download
CMD [ "go", "run", "server.go" ]
