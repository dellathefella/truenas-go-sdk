API_VERSION=$1
RELEASE=$2

curl --compressed https://www.truenas.com/docs/files/${API_VERSION}.json | yq -P > cfg/${RELEASE}.yaml 