FROM golang:1.22-alpine

RUN go version
ENV GOPATH=/

COPY ./ ./

EXPOSE ${HTTP_PORT}

RUN go mod download
RUN go build -o idler ./cmd/idler-service/main.go

CMD ["./idler"]