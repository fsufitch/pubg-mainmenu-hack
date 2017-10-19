#!/bin/bash

ROOT_PATH=$( cd $(dirname $0) ; pwd -P )/..
cd $ROOT_PATH

mkdir -p hacksim-bin

echo "
===> Building Windows hack simulation"
GOOS=windows GOARCH=amd64 go build -v -o hacksim-bin/PUBG-n00dzmod_4_a64.exe github.com/fsufitch/pubg-mainmenu-hack/hacksim

echo "
===> Building OSX hack simulation"
GOOS=darwin GOARCH=amd64 go build -v -o hacksim-bin/pubg-macport-1.0.1alpha.bin github.com/fsufitch/pubg-mainmenu-hack/hacksim

echo "
===> Building Linux hack simulation"
GOOS=linux GOARCH=amd64 go build -v -o hacksim-bin/battlegrounds-is-not-a-linux-game.bin github.com/fsufitch/pubg-mainmenu-hack/hacksim
