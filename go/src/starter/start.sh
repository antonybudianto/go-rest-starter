#!/bin/sh

if [ "$GO_DEP" = true ]; then
  dep ensure -update
else
  go install -v .
fi

if [ "$DEBUG" = true ] ; then
  modd -f=./modd-debug.conf
fi

if [ "$WATCH" = true ] ; then
  modd
fi

if [ "$BUILD" = true ] ; then
  go build -o $BIN_NAME .
  if [ "$EXEC" = true ] ; then
    ./$BIN_NAME
  fi
fi