#! /usr/bin/env sh

echo "shpr action starting..."
echo "platform: $INPUT_PLATFORM"
echo "build-folder: $INPUT_BUILD_FOLDER"
/shpr app start $INPUT_PLATFORM $INPUT_BUILD_FOLDER
/shpr app list

