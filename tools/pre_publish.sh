#!/usr/bin/env bash

fmt_path=github.com/pihao/desktop500px/src/...
go fmt $fmt_path
go test $fmt_path
