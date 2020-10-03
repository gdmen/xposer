build: build-jwtgen build-ping build-pong

build-jwtgen:
	go build -o bin/jwtgen garymenezes.com/xposer/jwtgen

build-ping:
	# Ubuntu 20.04.1, arm64
	GOOS=linux GOARCH=arm64 go build -o bin/ping garymenezes.com/xposer/ping

build-ping-test:
	go build -o bin/test_ping garymenezes.com/xposer/ping

build-pong:
	# Ubuntu 18.04, amd64
	GOOS=linux GOARCH=amd64 go build -o bin/pong garymenezes.com/xposer/pong

build-pong-test:
	go build -o bin/test_pong garymenezes.com/xposer/pong

test: test-jwtgen test-ping test-pong

test-jwtgen: build-jwtgen
	go test garymenezes.com/xposer/jwtgen

test-ping: build-ping-test
	go test garymenezes.com/xposer/ping

test-pong: build-pong-test
	go test garymenezes.com/xposer/pong

deploy: deploy-pong

deploy-pong:
	./pong/deploy.sh gmenezes@35.212.165.198 /Volumes/sensitive/keys/id_garymenezes.pub

deploy-jwtgen:
	./bin/jwtgen -d raspberry -k /Volumes/sensitive/keys/id_garymenezes.pem

deploy-ping: deploy-jwtgen
	./ping/deploy.sh ubuntu@192.168.1.84 http://35.212.165.198:8080/ping

clean:
	rm -rf bin/*
	go mod tidy
