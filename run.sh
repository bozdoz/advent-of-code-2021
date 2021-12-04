#!/bin/bash

day=${DAY:-$1}

if [ -z $day ]; then
  echo "You must set \$DAY or pass day directory as an arg"
  exit 1
fi

if [ ! -d $day ]; then
  echo "$day is not a directory!"
  exit 1
fi

cd $day 
gofile=$(ls | grep .go | grep -v test.go)

go run $gofile