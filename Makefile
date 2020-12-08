apps = 'gateway' 'registry' 'message' 'subscribe' 'cron' 'web'
.PHONY: run
run:
	for app in $(apps) ;\
	do \
		 go run ./cmd/$$app -f configs/$$app.yml  & \
	done
.PHONY: build
build:
	for app in $(apps) ;\
	do \
		GOOS=linux GOARCH="amd64" go build -o dist/$$app-linux-amd64 ./cmd/$$app/; \
	done
.PHONY: lint
lint:
	golint ./...
.PHONY: docker
docker-compose: build
	docker-compose -f deployments/docker-compose.yml up --build -d
.PHONY: proto
proto:
	protoc -I api/proto ./api/proto/* --go_out=api/proto
all: lint docker
