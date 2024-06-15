compile-proto:
	protoc --go_out=server/ --go-grpc_out=server/ server/proto/*.proto
compile-server:
	mkdir -p bin && \
	cd server/ && \
	go build -v -o server_executable computacao-distribuida && \
	cd .. && \
	mv server/server_executable .
