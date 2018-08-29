.DEFAULT_GOAL := server

.PHONY: proto
proto:
	@./gradlew generateProto

.PHONY: build
build:
	@./gradlew installDist

.PHONY: server
server: build
	@./build/install/server-java/bin/hello-world-server
