FROM golang:1.9.0 as builder
WORKDIR /go/src/github.com/amogower/shippy/auth-service
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service -a -installsuffix cgo main.go repository.go handler.go database.go token_service.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/amogower/shippy/auth-service/auth-service .
CMD ["./auth-service"]
