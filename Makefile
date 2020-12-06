docker/build:
				docker-compose build turdus

docker/run/api:
				docker-compose run --service-port turdus api

local/run/api:
				go run main.go api

.PHONY: install
install:
			go get -t ./... 

.PHONY: lint
lint:
			golint -min_confidence 0 ./...

.PHONY: test
test:
				go test ./...
