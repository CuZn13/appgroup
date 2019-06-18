BIN_DIR=_output/bin
RELEASE_VER=v0.0.1
REPO_PATH=github.com/cuzn/appgroup
GitSHA=`git rev-parse HEAD`
Date=`date "+%Y-%m-%d %H:%M:%S"`
REL_OSARCH="linux/amd64"
LD_FLAGS=" \
    -X '${REPO_PATH}/pkg/version.GitSHA=${GitSHA}' \
    -X '${REPO_PATH}/pkg/version.Built=${Date}'   \
    -X '${REPO_PATH}/pkg/version.Version=${RELEASE_VER}'"

appgroup: init
	go build -o=${BIN_DIR}/appgroup-controller ./cmd/appgroup

verify: generate-code
	hack/verify-gofmt.sh
	hack/verify-goimports.sh
	hack/verify-golint.sh
	hack/verify-gencode.sh

init:
	mkdir -p ${BIN_DIR}

generate-code:
	go build -o ${BIN_DIR}/deepcopy-gen ./cmd/deepcopy-gen/
	${BIN_DIR}/deepcopy-gen -i ./pkg/apis/scheduling/v1alpha1/ -O zz_generated.deepcopy
	${BIN_DIR}/deepcopy-gen -i ./pkg/apis/scheduling/v1alpha2/ -O zz_generated.deepcopy

rel_bins:
	go get github.com/mitchellh/gox
	CGO_ENABLED=0 gox -osarch=${REL_OSARCH} -ldflags ${LD_FLAGS} \
	-output=${BIN_DIR}/{{.OS}}/{{.Arch}}/kube-batch ./cmd/kube-batch

clean:
	rm -rf _output/
	rm -f kube-batch
