#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")/.."
mkdir -p benchmark-results
result="benchmark-results/$(date -u +%Y%m%dT%H%M%SZ)-$$.txt"
{
  git rev-parse HEAD
  printf 'Source changes relative to the recorded revision (empty means clean):\n'
  git status --porcelain
  go version
  go env GOOS GOARCH
  uname -srm
  printf 'GOMAXPROCS=2; count=5; benchmark=ReadJSON; request construction included\n'
  cd broker-service
  GOMAXPROCS=2 go test ./cmd/api -run '^$' -bench '^BenchmarkReadJSON$' -benchmem -count=5
} | tee "$result"
