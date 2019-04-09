PROTOFILES := $(shell find . -name "*.proto" -type f)
BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

cold-start:
	wire && go build -o ${BINARY} && ./${BINARY}

start:
	go build -o ${BINARY} && ./${BINARY}


engine: clean
	wire && go build -o ${BINARY}

build: clean
	go build -o ${BINARY}

migrate: build
	./${BINARY} mysql_migrate

seed: build
	./${BINARY} mysql_seed

install: 
	go build -o ${BINARY}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t go-clean-arch .

run:
	docker-compose up -d

stop:
	docker-compose down

generate-proto:
	@echo "Generating protobuf"; \
	for FILE in $(PROTOFILES); do \
		echo "processing" $$FILE; \
		protoc --go_out=plugins=grpc:. $$FILE;\
	done;

.PHONY: clean install unittest build docker run stop vendor
