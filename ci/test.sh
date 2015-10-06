#!/bin/bash -x
go test -race -v $(go list ./... | grep -v /vendor/)

