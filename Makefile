# export PATH="$PATH:$(go env GOPATH)/bin"

category:
	protoc --go_out=. --go-grpc_out=. proto/category/v1/category.proto
course:
	protoc --go_out=. --go-grpc_out=. proto/course/v1/course.proto

protoc: category course
																		