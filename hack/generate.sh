#!/bin/bash

export YEAR=${YEAR:-$(date +%Y)}
export DAY=${DAY:-$(date +%-d)}
export DAY0=${DAY0:-$(date +%d)}
export FOLDER_NAME=${FOLDER_NAME:-"calendar/${YEAR}/${DAY0}"}

mkdir -p "$FOLDER_NAME"
touch "$FOLDER_NAME/input.txt"
cp "template/main.go" "$FOLDER_NAME/main.go"
cp "template/Makefile" "$FOLDER_NAME/Makefile"
envsubst < "template/README.md" > "$FOLDER_NAME/README.md"