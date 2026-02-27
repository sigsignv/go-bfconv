.PHONY: build package clean

BIN ?= bootstrap
GOOS ?= linux
GOARCH ?= amd64
CGO_ENABLED ?= 0
LDFLAGS ?= -s -w

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) \
	go build -trimpath -ldflags "$(LDFLAGS)" -o $(BIN) .

package: build
	zip -j lambda-handler.zip $(BIN)

clean:
	$(RM) $(BIN) lambda-handler.zip
