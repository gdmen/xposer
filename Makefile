build: build-keygen build-ping build-pong

build-keygen:
	go build -o bin/keygen garymenezes.com/xfinity-xposer

build-ping:
	go build -o bin/ping garymenezes.com/xfinity-xposer/ping

build-pong:
	# Ubuntu 18.04
	GOOS=linux GOARCH=amd64 GIN_MODE=release go build -o bin/pong garymenezes.com/xfinity-xposer/pong-server

clean:
	rm -r bin/*

deploy-pong:
	./pong-server/deploy.sh gmenezes@35.212.165.198 /Volumes/sensitive/keys/id_garymenezes.pub

