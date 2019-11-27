#!/bin/sh

set -e

[ -z "$DEBUG" ] || set -x

echo "\n===> Generate image...\n"

docker build --no-cache -t c4-payment .

echo "\n===> Docker tag...\n"

docker tag c4-payment fernandocagale/c4-payment

echo "\n===> Docker push...\n"

docker push fernandocagale/c4-payment