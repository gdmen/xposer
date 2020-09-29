build: build-jwtgen build-ping build-pong

build-jwtgen:
	go build -o bin/jwtgen garymenezes.com/xfinity-xposer/jwtgen

build-ping:
	go build -o bin/ping garymenezes.com/xfinity-xposer/ping

build-pong:
	# Ubuntu 18.04
	GOOS=linux GOARCH=amd64 go build -o bin/pong garymenezes.com/xfinity-xposer/pong

build-pong-test:
	go build -o bin/test_pong garymenezes.com/xfinity-xposer/pong

test: test-jwtgen test-ping test-pong

test-jwtgen: build-jwtgen
	go test garymenezes.com/xfinity-xposer/jwtgen

test-ping: build-ping
	go test garymenezes.com/xfinity-xposer/ping

test-pong: build-pong-test
	go test garymenezes.com/xfinity-xposer/pong

clean:
	rm -r bin/*

deploy-pong:
	./pong/deploy.sh gmenezes@35.212.165.198 /Volumes/sensitive/keys/id_garymenezes.pub

