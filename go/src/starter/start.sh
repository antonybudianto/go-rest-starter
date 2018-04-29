#!/bin/sh

if [ "$GO_DEP" = true ]; then
  echo 'Running dep ensure...'
  dep ensure
else
  go install -v .
fi

if [ "$DEBUG" = true ] ; then
  modd -f=./modd-debug.conf
fi

if [ "$WATCH" = true ] ; then
  echo 'Running modd...'
  modd
fi

if [ "$BUILD" = true ] ; then
  echo 'Running build'
  go build -o $BIN_NAME .
  if [ "$EXEC" = true ] ; then
    ./$BIN_NAME
  fi
fi