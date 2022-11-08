NAMESPACE  := logicmonitor
REPOSITORY := lm-sdk-go
VERSION    ?= $(shell git describe --tags --always --dirty)
ifeq ($(VERSION),)
VERSION := latest
endif


default: build


build:
	docker build --build-arg VERSION=$(VERSION) -t $(NAMESPACE)/$(REPOSITORY):$(VERSION) .
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):$(VERSION) -r /go/src/github.com/logicmonitor/lm-sdk-go/client /out/
	docker run --rm -v "$(shell pwd)":/out --entrypoint=cp $(NAMESPACE)/$(REPOSITORY):$(VERSION) -r /go/src/github.com/logicmonitor/lm-sdk-go/models /out/
	go mod init github.com/logicmonitor/lm-sdk-go
	go mod tidy
