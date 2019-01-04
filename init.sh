#!/bin/sh

exec certmon -config $APP_CONFIG_PATH/config -listen ":$PORT" -ui "/assets/index.html"

