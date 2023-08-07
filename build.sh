#!/bin/bash


# add more builds if necessary

# build for mac
GOOS=darwin GOARCH=arm64 go build -o build/mac-m1/wad2linkedin

# build for windows
GOOS=windows GOARCH=amd64 go build -o build/windows-amd64/wad2linkedin.exe