#!/bin/bash

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | tr '[:upper:]' '[:lower:]')

BIN=".github/scripts/dist/yacar_ci_${OS}_${ARCH}"

echo $BIN

ROOT_DIR=$(git rev-parse --show-toplevel)
ROOT_DIR=${ROOT_DIR} ${BIN}
