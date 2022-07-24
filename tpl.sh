#!/bin/bash
ip=$1
out=$2
in=$3
while true
do
  ssh -o "StrictHostKeyChecking=no" -L 0.0.0.0:$2:$1:$3 -N root@$1 
done

