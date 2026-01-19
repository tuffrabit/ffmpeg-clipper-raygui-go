@echo off

setlocal
SET "CGO_ENABLED=1"
SET "GOOS=windows"
SET "GOARCH=amd64"
SET "CC=zig cc -lc -target x86_64-windows-gnu -Wl,--subsystem,windows -D__TIMESIZE=64 -D__LIBC -Wno-macro-redefined -Wno-dll-attribute-on-redeclaration -Wtautological-compare"
SET "CXX=zig c++ -lc -target x86_64-windows-gnu -Wl,--subsystem,windows -D__TIMESIZE=64 -D__LIBC -Wno-macro-redefined -Wno-dll-attribute-on-redeclaration -Wtautological-compare"
go build -ldflags "-s -w -H windowsgui"
endlocal
