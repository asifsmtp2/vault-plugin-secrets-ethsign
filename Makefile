VGO=go # Set to vgo if building in Go 1.10
BINARY_NAME=ethsign.exe
SRC_GOFILES := $(shell find . -name '*.go' -print)
.DELETE_ON_ERROR:

all: build
test: deps
		$(VGO) test  ./... -cover -coverprofile=coverage.txt -covermode=atomic
ethsign: ${SRC_GOFILES}
		
		$(VGO) build -o ${BINARY_NAME} -ldflags="-X 'main.Version=v1.0.0'" -tags=prod -v
build: ethsign
clean: 
		$(VGO) clean
		rm -f ${BINARY_NAME}
deps:
		$(VGO) get
