## Builder Image
FROM golang:1.13-stretch as builder
ENV GO111MODULE=on
COPY . /c4-payment
WORKDIR /c4-payment
RUN go get github.com/google/wire/cmd/wire
RUN wire ./...
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application cmd/c4/main.go cmd/c4/wire_gen.go

## Run Image
FROM scratch
COPY --from=builder /c4-payment/bin/application application
EXPOSE 8080
ENTRYPOINT ["./application"]