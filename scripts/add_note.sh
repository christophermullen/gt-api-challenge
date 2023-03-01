#!/usr/bin/env bash

body="{\"title\":\"${1}\",\"description\":\"${2}\"}"
echo -e "Adding one note:\n${body}"

curl 127.0.0.1:12345/notes \
     -H "Content-Type: application/json" \
     -H "Accept: application/json" \
     -d "${body}"