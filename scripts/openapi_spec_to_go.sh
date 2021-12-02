#!/usr/bin/env bash

set -euo pipefail

API_SPEC="./api/orka-macstadium.yaml"
CODEGEN="github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0"

go run "$CODEGEN" -package orka -generate types -o ./pkg/orka/types.gen.go "$API_SPEC"
go run "$CODEGEN" -package orka -generate client -o ./pkg/orka/client.gen.go "$API_SPEC"
