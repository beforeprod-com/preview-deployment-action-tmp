#! /usr/bin/env sh

echo "shpr action starting..."
echo "platform: $INPUT_PLATFORM"
echo "build-folder: $INPUT_BUILD_FOLDER"

# Capture the output of shpr app start and extract the URL
URL=$(/shpr app start $INPUT_PLATFORM $INPUT_BUILD_FOLDER | grep -o 'https://[^[:space:]]*')
echo "::set-output name=url::$URL"

/shpr app list

# Set the time output
echo "::set-output name=time::$(date -u +"%Y-%m-%dT%H:%M:%SZ")"
