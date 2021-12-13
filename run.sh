#!/bin/bash

day=${1:-$DAY}

if [ -z $day ]; then
  echo "You must set \$DAY or pass day directory as an arg"
  exit 1
fi

if [ ! -d $day ]; then
  echo "$day is not a directory!"
  exit 1
fi

cd $day 

go run .