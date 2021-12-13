#!/bin/bash

day=${1:-$DAY}

if [ -z $day ]; then
  echo "Set \$DAY or pass day directory as an arg to test single file"
  # test all
  go test ./...
elif [ ! -d $day ]; then
  echo "$day is not a directory!"
  exit 1
else
  # test individual day
  cd $day && go test
fi
