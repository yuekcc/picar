#!/bin/env bash
go build -ldflags="-s -w -X 'main.COMMIT_ID=`git log --pretty=format:'%h' -1`'" -v -a