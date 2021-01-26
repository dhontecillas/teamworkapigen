#!/bin/bash

# To get help about config options for language specifics:
# docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli config-help -g go

set -e

GENLANG=go
THISDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"


for spec in "project-manager" "projectsapigo.v1" "projectsapigo.v2" "projectsapigo.v3"
do
mkdir -p $THISDIR/sdks/$GENLANG/$spec
docker run --rm -u 1000:1000 -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/openapi_specs/$spec.yaml \
    -g $GENLANG \
    -c /local/configs/$GENLANG/$spec.json \
    -o /local/sdks/$GENLANG/$spec
done

find $THISDIR/sdks/go -name 'go.mod' | xargs rm
find $THISDIR/sdks/go -name 'go.sum' | xargs rm

TWGOOPENAPI=$PWD/../teamworkapigoclient/pkg/openapi
cd $TWGOOPENAPI
rm -rf pmv1
mv $THISDIR/sdks/go/project-manager ./pmv1
rm -rf projv1
mv $THISDIR/sdks/go/projectsapigo.v1 ./projv1
rm -rf projv2
mv $THISDIR/sdks/go/projectsapigo.v2 ./projv2
rm -rf projv3
mv $THISDIR/sdks/go/projectsapigo.v3 ./projv3
