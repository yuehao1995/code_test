#!/bin/bash
# Created by martin 01/02/2019

set -e

pushd $(dirname $0) > /dev/null
SCRIPTPATH=$(pwd -P)
popd > /dev/null
SCRIPTFILE=$(basename $0)

function log() {
    echo "================================================================================"
    echo "$(date +'%Y-%m-%d %H:%M:%S%z') [INFO] - $@"
    echo ""
}

function err() {
    echo "================================================================================"
    echo "$(date +'%Y-%m-%d %H:%M:%S%z') [ERRO] - $@" >&2
}


function showUsage() {
    echo -e ""
    echo -e "${SCRIPTFILE} [-h] --imageVersion=<imageVersion> --stage=<stage>"
    echo -e ""
    echo -e "    --stage=<stage>                 : Specify the stage, default to 'prod'"
    echo -e "    --registry=<registry>           : Specify the docker registry"
    echo -e "    --imageVersion=<imageVersion>   : Specify the image version"
    echo -e "    -h                              : Show this message"
    echo -e ""
}

function scpToServer() {
    scp -r  ${SCRIPTPATH}/pkg/api/epoll_client/bin root@116.62.118.133:/root/opt/epoll/bin
}
# ================================================================================


scpToServer