#!/bin/bash

# required env vars:
#   - GITHUB_TOKEN
#
# optional env vars:
# To get help about config options for language specifics:
# docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli config-help -g go

set -e

GENLANG=go
export CURDATE=$(date +%Y%m%d_%H%m%S)
export CURDIR=$PWD
export PROJDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && cd .. && pwd )"

echo "************************"
echo "PROJDIR: $PROJDIR"
echo "CURDIR: $CURDIR"
echo "PWD: $PWD"
echo "************************"

# export OPENAPIGEN="openapitools/openapi-generator-cli:v5.1.0"
export OPENAPIGEN="dhontecillas/openapi-generator-cli:v0.1"

if [[ -z "$TMPDIR" ]]
then
    # export TMPDIR="$HOME/GoBuild/$CURDATE"
    export TMPDIR="/GoBuild/$CURDATE"
fi

docker pull $OPENAPIGEN
docker pull golang:1.16-buster

# delete any existing files from any previous generation process
rm -rf $PROJDIR/sdks/$GENLANG

for spec in "project-manager" "projectsapigo.v1" "projectsapigo.v2" "projectsapigo.v3"
do
mkdir -p $PROJDIR/sdks/$GENLANG/$spec
docker run --rm \
    -v $PWD:/local \
    $OPENAPIGEN generate \
    -i /local/openapi_specs/$spec.yaml \
    -g $GENLANG \
    -c /local/configs/$GENLANG/$spec.json \
    -o /local/sdks/$GENLANG/$spec
done

# remove the module files created by the generator
find $PROJDIR/sdks/go -name 'go.mod' | xargs rm
find $PROJDIR/sdks/go -name 'go.sum' | xargs rm


git clone https://${GITHUB_TOKEN}@github.com/dhontecillas/teamworkapigoclient.git $TMPDIR
cd $TMPDIR

rm -rf pmv1
mv $PROJDIR/sdks/go/project-manager ./pmv1
rm -rf projv1
mv $PROJDIR/sdks/go/projectsapigo.v1 ./projv1
rm -rf projv2
mv $PROJDIR/sdks/go/projectsapigo.v2 ./projv2
rm -rf projv3
mv $PROJDIR/sdks/go/projectsapigo.v3 ./projv3

docker run --rm \
    -v $TMPDIR:/apps \
    -w /apps \
    golang:1.16-buster \
    go build -v ./examples/createtimelog \
    -o createtimelog
