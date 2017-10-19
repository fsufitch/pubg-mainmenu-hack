#!/bin/sh

cd $(dirname $(dirname $0))  # Return to root

GOOS=linux GOARCH=amd64 go build -v -o pubg-hack
npm run webpack

tar cvvfz pubg-hack.tar.gz pubg-hack dist
