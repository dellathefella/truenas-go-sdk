#!/bin/bash

set -e

RELEASE=$1

#!/bin/bash

rm -rf ./*.go
openapi-generator-cli generate -i cfg/openapi.yaml -c cfg/config.yaml -o . -g go --git-user-id dellathfella --git-repo-id truenas-go-sdk
rm -rf test # seems to be broken
patch client.go < client.patch
go mod tidy
go fmt ./...