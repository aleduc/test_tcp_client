FROM golang:1.17.6-alpine3.15 as builder

WORKDIR /workspace
COPY go.mod .
RUN go mod download

COPY cmd/test_tcp_client/main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o test_tcp_client main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/test_tcp_client .

ENTRYPOINT ["/test_tcp_client"]


