#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go build ./model.go
mv ./model ./bin/model
git add .
git commit -m'-'
git push origin master

say "提交成功"