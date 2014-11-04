#!/bin/bash
SCRIPTPATH="$(python -c 'import os,sys;print os.path.realpath(sys.argv[1])' c)/golang_talk"
#echo "$SCRIPTPATH"
EMPTY=''
SCRIPT="${SCRIPTPATH/\/c\///$EMPTY}"
UNAMEV=`uname`
PLATFORM=`echo $UNAMEV | awk '{print tolower($0)}'`
#echo "$SCRIPT"
export GOPATH="$SCRIPT/"
export GOROOT="$SCRIPT/.govenv"
export GOARCH="amd64"
export GOOS=$PLATFORM
export GOHOSTOS=$GOOS
export GOHOSTARCH=$GOARCH
export PATH=$PATH:$GOROOT/bin:$GOROOT/pkg:$GOPATH/bin
