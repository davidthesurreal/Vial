#!/bin/bash

FILE="vial"

if [[ -f "$FILE" ]]; then
    SIZE_KB=$(du -k "$FILE" | cut -f1)
    SIZE_MB=$(du -m "$FILE" | cut -f1)

    echo "Size of '$FILE':"
    echo "$SIZE_KB kB"
    echo "$SIZE_MB MB"
else
    echo "File '$FILE' not found."
fi
