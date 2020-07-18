.PHONY: setup get-dependencies generate-proto-def generate-go-files
setup: get-dependencies generate-proto-def generate-go-files


get-dependencies:
	  go get -v -t -d ./...
	  go get -u github.com/golang/protobuf/protoc-gen-go
	  go get github.com/golang/mock/mockgen@v1.4.3
	  go get github.com/google/wire/cmd/wire
	  go install golang.org/x/lint/golint

generate-proto-def:
	protoc --go_out=plugins=grpc:. idl/job.proto

generate-go-files:
	go generate ./...