#!/usr/bin/env bash

RESULT=0

go test github.com/quicklygabbing/users/test/... -v
[[ $? -eq 0 ]] || RESULT=$?

[[ ${RESULT} -eq 0 ]] || exit ${RESULT}
