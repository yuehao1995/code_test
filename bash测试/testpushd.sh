#!/usr/bin/env bash
pushd $(dirname $0) > /dev/null
echo $(pwd -P)
echo `pwd`
popd > /dev/null
echo `pwd`