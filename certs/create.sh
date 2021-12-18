#!/bin/bash

set -e

echo "AZURITE_CONTROLLER: INFO: cloning mkcert repository started"
git clone https://github.com/FiloSottile/mkcert
echo "AZURITE_CONTROLLER: INFO: cloning mkcert repository completed"

echo "AZURITE_CONTROLLER: INFO: building mkcert started"
cd mkcert && go build -ldflags "-X main.Version=$(git describe --tags)" && cd ../
echo "AZURITE_CONTROLLER: INFO: building mkcert completed"

echo "AZURITE_CONTROLLER: INFO: creating locally-trusted development certificates started"
./mkcert/mkcert -install
./mkcert/mkcert 127.0.0.1
echo "AZURITE_CONTROLLER: INFO: creating locally-trusted development certificates completed"

echo "AZURITE_CONTROLLER: INFO: deleting mkcert repository started"
rm -rf mkcert
echo "AZURITE_CONTROLLER: INFO: deleting mkcert repository completed"
