proto_out_path = internal/jiraya/interfaces/grpc_gw
proto_path = api/proto/jiraya/*.proto

.PHONY: proto
proto:
	protoc \
		--go_out=$(proto_out_path) \
        --go-grpc_out=$(proto_out_path) \
        --grpc-gateway_out=$(proto_out_path) \
		--grpc-gateway_opt=logtostderr=true \
		--grpc-gateway_opt=module=${module} \
		--grpc-gateway_opt=generate_unbound_methods=true \
		--grpc-gateway_opt=allow_delete_body=true \
		-Iapi/proto/jiraya \
        $(proto_path)

.PHONY: build
build:
	go build -o ./bin/jiraya/jiraya.exe ./cmd/jiraya/main.go

.PHONY: run
run:
	build; bin/jiraya/jiraya.exe