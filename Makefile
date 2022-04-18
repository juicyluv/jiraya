proto_out_path = internal/jiraya/interfaces/grpc_gw
protofile_path = api/proto/jiraya/*.proto
proto_path = api/proto/jiraya

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
		-I$(proto_path) \
        $(protofile_path)

.PHONY: doc
doc:
	protoc \
		--openapiv2_out=web/apidoc/v1 \
		--openapiv2_opt use_go_templates=true \
		--openapiv2_opt json_names_for_fields=false \
		--openapiv2_opt allow_delete_body=true \
		$(protofile_path)

.PHONY: build
build:
	go build -o ./bin/jiraya/jiraya.exe ./cmd/jiraya/main.go

.PHONY: run
run:
	build; bin/jiraya/jiraya.exe