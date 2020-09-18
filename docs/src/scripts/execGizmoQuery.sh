#!/bin/bash

if [ "$#" -lt 2 ]; then
    echo "Illegal number of parameters" >&2
    echo "Usage: execGizmoQuery <data-to-load> <query-file> [<out-file-name>]" >&2
    exit 1
fi

if [ "$#" -eq 3 ]; then
    cat $2 | cayley query  --logtostderr true --load $1 | cat > $3
else
    cat $2 | cayley query  --logtostderr true --load $1
fi
