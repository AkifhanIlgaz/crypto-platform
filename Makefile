gen-crypto-service:
	@protoc \
		--proto_path=shared/proto "shared/proto/crypto/crypto.proto" \
		--go_out=shared/proto/ --go_opt=paths=source_relative \
  	--go-grpc_out=shared/proto/ --go-grpc_opt=paths=source_relative
