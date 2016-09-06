all: whisper-client whisper-server

clean:
	rm -rf bin/

whisper-server: server/*.go
	go build -o bin/whisper-server server/*.go

whisper-client: client/*.go
	go build -o bin/whisper-client client/*.go
