#!/bin/bash --posix

ENV=$(uname -snr)
VER=$(cat .version)
UPT=$(date +"%Y/%m/%d %T %z")
TAG=$(echo $(git rev-parse --short HEAD)$([ -n "$(git status -s)"  ] && echo "-dev" || echo ""))

SOURCE="cmd/${1}/main.go"
BINARY="bin/${1}"

echo "making ${1}"
echo "Version:${TAG}"

go build -tags netgo                                \
    -installsuffix 'static'                         \
    -ldflags "                                      \
    -s -w                                        \
    -X '$(go list -m)/pkg/version.verStr=${VER}'    \
    -X '$(go list -m)/pkg/version.tagStr=${TAG}'    \
    -X '$(go list -m)/pkg/version.uptStr=${UPT}'    \
    -X '$(go list -m)/pkg/version.envStr=${ENV}'    \
    "                                               \
    -o ${BINARY} ${SOURCE}
