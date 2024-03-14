#!/usr/bin/bash

docker build --no-cache -t registry.rmutsv.app/authapon/apiserver:pro .
docker push registry.rmutsv.app/authapon/apiserver:pro