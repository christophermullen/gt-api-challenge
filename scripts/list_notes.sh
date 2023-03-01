#!/usr/bin/env bash

echo "Listing all notes"
curl 127.0.0.1:12345/notes \
     -H "Accept: application/json" \