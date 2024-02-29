$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$ go env
$ echo $PATH
$ export PATH="$PATH:$(go env GOPATH)/bin"

git tag v0.0.1
git push origin v0.0.1