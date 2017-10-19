#!/bin/sh

cd $(dirname $(dirname $0))  # Return to root

go build -v -o pubg-hack
npm run webpack

tar cvvfz pubg-hack.tar.gz pubg-hack dist
