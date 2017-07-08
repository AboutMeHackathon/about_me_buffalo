#!/bin/bash

HEROKU_APP=$1
BINARY_NAME=$2

echo -e "web: bin/${BINARY_NAME}" > Procfile
heroku create $HEROKU_APP --buildpack http://github.com/ryandotsmith/null-buildpack.git
heroku config:set GO_ENV=production
heroku addons:create heroku-postgresql:hobby-dev