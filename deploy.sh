#!/bin/bash

set -ex

BINARY_NAME=$1

rm -rfv public/assets
git checkout -b build
GOOS=linux buffalo build -z
git add .
git add -f "bin/${BINARY_NAME}"
git commit -a -m "binary commit"
git push heroku build:master --force
git checkout master
git branch -D build
rm -rf bin/

heroku run "bin/${BINARY_NAME}" migrate -e production