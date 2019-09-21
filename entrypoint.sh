#!/usr/bin/env sh

set -e

if [ -z "${INPUT_USERNAME}" ]; then
  echo "Username is empty. Please set with.username to login to docker registry."
fi

if [ -z "${INPUT_PASSWORD}" ]; then
  echo "Password is empty. Please set with.password to login to docker registry."
fi

echo "${INPUT_PASSWORD}" | docker login -u ${INPUT_USERNAME} --password-stdin ${INPUT_REGISTRY}

/publisher
