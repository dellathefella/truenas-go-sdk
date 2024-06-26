HOST=$1
RELEASE=$2

curl -k --compressed  http://${HOST}/api/v2.0 | yq -P > cfg/${RELEASE}.yaml
