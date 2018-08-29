.DEFAULT_GOAL := server

build:
	@./gradlew installDist

server: build
	@./build/install/server-java/bin/hello-world-server
