#!/usr/bin/env bash
set -e
go mod tidy
go build ./cmd/vnprider-node
go build ./cmd/vnprider-cli
