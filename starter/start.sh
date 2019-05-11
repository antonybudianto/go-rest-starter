#!/bin/sh

#if [ "$GO_DEP" = true ]; then
#  echo 'Running dep ensure...'
#  dep ensure
#fi

if [ "$WATCH" = true ] ; then
  echo 'Running realize...'
  realize start
fi

if [ "$BUILD" = true ] ; then
  echo 'Running build'
  go build -o $BIN_NAME ./cmd/$BIN_NAME/main.go
  if [ "$EXEC" = true ] ; then
    ./$BIN_NAME
  fi
fi