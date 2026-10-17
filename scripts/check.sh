#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")/.."
for module in authentication-service broker-service listener-service logger-service mail-service; do
  printf '\nChecking %s\n' "$module"
  (cd "$module"; go test ./...; go build ./...)
done
