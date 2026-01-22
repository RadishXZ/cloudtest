#!/bin/bash

PROJ_ROOT=$(dirname "${BASH_SOURCE[0]}")
OUTPUT_DIR=${PROJ_ROOT}/_output
VERSION_PACKAGE=github.com/RadishXZ/cloudtest/pkg/version

if [[ -z "${VERSION}" ]];then
    VERSION=$(git describe --tags --always --match='v*')
fi

GIT_TREE_STATE="dirty"

is_clean=$(git status --porcelain 2>/dev/null)
if [[ -z ${is_clean} ]];then
    GIT_TREE_STATE="clean"
fi

GIT_COMMIT=$(git rev-parse HEAD)

GO_LDFLAGS="-X ${VERSION_PACKAGE}.gitVersion=${VERSION} \
    -X ${VERSION_PACKAGE}.gitCommit=${GIT_COMMIT} \
    -X ${VERSION_PACKAGE}.gitTreeState=${GIT_TREE_STATE} \
    -X ${VERSION_PACKAGE}.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"

go build -v -ldflags "${GO_LDFLAGS}" -o ${OUTPUT_DIR}/fg-apiserver -v cmd/fg-apiserver/main.go