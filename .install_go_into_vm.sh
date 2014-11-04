#!/bin/bash
VERSION="1.3.1"
UNAMEV=`uname`
PLATFORM=`echo $UNAMEV | awk '{print tolower($0)}'`
PLATFORMVERSION=''
#echo $PLATFORM
if [ "$PLATFORM" == "darwin" ]; then
    PLATFORMVERSION="-osx10.8"
fi

if [ "$1" == "--32" ]; then
    DFILE="go$VERSION.$PLATFROM-386$PLATFORMVERSION.tar.gz"
elif [ "$1" == "--64" ]; then
    DFILE="go$VERSION.$PLATFORM-amd64$PLATFORMVERSION.tar.gz"
else
    # Default to the 64 bit version
    DFILE="go$VERSION.$PLATFORM-amd64$PLATFORMVERSION.tar.gz"
fi
#echo $DFILE
if [[ -d $GOROOT/ || -d $GOROOT/go ]]; then
    #echo "Installation directories already exist. Exiting."
    # If go is already installed just return normally
    exit 0
fi
mkdir $GOROOT
wget https://storage.googleapis.com/golang/$DFILE -O /tmp/go.tar.gz
if [ $? -ne 0 ]; then
    echo "Download failed! Exiting."
    exit 1
fi

tar -C $GOROOT -xzf /tmp/go.tar.gz
mv $GOROOT/go/* $GOROOT/
rm -fr $GOROOT/go/
mkdir -p $GOPATH/src
mkdir -p $GOPATH/pkg
mkdir -p $GOPATH/bin
rm -f /tmp/go.tar.gz
