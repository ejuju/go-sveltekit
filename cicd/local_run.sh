#!/bin/bash

set -e

printf 'Running app image: %s \n' "$REPO"

if test -z "$REPO"; then REPO="localhost/go-sveltekit"; fi
printf 'With REPO: %s \n' "$REPO"

if test -z "$TAG"; then TAG="$(git rev-parse --short HEAD)"; fi
printf 'With TAG: %s \n' "$TAG"

if test -z "$PORT"; then PORT="3420"; fi
printf 'With PORT: %s \n' "$PORT"

sudo docker run \
    -p "$PORT":"$PORT" \
    --expose="$PORT" \
    --env PORT="$PORT" \
    "$REPO:$TAG"