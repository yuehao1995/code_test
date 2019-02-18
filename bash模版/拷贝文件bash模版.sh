#!/bin/bash
# Created by g7tianyi on 29/9/2018

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

SPACE="    "
Target_Ip="182.61.47.204"

# SEED=$(python -c "import random; print ''.join([random.choice('0123456789abcdef') for n in xrange(10)])")
# WORKSPACE=/opt/eechain/ansible-fabric/${SEED}
WORKSPACE=${SCRIPTPATH}/tmp
CONFIGDIR=${SCRIPTPATH}/cfg

export PYTHONPATH=${SCRIPTPATH}:${PYTHONPATH}

# ================================================================================

function usage() {
    echo -e "${SPACE}"
    echo -e "${SPACE}用法: "
    echo -e "${SPACE}"
    echo -e "${SPACE}${SPACE}--help                     : 显示帮助信息"
    echo -e "${SPACE}${SPACE}--config=<your_config.yml> : 指定配置文件名，如: --config=example.yml"
    echo -e "${SPACE}"
    echo -e "${SPACE}示例: "
    echo -e "${SPACE}"
    echo -e "${SPACE}${SPACE}${SCRIPTFILE} --config=example.yml"
    echo -e "${SPACE}"
}

function run() {
    log "==>>【准备】1. 创建工作目录 ${WORKSPACE}.."
    mkdir -p ${WORKSPACE}

    log "==>>【准备】2. 复制msp证书到目标机器.."
    scp -r ${WORKSPACE}hyperledger-fabric-ansible/output /contract.ee-chain.com/ctypto-config root@127.0.0.1:/ebcgateway/assets/service
    log "==>>【准备】3.复制创世区块及交易信息到目标机器.."
    scp -r ${WORKSPACE}hyperledger-fabric-ansible/output /contract.ee-chain.com/channel-aitifacts root@127.0.0.1:/ebcgateway/assets/service

    log "==>>【部署】4. 复制多个配置文件到目标机器.."
    scp -r ${WORKSPACE}hyperledger-fabric-ansible/inventories /contract.ee-chain.com/two_config.yaml  root@127.0.0.1:/ebcgateway/assets/service

    log "==>>【完成】"

    cat << EOF

         ______ ______ _____
        |  ____|  ____/ ____|
        | |__  | |__ | |
        |  __| |  __|| |
        | |____| |___| |____
        |______|______\_____|

EOF
}

# ================================================================================

while [ $# -gt 0 ]; do
    case "$1" in
        --config=*)
            NETWORK_CONFIG="${1#*=}"
            ;;
        --help)
            usage
            exit
            ;;
    esac
    shift
done

run
