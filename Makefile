all: build-image run

get-depends:
	go get -v -t -d ./...

build-code:
	go build

build-image:
	docker build -t docker-linter:latest -f Dockerfile .

run:
	docker rm -f docker-linter || true
	docker run -itd --name docker-linter docker-linter:latest
	docker exec -it docker-linter bash
