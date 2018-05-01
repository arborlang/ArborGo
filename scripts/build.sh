#!/usr/bin/env bash
set -e 

ZIP=false
BUILDLAMBDA=build
ENVNAME=
APPNAME="rtbp-data-importers-rpmi"
DIR=$(pwd)

HELPMSG="A script for building the lambdas
usage:
    scripts/build.sh [-l|--lambdas <lambda,names>] [-a| --all] [-z|--zip]
flags:
    -l NAMES | --lambdas NAMES:
        Builds the lambdas with the names in the (comma seperated) list NAMES and put the output in the build directory as a zip file. 
    -h | --help:
        Dispaly this message
    -z | --zip:
        Bundle the executable into a zip file
"

POSITIONAL=()
while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -c|--command)
    BUILDLAMBDA="$2"
    shift
    shift
    ;;
    -e|--env)
    ENVNAME="$2-"
    shift
    shift
    ;;
    -z|--zip)
    ZIP=true
    shift
    ;;
    -h|--help)
    echo -e "$HELPMSG"
    exit 0
    shift
    ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

mkdir -p $DIR/build/
go build -o $DIR/build/arbor $DIR/cmd/arbor/main.go

for i in $(echo $BUILDLAMBDA | sed "s/,/ /g")
do
    go build -o $DIR/build/arbor-$i $DIR/cmd/arbor-$i/main.go
done