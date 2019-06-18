BIN_DIR=_output/bin

build: clean init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=${BIN_DIR}/appgroup-controller

init:
	mkdir -p ${BIN_DIR}

clean:
	rm -rf _output/
