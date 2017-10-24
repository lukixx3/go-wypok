#!/bin/bash

environment_filepath="$(pwd)/.environment_variables"
required_variables=( "WYKOPAPPKEY" "WYKOPSECRET" "WYKOPCONNECTIONKEY" )

if [ -f "$environment_filepath" ]; then
    source $environment_filepath

    for v in "${required_variables[@]}"; do
        if [ -z ${!v} ]; then
           echo "'$v' has not been set"
           exit 1
        fi
    done

    params="-v"
    if [ "$1" != "" ]; then
        params="$@"
    fi

    go -race -v -coverprofile .testCoverage.txt $params
else
    echo "You need to create '$environment_filepath' file with"
    echo "environment variables for test uses i. e.:"
    echo ""
    for v in "${required_variables[@]}"; do
        echo "export $v=xxx"
    done
    exit 1
fi
