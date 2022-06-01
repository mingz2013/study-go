


help:
	echo 'help message'


SRC_DIR=./protos
DST_DIR=./gen

build_proto:
	#	protoc -I=$(SRC_DIR) --go_out=$(DST_DIR) *.proto
#	protoc -I=./protos --go_out=./gen test.proto
#	protoc -I=./protos --go_out=./gen tutorial.person.proto
#
#	protoc -I=./protos --go_out=./gen info.proto

	protoc --go_out=./gen protos/*.proto
	#protoc --go_out=plugins=grpc:./gen protos/*.proto
	protoc --go-grpc_out=./gen protos/*.proto

run:
	go run main.go