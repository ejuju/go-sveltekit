#!/bin/bash

set -e

./cicd/local_check_code.sh

# todo: Watch local files and re-run on code change
./cicd/local_build.sh
./cicd/local_run.sh