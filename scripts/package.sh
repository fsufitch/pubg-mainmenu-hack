#!/bin/bash

ROOT_PATH=$( cd $(dirname $0) ; pwd -P )/..
cd $ROOT_PATH

GOOS=linux GOARCH=amd64 go build -v -o pubg-hack
npm run webpack

tar cvvfz pubg-hack.tar.gz pubg-hack dist
