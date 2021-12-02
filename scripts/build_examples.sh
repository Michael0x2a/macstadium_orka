#!/usr/bin/env bash

set -euo pipefail

mkdir -p ./build
rm -rf ./build/*

go build -o ./build ./examples/*

