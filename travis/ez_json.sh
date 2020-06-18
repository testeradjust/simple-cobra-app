#!/bin/bash
set -e          # abort on errors

echo "installing easyjson ..."
go get -u github.com/mailru/easyjson/...

echo "easyjson generate ..."
easyjson -all $(find . -name "ez_json*")

echo "easyjson done <3"
