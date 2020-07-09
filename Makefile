.PHONY: build clean deploy

build:
	go get -v all
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o bin/main  intern/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy --verbose
