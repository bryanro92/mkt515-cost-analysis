SHELL = /bin/bash
COMMIT = $(shell git rev-parse --short HEAD)

install: build 
	chmod 555 ./cost-analysis
	mv ./cost-analysis /usr/local/bin/cost-analysis

build:
	go build . -o cost-analysis

run:
	go run . 

clean:  
	rm -f ./cost-analysis

.PHONY: build install run clean