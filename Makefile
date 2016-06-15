all: client server

clean:
	rm -rf bin/

server: server/server.go
	go build -o bin/server server/*.go

client: client/client.go
	go build -o bin/client client/*.go