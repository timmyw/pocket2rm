#!/bin/sh

ABSPATH=$(readlink -f $0)
ABSDIR=$(dirname $ABSPATH)

CONFIG=pocket2rm.yaml
if [ $# -eq 1 ]; then
    CONFIG=$1
fi

rm -f $HOME/.config/pocket2rm.yaml 
ln -s $ABSDIR/../confs/$CONFIG $HOME/.config/pocket2rm.yaml 
