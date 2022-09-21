#!/bin/bash

set -e

# todo: Watch local files and re-run on code change

seconds=$(date +%s)

TAG="$seconds" ./cicd/local_build.sh
TAG="$seconds" ./cicd/local_run.sh