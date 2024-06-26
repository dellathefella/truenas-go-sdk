API_VERSION=$1
RELEASE=$2

swagger2openapi --yaml --outfile cfg/${RELEASE}.yaml https://www.truenas.com/docs/files/${API_VERSION}.json