build:
	go build -o .bin/main cmd/main.go
run: build
	 ./.bin/main

swag:
	swag init -g cmd/main.go


getblock_build:
	docker build -t getblock .

getblock:getblock_build
		 docker run -dt --publish 8000:8000 --name getblock getblock