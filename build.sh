#!/bin/bash

CGO_ENABLED=1 \
GOOS=linux \
GOARCH=amd64 \
CC="zig cc -lc -target x86_64-linux-gnu.2.39 -D__TIMESIZE=64 -D__LIBC -Wno-macro-redefined -I/usr/include -L/usr/lib -L/usr/lib64 -L/lib/x86_64-linux-gnu" \
CXX="zig c++ -lc -target x86_64-linux-gnu.2.39 -D__TIMESIZE=64 -D__LIBC -Wno-macro-redefined -I/usr/include -L/usr/lib -L/usr/lib64 -L/lib/x86_64-linux-gnu" \
go build -ldflags "-s -w"
