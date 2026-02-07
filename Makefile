
proto:
	protoc --go_out=./types ./types/proto/*.proto

clean_proto:
	@rm -f ./types/proto/*.pb.go