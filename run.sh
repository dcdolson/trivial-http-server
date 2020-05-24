#!/bin/sh

if [ $# -ne 2 ]
then
    echo "Usage: $0 <web-root-folder> <port>" >&2
    exit 2
fi

folder=$(realpath "$1")

./serve "$folder" "$2"

