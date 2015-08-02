#!/bin/bash

# Installs go packages, build main go artifact

# clean / remove old artifacts
rm -rf `find pkg -name "*.a" -print`
rm -rf main

# builds the packages first, then builds the entry point
go install bench
go build src/main.go

