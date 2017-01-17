#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -o socket.out .
