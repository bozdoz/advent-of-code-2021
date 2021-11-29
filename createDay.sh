#!/bin/bash

NEW_DAY=$1
NEW_DAY_NAME=$2
TEMPLATE=${3:-'01'}

if [ -z $NEW_DAY ]; then
  echo "Provide ## for new day directory"
  exit 1
fi

if [ -z $NEW_DAY_NAME ]; then
  echo "Provide go filename for new day: one, two, three"
  exit 1
fi

cp -r $TEMPLATE $NEW_DAY

cd $NEW_DAY

for f in *.go; do
  if [ $f != "*_test*" ]; then
    OLD_DAY_NAME=$(echo "${f/.go/}")
    break
  fi
done;

for f in *.go; do
  mv $f ${f/${OLD_DAY_NAME}/${NEW_DAY_NAME}}
done;