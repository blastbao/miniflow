#!/bin/bash

env GOOS=linux GOARCH=amd64 go build -o miniflow-linux-amd64 main.go
