generate-structs:
	mkdir -p pkg/user_v1
	protoc -I proto proto/sorm/sorm.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative
