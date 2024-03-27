#!/usr/bin/bash

docker build --no-cache -t registry.rmutsv.app/authapon/apiserver:pro2 .
docker push registry.rmutsv.app/authapon/apiserver:pro2