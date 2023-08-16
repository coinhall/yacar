#!/bin/bash

set -e

cd .github/scripts
ROOT_DIR=../.. go run cmd/yacar_ci.go

