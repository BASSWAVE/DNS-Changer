proto-gen:
	protoc -I ./proto \
		   --go_out ./api  \
		   --go-grpc_out ./api \
		   --grpc-gateway_out ./api \
		   --openapiv2_out ./swagger \
		   ./proto/dns/dns.proto
