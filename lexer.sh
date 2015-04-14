#!/bin/bash -e

ragel -Z lexer_go.rl 
go build -o app .
./app

if [ $(which dot) -a "$1" != "" ]; then
  ragel -V lexer_go.rl > app.dot
  dot -Tpng app.dot > app.png
  open app.png
fi
