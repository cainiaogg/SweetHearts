#!/bin/bash
if [ $# -ne 1 ];
then
   echo "Usage: $0 <conf.sh>"
   exit 0
fi

if [ ! -f "$1" ];
then
   echo "'$1' not found"
   exit 1
fi
source $1 #run in the same progress

password = ${password:-njSWf5r0qo4u}
outDir = ${outDir:-$outDir}
