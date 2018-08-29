go-proto:
	@protoc --go_out=plugins=grpc:client-go proto/helloworld.proto

java-proto:
	@cp -r ./proto ./server-java/src/main/proto
