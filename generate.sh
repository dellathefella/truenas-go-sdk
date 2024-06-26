#!/bin/bash

set -e

RELEASE=$1

rm -rf "${RELEASE}"
openapi-generator-cli generate -i cfg/"${RELEASE}".yaml -c cfg/"${RELEASE}"_config.yaml -o ./"${RELEASE}" -g go --git-user-id dellathefella --git-repo-id truenas-go-sdk --skip-validate-spec --legacy-discriminator-behavior
rm -rf "${RELEASE}"/test # seems to be broken
patch "${RELEASE}"/client.go < client.patch
mv "${RELEASE}"/go.mod .
mv "${RELEASE}"/go.sum .
sed -i '1s/.*/module github.com\/jdella\/truenas-go-sdk/' go.mod
cd "${RELEASE}" || exit
go mod tidy
go fmt ./...
cd ..