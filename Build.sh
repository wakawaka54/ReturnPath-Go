#!/usr/bin/env bash

#exit if any command fails
set -e

cd ./rpfrontend

go test -v

cd ..
cd ./rpapi

go test -v
