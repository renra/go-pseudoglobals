SOURCES=./
BINS=bin

dep:
	dep ensure

.PHONY: clean
clean:
	rm -rf ${BINS}/pseudoglobals

example:
	go run ${SOURCES}/examples/main.go

.DEFAULT_GOAL := test
test: example
