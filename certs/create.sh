#!/bin/bash

set -e

function logger_info () {
    echo "AZURITE_CONTROLLER: INFO: $1"
}

logger_info "cloning mkcert repository started"
git clone https://github.com/FiloSottile/mkcert
logger_info "cloning mkcert repository completed"

logger_info "building mkcert started"
cd mkcert && go build -ldflags "-X main.Version=$(git describe --tags)" && cd ../
logger_info "building mkcert completed"

logger_info "creating locally-trusted development certificates started"
./mkcert/mkcert -install
./mkcert/mkcert 127.0.0.1
logger_info "creating locally-trusted development certificates completed"

logger_info "deleting mkcert repository started"
rm -rf mkcert
logger_info "deleting mkcert repository completed"
