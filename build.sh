#!/bin/bash

export GOPATH=`pwd`
export GOBIN=`pwd`/bin
go get github.com/bobby96333/goSqlHelper
go install csvdump
bin/csvdump -h