#!/bin/bash

TMP_DIR="tmp"

PYTHON_ARCHIVE_NAME=$(ls sdk/python/bin/dist/*.tar.gz)
ARCHIVE_TMP_DIR=${TMP_DIR}/$(basename ${PYTHON_ARCHIVE_NAME} .tar.gz)

S3_DEST="s3://${TARGET_S3_BUCKET}${TARGET_DIRECTORY}$(basename ${PYTHON_ARCHIVE_NAME})"
S3_HTTP_URL="https://${TARGET_S3_BUCKET}.s3.${AWS_DEFAULT_REGION}.amazonaws.com${TARGET_DIRECTORY}"

ROOT=$(dirname $0)/..
VERSION=$(jq -r '.version' < "${ROOT}/sdk/nodejs/bin/package.json")
PULUMI_PLUGIN_NAME="pulumi-resource-heroku-${VERSION}-$(go env GOOS)-$(go env GOARCH).tar.gz"

echo "Unzip python archive ${PYTHON_ARCHIVE_NAME}"
mkdir ${TMP_DIR} && tar -xvzf ${PYTHON_ARCHIVE_NAME} -C ${TMP_DIR}

echo "Manipulate setup.py"
sed -i "s|'pulumi', 'plugin', 'install',|'pulumi', 'plugin', 'install', '--server', '${S3_HTTP_URL}',|g" ${ARCHIVE_TMP_DIR}/setup.py

echo "Create new archive"
tar cfvz ${PYTHON_ARCHIVE_NAME} -C ${ARCHIVE_TMP_DIR} .


echo "Publish to ${S3_DEST}"
aws s3 cp --only-show-errors ${PYTHON_ARCHIVE_NAME} ${S3_DEST}
