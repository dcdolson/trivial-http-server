#!/bin/sh

if [ $# -ne 1 ]
then
    echo "Usage: $0 <web-root-folder>" >&2
    exit 2
fi

folder=$(realpath "$1")

./serve "$folder"

