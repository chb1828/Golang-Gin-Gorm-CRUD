FROM golang:alpine3.14

WORKDIR /build

COPY go.mod go.sum server.go ./

RUN go mod download

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]
