BINARY=no-name
TEST=test.sh

.DEFAULT_GOAL: ${BINARY}

buildandrun: clean build run

check: test lint vet

lint:
	@go list ./...  | grep -v /vendor/ |  xargs -L1 golint -set_exit_status

vet:
	@go vet $(shell go list ./... | grep -v /vendor/)

test:
	@if [ -f ${TEST} ] ; then ./${TEST} ; fi

build:
	@go build -o ${BINARY} .

run:
	@if [ -f ${BINARY} ] ; then ./${BINARY} -d db.txt -f ip.txt -o html ; fi

install:
	@go install $(shell go list ./... | grep -v /vendor/)

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi