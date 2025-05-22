#!/usr/bin/env bash
set -e
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
