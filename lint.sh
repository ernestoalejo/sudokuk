#!/bin/bash


FILTER="-runtime/references,-whitespace/labels,-legal/copyright,-readability/streams,-build/include"

if [ "$1" = "" ]; then
  find -name '*.cc' -o -name '*.h' | xargs 'cpplint.py' --filter=$FILTER
else
  cpplint.py --filter=$FILTER $1 
fi
