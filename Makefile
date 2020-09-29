test: test-jwtgen

test-jwtgen: build-jwtgen
	go test garymenezes.com/xfinity-xposer/jwtgen

build: build-jwtgen build-ping build-pong

build-jwtgen:
	go build -o bin/jwtgen garymenezes.com/xfinity-xposer/jwtgen

build-ping:
	go build -o bin/ping garymenezes.com/xfinity-xposer/ping

build-pong:
	# Ubuntu 18.04
	GOOS=linux GOARCH=amd64 GIN_MODE=release go build -o bin/pong garymenezes.com/xfinity-xposer/pong-server

clean:
	rm -r bin/*

deploy-pong:
	./pong-server/deploy.sh gmenezes@35.212.165.198 /Volumes/sensitive/keys/id_garymenezes.pub

