#!/bin/bash

export YEAR=${YEAR:-$(date +%Y)}
export DAY=${DAY:-$(date +%-d)}
DAY0="$DAY"
if [ ${#DAY} -eq 1 ]; then
  DAY0="0${DAY}"
fi
FOLDER_NAME=${FOLDER_NAME:-"calendar/${YEAR}/${DAY0}"}
mkdir -p "$FOLDER_NAME"
touch "$FOLDER_NAME/input.txt"
cp "template/main.go" "$FOLDER_NAME/main.go"
cp "template/Makefile" "$FOLDER_NAME/Makefile"
envsubst < "template/README.md" > "$FOLDER_NAME/README.md"