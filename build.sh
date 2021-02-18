#!/bin/bash

BUILD_DIR="build"
BIN_NAME="pvs_differ"

SRC_FILES="main.go alertion.go parser.go"

! [ -d "${BUILD_DIR}" ] && mkdir ${BUILD_DIR}

gofmt -w ${SRC_FILES}
go build -o "./${BUILD_DIR}/${BIN_NAME}" ${SRC_FILES}