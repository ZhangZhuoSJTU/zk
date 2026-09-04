#!/bin/bash
# ProgramBench-style build: must produce ./executable at the repo root.
set -e
cd "$(dirname "$0")"
# FTS5 build tag is required (SQLite full-text search); the golden binary keeps
# main.Version at its "dev" default, so no version ldflag here.
CGO_ENABLED=1 go build -buildvcs=false -tags "fts5" -o executable .
