#!/bin/bash

ROOT_DIR=$(git rev-parse --show-toplevel)
cd $ROOT_DIR/.github/scripts
ROOT_DIR=$ROOT_DIR go run cmd/yacar_ci.go
