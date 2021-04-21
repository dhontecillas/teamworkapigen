#!/bin/bash

# required env vars:
#   - GITHUB_TOKEN
#
# optional env vars:
# To get help about config options for language specifics:
# docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli config-help -g go

GENLANG=go
export CURDATE=$(date +%Y%m%d_%H%m%S)
export CURDIR=$PWD
export PROJDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && cd .. && pwd )"

echo "************************"
echo "PROJDIR: $PROJDIR"
echo "CURDIR: $CURDIR"
echo "PWD: $PWD"
echo "GH: $GITHUB_TOKEN"
echo "************************"

# export OPENAPIGEN="openapitools/openapi-generator-cli:v5.1.0"
export OPENAPIGEN="dhontecillas/openapi-generator-cli:v0.1"

if [[ -z "$TMPDIR" ]]
then
    export TMPDIR="$HOME/TMP_$CURDATE"
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


echo "***********************"
echo "-----------------------"
echo "git clone https://${GITHUB_TOKEN}@github.com/dhontecillas/teamworkapigoclient.git $TMPDIR"
echo "-----------------------"
echo "***********************"

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
    go mod download

docker run --rm \
    -v $TMPDIR:/apps \
    -w /apps \
    golang:1.16-buster \
    go mod tidy

docker run --rm \
    -v $TMPDIR:/apps \
    -w /apps \
    golang:1.16-buster \
    go build -o createtimelog ./examples/createtimelog

# TODO: check if the examples don't compile

if [[ -z "$GO_CLIENT_BRANCH" ]]; then
    echo "GO_CLIENT_BRANCH: ** $GO_CLIENT_BRANCH ** - GITHUB_HEAD_REF: ** $GITHUB_HEAD_REF **"
    GO_CLIENT_BRANCH="update/at_$CURDATE"
else
    echo "GO_CLIENT_BRANCH: ** $GO_CLIENT_BRANCH ** - GITHUB_HEAD_REF: ** $GITHUB_HEAD_REF **"
fi

if ! [[ -z "CREATE_PR" ]]
then
    echo "Commiting to a new branch: $GO_CLIENT_BRANCH"
    GO_CLIENT_MSG="Update API with $GO_CLIENT_BRANCH at $CURDATE"
    git config --global user.email $GH_USER_EMAIL
    git config --global user.name $GH_USER_NAME
    git checkout $GO_CLIENT_BRANCH || git checkout -b $GO_CLIENT_BRANCH
    git add ./pmv1
    git add ./projv1
    git add ./projv2
    git add ./projv3
    echo "going to commit: $GO_CLIENT_MSG"
    git commit -am "$GO_CLIENT_MSG"
    git push -u origin $GO_CLIENT_BRANCH
    gh pr create --base main \
        --head $GO_CLIENT_BRANCH \
        --title "$GO_CLIENT_MSG" \
        --body "$GO_CLIENT_MSG"
fi
