build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/amogower/shippy/auth-service \
		proto/auth/auth.proto
	# docker build -t auth-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns auth-service
