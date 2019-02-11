SOURCES=./
BINS=bin

dep:
	dep init

.PHONY: clean
clean:
	rm -rf ${BINS}/pseudoglobals

example:
	go run ${SOURCES}/examples/main.go

.DEFAULT_GOAL := test
test: example
