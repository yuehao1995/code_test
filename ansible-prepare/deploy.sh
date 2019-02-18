#!/usr/bin/env bash

#!/bin/bash
# Created by g7tianyi on 8/10/2018

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

# ================================================================================

export ANSIBLE_HOST_KEY_CHECKING=False

HOSTS_FILE=${SCRIPTPATH}/inventories/eecontract.ee-chain.com/hosts

log "==>>【部署】4.1 初始化目标机器.."
ansible-playbook -i ${HOSTS_FILE} ${SCRIPTPATH}/deploy_prepare.yml