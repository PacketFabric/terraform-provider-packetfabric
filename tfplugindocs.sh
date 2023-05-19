#!/usr/bin/env bash
#set -x # uncomment for debug
set -e

if [ ! -d "docs" ] || [ ! -d "templates" ] || [ ! -d "examples" ]; then
    echo "One or more required directories ('docs', 'templates', 'examples') do not exist. Exiting."
    exit 1
fi

os_name=$(uname)

echo -e "\nFormat examples..."
terraform fmt -recursive examples

echo -e "\nGenerating docs..."
if [ "$os_name" == "Linux" ]; then
    tfplugindocs generate --provider-name packetfabric
    cd docs/resources
    for file in *; do mv $file packetfabric_${file%%}; done
    cd ../data-sources
    for file in *; do mv $file packetfabric_${file%%}; done
    cd ../..
    find docs/* -name "*.md" -type f -exec sed -i 's/ Defaults: 0//g' {} \;
    echo -e "\nDocs Updated"

elif [ "$os_name" == "Darwin" ]; then
    tfplugindocs generate --provider-name packetfabric
    cd docs/resources
    for file in *; do mv $file packetfabric_${file%%}; done
    cd ../data-sources
    for file in *; do mv $file packetfabric_${file%%}; done
    cd ../..
    find docs/* -name "*.md" -type f -exec sed -i '' 's/ Defaults: 0//g' {} \;
    echo -e "\nDocs Updated"
else
    echo "This script is only compatible with Linux and macOS."
fi
