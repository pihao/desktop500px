#!/usr/bin/env bash

dev_dir=$GOPATH/src/github.com/pihao/desktop500px

fmt_path=github.com/pihao/desktop500px/src/...
go fmt $fmt_path
go test $fmt_path
