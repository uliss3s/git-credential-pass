#!/bin/bash
ROOT_DIR=$(pwd)
cd cmd/git-credential-pass
go build -ldflags="-s -w" -o $ROOT_DIR/target/git-credential-pass
cp $ROOT_DIR/target/git-credential-pass ~/.local/bin/git-credential-pass