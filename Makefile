bin/iam-service: $(shell find cmd internal generated -type f)
	go build -o bin/iam-service cmd/iam-service/main.go

generated/jsonmultiplex/iam/v0/iam.pb.go: $(shell find proto -type f)
	protoc \
		--proto_path=proto \
		--go_out=plugins=grpc:generated \
		jsonmultiplex/iam/v0/iam.proto

generated/jsonmultiplex/iam/v0/iam.pb.gw.go: $(shell find proto -type f)
	protoc \
		--proto_path=proto \
		--grpc-gateway_out=logtostderr=true:generated \
		jsonmultiplex/iam/v0/iam.proto
