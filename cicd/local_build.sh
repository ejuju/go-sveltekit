#!/bin/bash

start=$(date +%s)

if test -z "$REPO"; then REPO="localhost/go-sveltekit"; fi
printf 'Using repo: %s \n' "$REPO"

if test -z "$TAG"; then TAG="$(git rev-parse --short HEAD)"; fi
printf 'Using tag: %s \n' "$TAG"

podman build . \
    -f ./containerfile \
    -t "$REPO:$TAG" \
    --progress plain

end=$(date +%s)
runtime=$((end-start))
printf 'Took %s seconds to build container image \n\n' "$runtime"